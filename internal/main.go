package main

import (
	"hotel-merger/adapters/api"
	"hotel-merger/app"
	"hotel-merger/ports"
)

func main() {
	hotelRepo := api.NewRepository()
	service := app.NewHotelService(hotelRepo)
	cli := ports.NewCLI(service)
	cli.Execute()
}
