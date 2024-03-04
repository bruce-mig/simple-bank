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

func (server *Server) GetTransfer(ctx context.Context, req *pb.GetTransferRequest) (*pb.GetTransferResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole, util.DepositorRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateGetTransferRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	transfer, err := server.store.GetTransfer(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"transfer not found",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to get transfer",
		)
	}

	fromAccount, err := server.store.GetAccount(ctx, transfer.FromAccountID)
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

	toAccount, err := server.store.GetAccount(ctx, transfer.ToAccountID)
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
	if authPayload.Role != util.BankerRole && fromAccount.Owner != authPayload.Username && toAccount.Owner != authPayload.Username {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"transfer doesn't belong to the authenticated user",
		)
	}

	res := &pb.GetTransferResponse{
		Transfer: convertTransfer(transfer),
	}
	return res, nil
}

func validateGetTransferRequest(req *pb.GetTransferRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateTransferID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
