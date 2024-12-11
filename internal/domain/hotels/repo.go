package hotels

import "context"

type Repository interface {
	GetHotelsWithFilters(ctx context.Context, filter Filter) ([]Hotel, error)
}

type Filter struct {
	DestinationIds []int
	HotelIds       []string
}
