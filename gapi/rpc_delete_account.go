package gapi

import (
	"context"
	"errors"
	"fmt"

	db "github.com/bruce-mig/simple-bank/db/sqlc"
	"github.com/bruce-mig/simple-bank/pb"
	"github.com/bruce-mig/simple-bank/util"
	"github.com/bruce-mig/simple-bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Only bankers are authorized to delete for the following reasons:
// 1. to prevent depositors with overdrawn accounts from just deleting the account to avoid paying their overdraft amounts.
// 2. to prevent depositors from accidentally deleting their accounts and thus "lose" their funds.
func (server *Server) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateDeleteAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.Role != util.BankerRole {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"you do not have the required permissions to execute the task",
		)
	}

	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"account not found",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to get account",
		)
	}

	if account.Owner != req.GetOwner() {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"account id and owner mismatch",
		)
	}

	arg := db.DeleteAccountParams{
		ID:    req.GetId(),
		Owner: req.GetOwner(),
	}

	err = server.store.DeleteAccount(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"account not found",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to delete account",
		)
	}
	res := &pb.DeleteAccountResponse{
		Response: fmt.Sprintf("account [id:%v | owner: %s] has been successfully deleted", req.GetId(), req.GetOwner()),
	}

	return res, nil
}

func validateDeleteAccountRequest(req *pb.DeleteAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateAccountID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := val.ValidateUsername(req.GetOwner()); err != nil {
		violations = append(violations, fieldViolation("owner", err))
	}

	return violations
}
