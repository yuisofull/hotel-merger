package api

import (
	"context"
	"hotel-merger/adapters/api/supplier/acme"
	"hotel-merger/adapters/api/supplier/paperflies"
	"hotel-merger/adapters/api/supplier/patagonia"
	"hotel-merger/domain/hotels"
)

type HotelSupplier interface {
	FillHotel(hotel *hotels.Hotel)
	GetIds() []string
}

type repository struct {
	suppliers []HotelSupplier
	Hotels    []hotels.Hotel
}

func (r repository) GetHotelsWithFilters(ctx context.Context, filter hotels.Filter) ([]hotels.Hotel, error) {
	var res []hotels.Hotel
	for _, hotel := range r.Hotels {
		if len(filter.HotelIds) > 0 && !contains(filter.HotelIds, hotel.Id) {
			continue
		}
		if len(filter.DestinationIds) > 0 && !contains(filter.DestinationIds, hotel.DestinationId) {
			continue
		}
		res = append(res, hotel)
	}
	return res, nil
}

func contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func NewRepository() *repository {
	acmeEndpoint := "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"
	paperfliesEndpoint := "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"
	patagoniaEndpoint := "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"
	var suppliers []HotelSupplier
	acmeSupplier, err := acme.NewSupplier(acmeEndpoint)
	if err != nil {
		panic(err)
	}
	suppliers = append(suppliers, acmeSupplier)
	paperfliesSupplier, err := paperflies.NewSupplier(paperfliesEndpoint)
	if err != nil {
		panic(err)
	}
	suppliers = append(suppliers, paperfliesSupplier)
	patagoniaSupplier, err := patagonia.NewSupplier(patagoniaEndpoint)
	if err != nil {
		panic(err)
	}
	suppliers = append(suppliers, patagoniaSupplier)
	hs := hotelsWithIds(mergeIds(
		acmeSupplier.GetIds(),
		paperfliesSupplier.GetIds(),
		patagoniaSupplier.GetIds(),
	))

	for i := range hs {
		for _, supplier := range suppliers {
			supplier.FillHotel(&hs[i])
		}
	}
	return &repository{
		suppliers: suppliers,
		Hotels:    hs,
	}
}

func mergeIds(idss ...[]string) []string {
	var existingIds = make(map[string]bool)
	var ids []string
	for _, idList := range idss {
		for _, id := range idList {
			if _, ok := existingIds[id]; !ok {
				ids = append(ids, id)
				existingIds[id] = true
			}
		}
	}
	return ids
}

func hotelsWithIds(ids []string) []hotels.Hotel {
	var res []hotels.Hotel
	for _, id := range ids {
		res = append(res, hotels.Hotel{Id: id})
	}
	return res
}
