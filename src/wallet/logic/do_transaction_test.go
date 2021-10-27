package logic

import (
	"context"
	"fmt"
	"strings"
	"testing"

	mocks "github.com/dinislamdarkhan/simple-wallet/mocks/src/wallet/logic"
	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"
	"github.com/stretchr/testify/mock"
)

func TestDoTransaction(t *testing.T) {
	tests := []struct {
		name                           string
		wantError                      error
		request                        *domain.PostDoTransactionRequest
		wantResponse                   *domain.PostDoTransactionResponse
		returnGetWalletAmountCassandra float64
		returnCheckAmountExists        int
		returnCassandraError           error
	}{
		{
			name:      "Correct EUR Deposit",
			wantError: nil,
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "EUR",
				Amount:     1000,
				Type:       "deposit",
				TimePlaced: "24-JAN-20 10:27:44",
			},
			wantResponse:                   &domain.PostDoTransactionResponse{Message: respMessage + fmt.Sprintf("%.3f", 1010.5)},
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Correct USD Withdrawal",
			wantError: nil,
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "USD",
				Amount:     1000,
				Type:       "withdrawal",
				TimePlaced: "24-JAN-20 10:27:44",
			},
			wantResponse:                   &domain.PostDoTransactionResponse{Message: respMessage + fmt.Sprintf("%.3f", -989.5)},
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Incorrect Currency",
			wantError: errors.BadRequest.SetMessage(currencyError),
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "KZT",
				Amount:     1000,
				Type:       "deposit",
				TimePlaced: "24-JAN-20 10:27:44",
			},
			wantResponse:                   nil,
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Incorrect Type",
			wantError: errors.BadRequest.SetMessage(operationTypeError),
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "EUR",
				Amount:     1000,
				Type:       "ERROR",
				TimePlaced: "24-JAN-20 10:27:44",
			},
			wantResponse:                   nil,
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Error Basic Structure Validation",
			wantError: errors.BadRequest,
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "EUR",
				Amount:     -2,
				Type:       "deposit",
				TimePlaced: "24-JAN-20 10:27:44",
			},
			wantResponse:                   nil,
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
		{
			name:      "Time Formatting Error",
			wantError: errors.DeserializeError.SetMessage("Error at formatting time"),
			request: &domain.PostDoTransactionRequest{
				UserID:     "134256",
				Currency:   "EUR",
				Amount:     1000,
				Type:       "deposit",
				TimePlaced: "error",
			},
			wantResponse:                   nil,
			returnGetWalletAmountCassandra: 10.5,
			returnCheckAmountExists:        1,
			returnCassandraError:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m := &mocks.DoTransactionRepository{}
			m.On("CheckAmountExistsCassandra", ctx, tt.request.Currency, tt.request.UserID).Return(tt.returnCheckAmountExists, tt.returnCassandraError)
			m.On("GetWalletAmountCassandra", ctx, tt.request.Currency, tt.request.UserID).Return(tt.returnGetWalletAmountCassandra, tt.returnCassandraError)
			m.On("UpdateWalletAmountCassandra", ctx, tt.request.Currency, tt.request.UserID, mock.Anything, mock.Anything).Return(tt.returnCassandraError)

			response, err := DoTransaction(ctx, m, tt.request)

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

			if response.Message != tt.wantResponse.Message {
				t.Errorf("want response %s, but got %s", tt.wantResponse, response)
				return
			}
		})
	}
}
