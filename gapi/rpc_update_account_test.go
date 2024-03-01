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

func TestUpdateAccountAPI(t *testing.T) {
	user, _ := randomUser(t, util.BankerRole)
	other, _ := randomUser(t, util.DepositorRole)
	account := randomAccount(other.Username)
	amount := util.RandomMoney()

	testCases := []struct {
		name          string
		req           *pb.UpdateAccountRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdateAccountResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {

				arg := db.AddAccountBalanceParams{
					Amount: amount,
					ID:     account.ID,
				}

				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(account, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "Invalid ID",
			req: &pb.UpdateAccountRequest{
				Id:     0,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "Invalid Amount",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: 1000000000,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "Account Doesn't Exist",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, db.ErrRecordNotFound)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "Depositor Cannot Update An Account",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					AddAccountBalance(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, other.Username, other.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			req: &pb.UpdateAccountRequest{
				Id:     account.ID,
				Amount: amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateAccountResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
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
			res, err := server.UpdateAccount(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
