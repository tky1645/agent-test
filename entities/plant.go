package entities

import (
	"errors"
	"time"
	"github.com/google/uuid"
)

type Plant struct {
	ID          string     `json:"id"`
	Name        PlantName  `json:"name"`
	Description *string    `json:"description"`
	ImageURL    *string    `json:"image_url"`
	WateringDate *time.Time `json:"watering_date"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
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
		ID:          uuid.New().String(),
		Name:        PlantName,
		Description: nil,
		ImageURL:    nil,
		WateringDate: nil,
	}

	for _,arg := range args{
		arg(&p)
	}

	return &p
}


func (p *Plant) UpdateWatering() {
	d := time.Now()
	p.WateringDate = &d
}

func WithDescription(desc *string) func(*Plant){
	return func(p *Plant){
		p.Description = desc
	}
}

func WithImageURL(url *string) func(*Plant){
	return func(p *Plant){
		p.ImageURL = url
	}
}
