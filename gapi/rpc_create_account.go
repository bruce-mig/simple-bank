package gapi

import (
	"context"

	db "github.com/bruce-mig/simple-bank/db/sqlc"
	"github.com/bruce-mig/simple-bank/pb"
	"github.com/bruce-mig/simple-bank/util"
	"github.com/bruce-mig/simple-bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole, util.DepositorRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if authPayload.Role != util.BankerRole && authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"cannot create account for other users",
		)
	}

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			return nil, status.Errorf(
				codes.AlreadyExists,
				"account already exists",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to create account",
		)
	}

	res := &pb.CreateAccountResponse{
		Account: convertAccount(account),
	}

	return res, nil
}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateCurrency(req.GetCurrency()); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
