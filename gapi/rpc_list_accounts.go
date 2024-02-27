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

func (server *Server) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole, util.DepositorRole})
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListAccountsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to fetch accounts",
		)
	}

	for _, acc := range accounts {
		if authPayload.Role != util.BankerRole && acc.Owner != authPayload.Username {
			return nil, status.Errorf(
				codes.PermissionDenied,
				"account doesn't belong to the authenticated user",
			)
		}
	}

	res := &pb.ListAccountsResponse{
		Accounts: convertAccounts(accounts),
	}
	return res, nil
}

func validateListAccountsRequest(req *pb.ListAccountsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidatePageID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}

	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}

	return violations
}
