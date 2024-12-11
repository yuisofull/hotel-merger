package paperflies

import (
	"cmp"
	"encoding/json"
	"hotel-merger/domain/hotels"
	"net/http"
)

type HotelModel struct {
	HotelId       string `json:"hotel_id"`
	DestinationId int    `json:"destination_id"`
	HotelName     string `json:"hotel_name"`
	Location      struct {
		Address string `json:"address"`
		Country string `json:"country"`
	} `json:"location"`
	Details   string `json:"details"`
	Amenities struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities"`
	Images struct {
		Rooms []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"rooms"`
		Site []struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		} `json:"site"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
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
		hotelMap[hotel.HotelId] = hotel
	}
	return &supplier{
		Hotels:   hotelMap,
		EndPoint: endPoint,
	}, nil
}

func (s *supplier) FillHotel(hotel *hotels.Hotel) {
	if paperfliesHotel, ok := s.Hotels[hotel.Id]; ok {
		hotel.DestinationId = cmp.Or(hotel.DestinationId, paperfliesHotel.DestinationId)
		hotel.Name = cmp.Or(hotel.Name, paperfliesHotel.HotelName)
		hotel.Location.Address = cmp.Or(hotel.Location.Address, paperfliesHotel.Location.Address)
		hotel.Location.Country = cmp.Or(hotel.Location.Country, paperfliesHotel.Location.Country)
		hotel.Description = cmp.Or(hotel.Description, paperfliesHotel.Details)
		if len(hotel.Amenities.General) == 0 {
			hotel.Amenities.General = paperfliesHotel.Amenities.General
		}
		if len(hotel.Amenities.Room) == 0 {
			hotel.Amenities.Room = paperfliesHotel.Amenities.Room
		}
		if len(hotel.BookingConditions) == 0 {
			hotel.BookingConditions = paperfliesHotel.BookingConditions
		}
		if len(hotel.Images.Rooms) == 0 {
			for _, room := range paperfliesHotel.Images.Rooms {
				hotel.Images.Rooms = append(hotel.Images.Rooms, struct {
					Link        string `json:"link"`
					Description string `json:"description"`
				}{
					Link:        room.Link,
					Description: room.Caption,
				})
			}
		}
		if len(hotel.Images.Site) == 0 {
			for _, site := range paperfliesHotel.Images.Site {
				hotel.Images.Site = append(hotel.Images.Site, struct {
					Link        string `json:"link"`
					Description string `json:"description"`
				}{
					Link:        site.Link,
					Description: site.Caption,
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
