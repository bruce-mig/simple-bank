package gapi

import (
	"context"
	"errors"
	"time"

	db "github.com/bruce-mig/simple-bank/db/sqlc"
	"github.com/bruce-mig/simple-bank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RenewAccessToken(ctx context.Context, req *pb.RenewAccessTokenRequest) (*pb.RenewAccessTokenResponse, error) {
	refreshPayload, err := server.tokenMaker.VerifyToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to verify refresh token",
		)
	}
	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"user session not found",
			)
		}
		return nil, status.Errorf(
			codes.Internal,
			"failed to get session",
		)
	}

	if session.IsBlocked {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"blocked session",
		)
	}

	if session.Username != refreshPayload.Username {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"incorrect session user",
		)
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"mismatched session token",
		)
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"expired session",
		)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		refreshPayload.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to generate access token",
		)
	}
	res := &pb.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return res, nil
}
