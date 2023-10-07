package main

import "fmt"

type AlchemicalStation interface {
	ProducePotion() string
}

type HerbStation struct{}

func (h *HerbStation) ProducePotion() string {
	return "Producing Healing Potion using Herb Station!"
}

type CrystalStation struct{}

func (c *CrystalStation) ProducePotion() string {
	return "Producing Defence Potion using Crystal Station!"
}

type AlchemicalStationDecorator interface {
	ProducePotion() string
}

type EnhancedStation struct {
	Station AlchemicalStation
}

func (e *EnhancedStation) ProducePotion() string {
	return e.Station.ProducePotion() + " Extended duration!"
}

func main() {
	herbStation := &HerbStation{}
	fmt.Println(herbStation.ProducePotion())

	crystalStation := &CrystalStation{}
	fmt.Println(crystalStation.ProducePotion())

	enhancedHerbStation := &EnhancedStation{Station: herbStation}
	fmt.Println(enhancedHerbStation.ProducePotion())

	enhancedCrystalStation := &EnhancedStation{Station: crystalStation}
	fmt.Println(enhancedCrystalStation.ProducePotion())
}
