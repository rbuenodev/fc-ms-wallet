package createtransaction

import (
	"context"
	"testing"

	"github.com/rbueno/fc-ms-wallet/internal/entity"
	"github.com/rbueno/fc-ms-wallet/internal/event"
	"github.com/rbueno/fc-ms-wallet/internal/usecase/mocks"
	"github.com/rbueno/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "client1@com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "client2@com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	output, err := uc.Execute(ctx, inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
