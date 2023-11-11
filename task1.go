// Functionality
package main

import "context"

func (s *service) NextPage(ctx context.Context, bankID int64) (*bank.Bank, error) {
	currentBank, err := s.store.getNextPageBankByID(ctx, bankID)
	if err != nil {
		return nil, bank.ErrorGettingNextPage
	}

	nextBank, err := s.store.Get(ctx, currentBank.ID)
	if err != nil {
		return nil, err
	}

	if nextBank == nil {
		return nil, bank.ErrorNoMoreNextPages
	}

	return nextBank, nil
}

func (s *service) PreviousPage(ctx context.Context, bankID int64) (*bank.Bank, error) {
	currentBank, err := s.store.getPreviousPageBankByID(ctx, bankID)
	if err != nil {
		return nil, bank.ErrorGettingNextPage
	}

	nextBank, err := s.store.Get(ctx, currentBank.ID)
	if err != nil {
		return nil, err
	}

	if nextBank == nil {
		return nil, bank.ErrorNoMoreNextPages
	}

	return nextBank, nil
}
