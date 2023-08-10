package model

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"

type Ship struct {
	ID            string `json:"id"`
	ShipNumber    string `json:"ship_number"`
	DepartureTime string `json:"departure_time"`
}

func (s *Ship) Options() (options []*selectfield.Option, Error error) {
	getList := []Ship{
		{
			ID:            "123",
			ShipNumber:    "T101",
			DepartureTime: "2023-09-01 10:00:00",
		},
	}
	for _, v := range getList {
		option := &selectfield.Option{
			Label: v.ShipNumber,
			Value: v.DepartureTime,
		}
		options = append(options, option)
	}
	return options, nil
}
