package patagonia

import (
	"cmp"
	"encoding/json"
	"hotel-merger/common"
	"hotel-merger/domain/hotels"
	"net/http"
	"strings"
)

type HotelModel struct {
	Id          string                `json:"id"`
	Destination int                   `json:"destination"`
	Name        string                `json:"name"`
	Lat         common.DynamicFloat64 `json:"lat"`
	Lng         common.DynamicFloat64 `json:"lng"`
	Address     string                `json:"address"`
	Info        string                `json:"info"`
	Amenities   []string              `json:"amenities"`
	Images      struct {
		Rooms []struct {
			Url         string `json:"url"`
			Description string `json:"description"`
		} `json:"rooms"`
		Amenities []struct {
			Url         string `json:"url"`
			Description string `json:"description"`
		} `json:"amenities"`
	} `json:"images"`
}

type supplier struct {
	Hotels   map[string]HotelModel
	EndPoint string
}

func (s *supplier) FillHotel(hotel *hotels.Hotel) {
	if patagoniaHotel, ok := s.Hotels[hotel.Id]; ok {
		hotel.DestinationId = cmp.Or(hotel.DestinationId, patagoniaHotel.Destination)
		hotel.Name = cmp.Or(hotel.Name, patagoniaHotel.Name)
		if patagoniaHotel.Lat != 0 {
			tmp := float64(patagoniaHotel.Lat)
			hotel.Location.Lat = cmp.Or(hotel.Location.Lat, &tmp)
		}
		if patagoniaHotel.Lng != 0 {
			tmp := float64(patagoniaHotel.Lng)
			hotel.Location.Lng = cmp.Or(hotel.Location.Lng, &tmp)
		}
		if strings.Compare(hotel.Location.Address, patagoniaHotel.Address) == -1 {
			hotel.Location.Address = patagoniaHotel.Address
		}
		hotel.Description = cmp.Or(hotel.Description, patagoniaHotel.Info)
		if len(hotel.Amenities.General) == 0 {
			hotel.Amenities.General = patagoniaHotel.Amenities
		}
		if len(hotel.Images.Rooms) == 0 {
			for _, room := range patagoniaHotel.Images.Rooms {
				hotel.Images.Rooms = append(hotel.Images.Rooms, struct {
					Link        string `json:"link"`
					Description string `json:"description"`
				}{
					Link:        room.Url,
					Description: room.Description,
				})
			}
		}
		if len(hotel.Images.Amenities) == 0 {
			for _, amenity := range patagoniaHotel.Images.Amenities {
				hotel.Images.Amenities = append(hotel.Images.Amenities, struct {
					Link        string `json:"link"`
					Description string `json:"description"`
				}{
					Link:        amenity.Url,
					Description: amenity.Description,
				})
			}
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
