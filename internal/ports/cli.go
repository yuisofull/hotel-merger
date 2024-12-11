package ports

import (
	"context"
	"encoding/json"
	"hotel-merger/app"
	"log"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	service *app.HotelService
}

func NewCLI(service *app.HotelService) *CLI {
	return &CLI{service: service}
}

func (c *CLI) Execute() {
	hotelIds, destinationIds, err := parseArgs()
	if err != nil {
		log.Fatalf("failed to parse args: %v", err)
	}
	if hotelIds == nil && destinationIds == nil {
		hotels, err := c.service.GetAllHotels(context.Background())
		if err != nil {
			log.Fatalf("failed to get all hotels: %v", err)
		}
		jsonData, _ := json.MarshalIndent(hotels, "", "  ")
		log.Printf("%s", jsonData)
		return
	}
	if hotelIds != nil && destinationIds != nil {
		hotels, err := c.service.GetHotelsByDestinationIdsAndHotelIds(context.Background(), destinationIds, hotelIds)
		if err != nil {
			log.Fatalf("failed to get hotels by destination ids and hotel ids: %v", err)
		}
		jsonData, _ := json.MarshalIndent(hotels, "", "  ")
		log.Printf("%s", jsonData)
		return
	}
	if hotelIds != nil {
		hotels, err := c.service.GetHotelsByHotelIds(context.Background(), hotelIds)
		if err != nil {
			log.Fatalf("failed to get hotels by hotel ids: %v", err)
		}
		jsonData, _ := json.MarshalIndent(hotels, "", "  ")
		log.Printf("%s", jsonData)
		return
	}
	hotels, err := c.service.GetHotelsByDestinationIds(context.Background(), destinationIds)
	if err != nil {
		log.Fatalf("failed to get hotels by destination ids: %v", err)
	}
	jsonData, _ := json.MarshalIndent(hotels, "", "  ")
	log.Printf("%s", jsonData)
	return

}

func parseArgs() ([]string, []int, error) {
	if len(os.Args) < 2 {
		return nil, nil, nil
	}
	hotelIds, err := extractHotelIds(os.Args[1])
	if err != nil {
		return nil, nil, err
	}
	if len(os.Args) < 3 {
		return hotelIds, nil, nil
	}
	destinationIds, err := extractDestinationIds(os.Args[2])
	if err != nil {
		return nil, nil, err
	}
	return hotelIds, destinationIds, nil
}

func extractHotelIds(arg string) ([]string, error) {
	if arg == "none" || arg == "" {
		return nil, nil
	}
	return strings.Split(arg, ","), nil
}

func extractDestinationIds(arg string) ([]int, error) {
	if arg == "none" || arg == "" {
		return nil, nil
	}
	var destinationIds []int
	for _, id := range strings.Split(arg, ",") {
		tmp, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("failed to parse destination id: %v", err)
		}
		destinationIds = append(destinationIds, tmp)
	}
	return destinationIds, nil
}
