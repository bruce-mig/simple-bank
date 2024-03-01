package gapi

import (
	"context"
	"database/sql"
	"testing"
	"time"

	mockdb "github.com/bruce-mig/simple-bank/db/mock"
	db "github.com/bruce-mig/simple-bank/db/sqlc"
	"github.com/bruce-mig/simple-bank/pb"
	"github.com/bruce-mig/simple-bank/token"
	"github.com/bruce-mig/simple-bank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteAccountAPI(t *testing.T) {
	user, _ := randomUser(t, util.BankerRole)
	other, _ := randomUser(t, util.DepositorRole)
	account := randomAccount(user.Username)

	testCases := []struct {
		name          string
		req           *pb.DeleteAccountRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.DeleteAccountResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.DeleteAccountParams{
					ID:    account.ID,
					Owner: account.Owner,
				}

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(arg.ID)).
					Times(1).
					Return(account, nil)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "Account NotFound",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.DeleteAccountParams{
					ID:    account.ID,
					Owner: user.Username,
				}

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(arg.ID)).
					Times(1).
					Return(db.Account{}, db.ErrRecordNotFound)
				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.DeleteAccountParams{
					ID:    account.ID,
					Owner: user.Username,
				}

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(arg.ID)).
					Times(1).
					Return(account, nil)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InvalidID",
			req: &pb.DeleteAccountRequest{
				Id:    0,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "Invalid Owner Name",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: "0wn3r#",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "Depositor Cannot Delete An Account",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, other.Username, other.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: account.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "Account ID and Owner Mismatch",
			req: &pb.DeleteAccountRequest{
				Id:    account.ID,
				Owner: other.Username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.DeleteAccountParams{
					ID:    account.ID,
					Owner: other.Username,
				}

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(arg.ID)).
					Times(1).
					Return(account, nil)

				store.EXPECT().
					DeleteAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)

			tc.buildStubs(store)
			server := newTestServer(t, store, nil)

			ctx := tc.buildContext(t, server.tokenMaker)
			res, err := server.DeleteAccount(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
