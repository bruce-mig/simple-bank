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

func TestListAccountAPI(t *testing.T) {
	user, _ := randomUser(t, util.DepositorRole)
	other, _ := randomUser(t, util.DepositorRole)
	// banker, _ := randomUser(t, util.BankerRole)

	n := 5
	accounts := make([]db.Account, n)
	for i := 0; i < n; i++ {
		accounts[i] = randomAccount(user.Username)
	}

	testCases := []struct {
		name          string
		req           *pb.ListAccountsRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.ListAccountsResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.ListAccountsRequest{
				PageId:   1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListAccountsParams{
					Owner:  user.Username,
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(accounts, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
				accArr := res.GetAccounts()
				resp := convertAccounts(accounts)
				require.Equal(t, accArr, resp)
			},
		},
		// {
		// 	name: "BankerCanGetAccountForUser",
		// 	req: &pb.ListAccountsRequest{
		// 		PageId:   1,
		// 		PageSize: int32(n),
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		arg := db.ListAccountsParams{
		// 			Owner:  user.Username,
		// 			Limit:  int32(n),
		// 			Offset: 0,
		// 		}

		// 		store.EXPECT().
		// 			ListAccounts(gomock.Any(), gomock.Eq(arg)).
		// 			Times(1).
		// 			Return(accounts, nil)
		// 	},
		// 	buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
		// 		return newContextWithBearerToken(t, tokenMaker, banker.Username, banker.Role, time.Minute)
		// 	},
		// 	checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
		// 		require.NoError(t, err)
		// 		require.NotNil(t, res)
		// 		st, ok := status.FromError(err)
		// 		require.True(t, ok)
		// 		require.Equal(t, codes.OK, st.Code())
		// 		accArr := res.GetAccounts()
		// 		resp := convertAccounts(accounts)
		// 		require.Equal(t, accArr, resp)
		// 	},
		// },
		{
			name: "NoAuthorization",
			req: &pb.ListAccountsRequest{
				PageId:   1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.ListAccountsRequest{
				PageId:   1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Account{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InvalidPageID",
			req: &pb.ListAccountsRequest{
				PageId:   -1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InvalidPageSize",
			req: &pb.ListAccountsRequest{
				PageId:   1,
				PageSize: int32(100000),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "OtherDepositorCannotListAccountsBelongingToOtherUser",
			req: &pb.ListAccountsRequest{
				PageId:   1,
				PageSize: int32(n),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(1).
					Return(accounts, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, other.Username, other.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.PermissionDenied, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Username, user.Role, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.ListAccountsResponse, err error) {
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			// build stubs
			tc.buildStubs(store)

			// start test server and send request
			server := newTestServer(t, store, nil)

			ctx := tc.buildContext(t, server.tokenMaker)
			res, err := server.ListAccounts(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
