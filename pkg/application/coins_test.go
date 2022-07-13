package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mcoins "klever.io/interview/mocks/coins"
	"klever.io/interview/pkg/domain"
)

func TestCoinApplication_FindByID(t *testing.T) {
	type args struct {
		coinID uint
	}
	tests := []struct {
		name          string
		args          args
		mockedCoin    *domain.Coin
		mockedError   error
		expected      *domain.Coin
		expectedError error
		wantError     bool
	}{
		{
			name: "succes find coin by id",
			args: args{
				coinID: 1,
			},
			mockedCoin: &domain.Coin{
				ID:          pointer.ToUint(1),
				Description: "Coin 1",
				Short:       "C1",
				Votes:       0,
			},
			expected: &domain.Coin{
				ID:          pointer.ToUint(1),
				Description: "Coin 1",
				Short:       "C1",
				Votes:       0,
			},
			wantError: false,
		},
		{
			name: "could not get coin by id",
			args: args{
				coinID: 1,
			},
			mockedError:   fmt.Errorf("generic error"),
			mockedCoin:    &domain.Coin{},
			expected:      &domain.Coin{},
			expectedError: fmt.Errorf("generic error"),
			wantError:     true,
		},
	}
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepCoin := mcoins.NewMockRepository(mockCtrl)

			mockRepCoin.EXPECT().
				FindByID(ctx, tt.args.coinID).
				Return(tt.mockedCoin, tt.mockedError)

			a := CoinApplication{
				coinRepository: mockRepCoin,
			}

			got, err := a.FindByID(ctx, tt.args.coinID)

			if tt.wantError {
				assert.Error(t, err, "Error shouldn't be nil")
				assert.EqualError(t, err, tt.expectedError.Error(), "Divergent erros")
			} else {
				assert.NoError(t, err, "Error should be nil")
				assert.NotNil(t, got, "Coin should not be nil ")
				assert.EqualValues(t, tt.expected, got, "Divergent values")
			}
		})
	}
}

func TestCoinApplication_UpdateCoin(t *testing.T) {
	type args struct {
		domainCoin *domain.Coin
	}
	tests := []struct {
		name                string
		args                args
		mockedUpdateError   error
		mockedFindByIDError error
		mockedCoin          *domain.Coin

		expected      *domain.Coin
		expectedError error
		wantErr       bool
	}{
		{
			name: "success coin update",
			args: args{
				domainCoin: &domain.Coin{
					ID:          pointer.ToUint(1),
					Description: "Coin 1",
					Short:       "CO1",
					Votes:       0,
				},
			},
			mockedUpdateError:   nil,
			mockedFindByIDError: nil,
			mockedCoin: &domain.Coin{
				ID:          pointer.ToUint(1),
				Description: "Coin 1",
				Short:       "CO1",
				Votes:       0,
			},
			expected: &domain.Coin{
				ID:          pointer.ToUint(1),
				Description: "Coin 1",
				Short:       "CO1",
				Votes:       0,
			},
			expectedError: nil,
			wantErr:       false,
		},
		{
			name: "coin update error",
			args: args{
				domainCoin: &domain.Coin{
					ID:          pointer.ToUint(1),
					Description: "Coin 1",
					Short:       "CO1",
					Votes:       0,
				},
			},
			mockedUpdateError:   fmt.Errorf("update generic error"),
			mockedFindByIDError: nil,
			mockedCoin:          &domain.Coin{},
			expected:            &domain.Coin{},
			expectedError:       fmt.Errorf("update generic error"),
			wantErr:             true,
		},
		{
			name: "find error after update coin",
			args: args{
				domainCoin: &domain.Coin{
					ID:          pointer.ToUint(1),
					Description: "Coin 1",
					Short:       "CO1",
					Votes:       0,
				},
			},
			mockedFindByIDError: fmt.Errorf("coin not found"),
			mockedCoin:          &domain.Coin{},
			expected:            &domain.Coin{},
			expectedError:       fmt.Errorf("coin not found"),
			wantErr:             true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepCoin := mcoins.NewMockRepository(mockCtrl)

			mockRepCoin.EXPECT().
				Update(ctx, tt.args.domainCoin).
				Return(tt.mockedUpdateError)

			mockRepCoin.EXPECT().
				FindByID(ctx, pointer.GetUint(tt.args.domainCoin.ID)).
				MinTimes(0).
				MaxTimes(1).
				Return(tt.mockedCoin, tt.mockedFindByIDError)

			a := &CoinApplication{
				coinRepository: mockRepCoin,
			}

			got, err := a.UpdateCoin(ctx, tt.args.domainCoin)
			if tt.wantErr {
				assert.Error(t, err, "The error shouldn't nil")
				assert.EqualError(
					t,
					err,
					tt.expectedError.Error(),
					"Error different that expected",
				)
			} else {
				assert.NoError(t, err, "The error should be nil")
				assert.NotNil(t, got, "The coin should be nil")
				assert.EqualValues(
					t,
					tt.expected,
					got,
					"The coin values different that expected",
				)
			}
		})
	}
}
