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

func (server *Server) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.CreateTransferResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole, util.DepositorRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := ValidateCreateTransferRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	fromAccount, valid, _ := server.validAccount(ctx, req.GetFromAccountId(), req.GetCurrency())

	if !valid {
		return nil, status.Errorf(
			codes.Internal,
			"account validation failed",
		)
	}

	if authPayload.Role != util.BankerRole && fromAccount.Owner != authPayload.Username {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"from account doesn't belong to the authenticated user",
		)
	}
	_, valid, _ = server.validAccount(ctx, req.GetToAccountId(), req.GetCurrency())
	if !valid {
		return nil, status.Errorf(
			codes.Internal,
			"account validation failed",
		)
	}

	arg := db.TransferTxParams{
		FromAccountID: req.GetFromAccountId(),
		ToAccountID:   req.GetToAccountId(),
		Amount:        req.GetAmount(),
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to execute transfer",
		)
	}

	res := convertTranferTxResponse(result)
	return res, nil
}

func (server *Server) validAccount(ctx context.Context, accountID int64, currency string) (db.Account, bool, error) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return account, false, status.Errorf(
				codes.NotFound,
				"account doesn't exist",
			)
		}
		return account, false, status.Errorf(
			codes.Internal,
			"failed to get account",
		)
	}

	if account.Currency != currency {
		return account, false, status.Errorf(
			codes.Internal,
			fmt.Sprintf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency),
		)

	}

	return account, true, nil
}

func ValidateCreateTransferRequest(req *pb.CreateTransferRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateAccountID(req.GetFromAccountId()); err != nil {
		violations = append(violations, fieldViolation("from_account_id", err))
	}

	if err := val.ValidateAccountID(req.GetToAccountId()); err != nil {
		violations = append(violations, fieldViolation("to_account_id", err))
	}

	if err := val.ValidateTransferAmount(req.GetAmount()); err != nil {
		violations = append(violations, fieldViolation("amount", err))
	}

	if err := val.ValidateCurrency(req.GetCurrency()); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	return violations
}
