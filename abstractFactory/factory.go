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

type AlchemicalStationFactory interface {
	CreateHerbStation() AlchemicalStation
	CreateCrystalStation() AlchemicalStation
}

type AlchemicalStationProducer struct{}

func (asp *AlchemicalStationProducer) CreateHerbStation() AlchemicalStation {
	return &HerbStation{}
}

func (asp *AlchemicalStationProducer) CreateCrystalStation() AlchemicalStation {
	return &CrystalStation{}
}

func main() {
	producer := &AlchemicalStationProducer{}

	herbStation := producer.CreateHerbStation()
	fmt.Println(herbStation.ProducePotion())

	crystalStation := producer.CreateCrystalStation()
	fmt.Println(crystalStation.ProducePotion())
}
