package hotels

type Hotel struct {
	Id            string `json:"id"`
	DestinationId int    `json:"destination_id"`
	Name          string `json:"name"`
	Location      struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
		City    string  `json:"city"`
		Country string  `json:"country"`
	} `json:"location"`
	Description string `json:"description"`
	Amenities   struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities"`
	Images struct {
		Rooms []struct {
			Link        string `json:"link"`
			Description string `json:"description"`
		} `json:"rooms"`
		Site []struct {
			Link        string `json:"link"`
			Description string `json:"description"`
		} `json:"site"`
		Amenities []struct {
			Link        string `json:"link"`
			Description string `json:"description"`
		} `json:"amenities"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
}
