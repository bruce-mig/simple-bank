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

func randomTransfer(fromAccountID int64, toAccountID int64, amount int64) db.Transfer {
	return db.Transfer{
		ID:            util.RandomInt(1, 1000),
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
}

func TestGetTransferAPI(t *testing.T) {
	amount := int64(10)

	user1, _ := randomUser(t, util.DepositorRole)
	user2, _ := randomUser(t, util.DepositorRole)
	user3, _ := randomUser(t, util.DepositorRole)
	user4, _ := randomUser(t, util.BankerRole)

	account1 := randomAccount(user1.Username)
	account2 := randomAccount(user2.Username)
	account3 := randomAccount(user3.Username)

	account1.Currency = util.USD
	account2.Currency = util.USD
	account3.Currency = util.ZAR

	transfer1 := randomTransfer(account1.ID, account2.ID, amount)
	transfer2 := randomTransfer(account2.ID, account1.ID, amount)

	testCases := []struct {
		name          string
		req           *pb.GetTransferRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.GetTransferResponse, err error)
	}{
		{
			name: "OKasSender",
			req: &pb.GetTransferRequest{
				Id: transfer1.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(1).
					Return(transfer1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(1).
					Return(account1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(1).
					Return(account2, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
			},
		},
		{
			name: "OKasReceiver",
			req: &pb.GetTransferRequest{
				Id: transfer2.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer2.ID)).
					Times(1).
					Return(transfer1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(1).
					Return(account1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(1).
					Return(account2, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
			},
		},
		{
			name: "BankerCanGetTransfer",
			req: &pb.GetTransferRequest{
				Id: transfer1.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(1).
					Return(transfer1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(1).
					Return(account1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(1).
					Return(account2, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user4.Username, user4.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.OK, st.Code())
			},
		},
		{
			name: "OtherDepositorCannotGetTransferBelongingToOtherUser",
			req: &pb.GetTransferRequest{
				Id: transfer1.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(1).
					Return(transfer1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(1).
					Return(account1, nil)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(1).
					Return(account2, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user3.Username, user3.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.PermissionDenied, st.Code())
			},
		},
		{
			name: "NoAuthorization",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.GetTransferRequest{
				Id: transfer1.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(1).
					Return(db.Transfer{}, sql.ErrConnDone)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "NotFound",
			req: &pb.GetTransferRequest{
				Id: transfer1.ID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(1).
					Return(db.Transfer{}, db.ErrRecordNotFound)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "InvalidID",
			req: &pb.GetTransferRequest{
				Id: -1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "ExpiredToken",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTransfer(gomock.Any(), gomock.Eq(transfer1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Times(0)
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user1.Username, user1.Role, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetTransferResponse, err error) {
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
			res, err := server.GetTransfer(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
