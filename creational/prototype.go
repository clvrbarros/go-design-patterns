package creational

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}
const (
	White = 1
	Black = 2
	Blue  = 3
)
func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}
type ShirtsCache struct {}
func (s *ShirtsCache)GetClone (m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}
type ShirtColor byte
type Shirt struct {
	Price float32
	SKU string
	Color ShirtColor
}
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color " +
		"id %d that costs %f\n", s.SKU, s.Color, s.Price)
}
var whitePrototype *Shirt = &Shirt{
	15.00,
	"empty",
	White,
}
var blackPrototype *Shirt = &Shirt{
	16.00,
	"empty",
	Black,
}
var bluePrototype *Shirt = &Shirt{
	17.00,
	"empty",
	Blue,
}
func (i *Shirt) GetPrice() float32 {
	return i.Price
}

