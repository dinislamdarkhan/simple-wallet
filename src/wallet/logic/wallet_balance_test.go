package logic

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	mocks "github.com/dinislamdarkhan/simple-wallet/mocks/src/wallet/logic"
	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"
	"github.com/stretchr/testify/mock"
)

func TestWalletBalance(t *testing.T) {
	tests := []struct {
		name                           string
		wantError                      error
		request                        *domain.GetWalletBalanceRequest
		wantResponse                   *domain.GetWalletBalanceResponse
		returnGetWalletAmountCassandra float64
		returnTime                     time.Time
		returnCheckAmountExists        int
		returnCassandraError           error
	}{
		{
			name:      "Correct UserID",
			wantError: nil,
			request: &domain.GetWalletBalanceRequest{
				UserID: "134256",
			},
			wantResponse: &domain.GetWalletBalanceResponse{Wallet: []domain.Account{
				{
					Currency:    "EUR",
					Amount:      10.5,
					UpdatedTime: time.Time{}.String(),
				},
				{
					Currency:    "USD",
					Amount:      10.5,
					UpdatedTime: time.Time{}.String(),
				},
			}},
			returnGetWalletAmountCassandra: 10.5,
			returnTime:                     time.Time{},
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "InCorrect UserID",
			wantError: errors.BadRequest,
			request: &domain.GetWalletBalanceRequest{
				UserID: "",
			},
			wantResponse:                   nil,
			returnGetWalletAmountCassandra: 10.5,
			returnTime:                     time.Time{},
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Not Found UserID",
			wantError: nil,
			request: &domain.GetWalletBalanceRequest{
				UserID: "123",
			},
			wantResponse:                   &domain.GetWalletBalanceResponse{Wallet: nil},
			returnGetWalletAmountCassandra: 10.5,
			returnTime:                     time.Time{},
			returnCheckAmountExists:        0,
			returnCassandraError:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m := &mocks.WalletBalanceRepository{}
			m.On("CheckAmountExistsCassandra", ctx, mock.Anything, tt.request.UserID).Return(tt.returnCheckAmountExists, tt.returnCassandraError)
			m.On("GetWalletAmountAndTimeCassandra", ctx, mock.Anything, tt.request.UserID).Return(tt.returnGetWalletAmountCassandra, time.Time{}, tt.returnCassandraError)

			resp, err := WalletBalance(ctx, m, tt.request)

			if tt.wantError != nil {
				if err == nil {
					t.Errorf("want err: %s, but got err=nil", tt.wantError)
					return
				}

				if !strings.Contains(err.Error(), tt.wantError.Error()) {
					t.Errorf("unexpected error. Err must contain %s. Got err: %s", tt.wantError.Error(), err.Error())
				}

				return
			} else if err != nil {
				t.Errorf("unexpected error. got %s, but want err=nil", err)
				return
			}

			if !reflect.DeepEqual(tt.wantResponse.Wallet, resp.Wallet) {
				respBytes, err := json.Marshal(resp)
				if err != nil {
					t.Error("json deserialization bug ", err)
				}
				wantRespBytes, err := json.Marshal(tt.wantResponse)
				if err != nil {
					t.Error("json deserialization bug ", err)
				}
				t.Errorf("want response %s, but got %s", string(respBytes), string(wantRespBytes))
				return
			}
		})
	}
}
