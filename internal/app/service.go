package app

import (
	"context"
	"hotel-merger/domain/hotels"
)

type HotelService struct {
	hotelRepo hotels.Repository
}

func NewHotelService(hotelRepo hotels.Repository) *HotelService {
	return &HotelService{hotelRepo: hotelRepo}
}

func (s *HotelService) GetAllHotels(ctx context.Context) ([]hotels.Hotel, error) {
	return s.hotelRepo.GetHotelsWithFilters(ctx, hotels.Filter{})
}

func (s *HotelService) GetHotelsByDestinationIds(ctx context.Context, destinationIds []int) ([]hotels.Hotel, error) {
	return s.hotelRepo.GetHotelsWithFilters(ctx, hotels.Filter{DestinationIds: destinationIds})
}

func (s *HotelService) GetHotelsByHotelIds(ctx context.Context, hotelIds []string) ([]hotels.Hotel, error) {
	return s.hotelRepo.GetHotelsWithFilters(ctx, hotels.Filter{HotelIds: hotelIds})
}

func (s *HotelService) GetHotelsByDestinationIdsAndHotelIds(ctx context.Context, destinationIds []int, hotelIds []string) ([]hotels.Hotel, error) {
	return s.hotelRepo.GetHotelsWithFilters(ctx, hotels.Filter{DestinationIds: destinationIds, HotelIds: hotelIds})
}
