package acme

import (
	"cmp"
	"encoding/json"
	"hotel-merger/common"
	"hotel-merger/domain/hotels"
	"net/http"
)

type HotelModel struct {
	Id            string                `json:"Id"`
	DestinationId int                   `json:"DestinationId"`
	Name          string                `json:"Name"`
	Latitude      common.DynamicFloat64 `json:"Latitude"`
	Longitude     common.DynamicFloat64 `json:"Longitude"`
	Address       string                `json:"Address"`
	City          string                `json:"City"`
	Country       string                `json:"Country"`
	PostalCode    string                `json:"PostalCode"`
	Description   string                `json:"Description"`
	Facilities    []string              `json:"Facilities"`
}

type supplier struct {
	Hotels   map[string]HotelModel
	EndPoint string
}

func NewSupplier(endPoint string) (*supplier, error) {
	resp, err := http.Get(endPoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var hotels []HotelModel
	if err := json.NewDecoder(resp.Body).Decode(&hotels); err != nil {
		return nil, err
	}
	hotelMap := make(map[string]HotelModel)
	for _, hotel := range hotels {
		hotelMap[hotel.Id] = hotel
	}
	return &supplier{
		Hotels:   hotelMap,
		EndPoint: endPoint,
	}, nil
}

func (s *supplier) FillHotel(hotel *hotels.Hotel) {
	if acmeHotel, ok := s.Hotels[hotel.Id]; ok {
		hotel.Id = cmp.Or(hotel.Id, acmeHotel.Id)
		hotel.DestinationId = cmp.Or(hotel.DestinationId, acmeHotel.DestinationId)
		hotel.Name = cmp.Or(hotel.Name, acmeHotel.Name)
		hotel.Location.Lat = cmp.Or(hotel.Location.Lat, float64(acmeHotel.Latitude))
		hotel.Location.Lng = cmp.Or(hotel.Location.Lng, float64(acmeHotel.Longitude))
		hotel.Location.Address = cmp.Or(hotel.Location.Address, acmeHotel.Address)
		hotel.Location.City = cmp.Or(hotel.Location.City, acmeHotel.City)
		hotel.Location.Country = cmp.Or(hotel.Location.Country, acmeHotel.Country)
		hotel.Description = cmp.Or(hotel.Description, acmeHotel.Description)
	}
}

func (s *supplier) GetIds() []string {
	var ids []string
	for id := range s.Hotels {
		ids = append(ids, id)
	}
	return ids
}
