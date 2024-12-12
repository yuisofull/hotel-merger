package acme

import (
	"cmp"
	"encoding/json"
	"hotel-merger/common"
	"hotel-merger/domain/hotels"
	"net/http"
	"strings"
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
		if acmeHotel.Latitude != 0 {
			tmp := float64(acmeHotel.Latitude)
			hotel.Location.Lat = cmp.Or(hotel.Location.Lat, &tmp)
		}
		if acmeHotel.Longitude != 0 {
			tmp := float64(acmeHotel.Longitude)
			hotel.Location.Lng = cmp.Or(hotel.Location.Lng, &tmp)
		}
		if acmeHotel.PostalCode != "" {
			hotel.Location.Address = cmp.Or(hotel.Location.Address, acmeHotel.PostalCode)
		}
		if strings.Compare(hotel.Location.Address, acmeHotel.Address) == -1 {
			hotel.Location.Address = acmeHotel.Address
		}
		if strings.Compare(hotel.Location.City, acmeHotel.City) == -1 {
			hotel.Location.City = acmeHotel.City
		}
		if strings.Compare(hotel.Location.Country, acmeHotel.Country) == -1 {
			hotel.Location.Country = acmeHotel.Country
		}
		if strings.Compare(hotel.Description, acmeHotel.Description) == -1 {
			hotel.Description = acmeHotel.Description
		}
	}
}

func (s *supplier) GetIds() []string {
	var ids []string
	for id := range s.Hotels {
		ids = append(ids, id)
	}
	return ids
}
