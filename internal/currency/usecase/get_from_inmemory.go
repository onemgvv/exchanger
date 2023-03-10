package usecase

import (
	"context"
	"errors"
	"exchanger/internal/models"
	"fmt"
)

func (uc useCase) GetFromInMemory(ctx context.Context, params models.CurrencyParams) (float64, error) {
	rate, ok := uc.inMemory.Get(fmt.Sprint(params.CurrencyFrom, params.CurrencyTo))
	if !ok {
		return 0, errors.New("pair not found in inMemory storage")
	}

	return rate, nil
}
