package gapi

import (
	"context"
	"errors"

	db "github.com/bruce-mig/simple-bank/db/sqlc"
	"github.com/bruce-mig/simple-bank/pb"
	"github.com/bruce-mig/simple-bank/util"
	"github.com/bruce-mig/simple-bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Only bankers are authorized to update for the following reason:
// 1. to prevent depositors from fraudulently altering their account balances.
func (server *Server) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.Role != util.BankerRole {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"you do not have the required permissions to execute the task",
		)
	}

	arg := db.AddAccountBalanceParams{
		Amount: req.GetAmount(),
		ID:     req.GetId(),
	}

	account, err := server.store.AddAccountBalance(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"account does not exist",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to update balance",
		)
	}

	res := &pb.UpdateAccountResponse{
		Account: convertAccount(account),
	}

	return res, nil
}

func validateUpdateAccountRequest(req *pb.UpdateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateAccountID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := val.ValidateAmount(req.GetAmount()); err != nil {
		violations = append(violations, fieldViolation("amount", err))
	}

	return violations
}
