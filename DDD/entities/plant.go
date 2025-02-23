package entities

import (
	"errors"
	"time"
)

type Plant struct {
	ID   int    `json:"id"`
	Name PlantName `json:"name"`
	WateringDate *time.Time `json:"watering_date"`
}

type PlantName string

func NewPlantName(name string)(PlantName, error){
	if name == ""{
		return "", errors.New("name is empty")
	}
	return PlantName(name), nil
}
func NewPlant(name string, args ...func(*Plant))*Plant{
	PlantName, err  := NewPlantName(name)

	if err != nil {
		return &Plant{}
	}

	p :=  Plant{
		ID:  0,
		Name: PlantName,
		WateringDate: nil,
	}

	for _,arg := range args{
		arg(&p)
	}

	return &p
}

func withWateringDate(d *time.Time) func(*Plant){
	return func(p *Plant){
		p.WateringDate = d
	}
}

func (p *Plant) UpdateWatering() {
	d := time.Now()
	p.WateringDate = &d
}