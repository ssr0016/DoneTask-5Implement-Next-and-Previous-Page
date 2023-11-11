// Define Service
package main

import "context"

type Service interface {
	NextPage(ctx context.Context, bankID int64) (*Bank, error)
	PreviousPage(ctx context.Context, bankID int64) (*Bank, error)
}
