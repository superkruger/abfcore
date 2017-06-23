package model

type Unit string

type Emission struct {
	Impact string
	Amount float64
}

type LCA struct {
	Name string
	Unit Unit
	Emissions []Emission
}

type Weight struct {
	Grams float64 `json:"grams"`
}

type Distance struct {
	Km float64 `json:"km"`
}

type Joule struct {
	Joule float64 `json:"joule"`
}

type Material struct {
	Name string
	Weight Weight
	TransportAlloc TransportAlloc

	Lca LCA
}

type Transport struct {
	Name string
	Distance Distance

	Lca LCA
}

type Energy struct {
	Name string
	Joule Joule

	Lca LCA
}

type Process struct {
	Name string
	MaterialAlloc MaterialAlloc
	ProductAlloc ProductAlloc
	EnergyAlloc EnergyAlloc
}

type Product struct {
	Name string
	Produced int
	MaterialAlloc MaterialAlloc
}

type MaterialAlloc struct {
	Name string
	Materials []Material
	Products []Product
}

type ProductAlloc struct {
	Name string
	Products []Product
}

type EnergyAlloc struct {
	Name string
	Energies []Energy
}

type TransportAlloc struct {
	Name string
	Transports []Transport
}

func (product *Product) getWeight() (float64) {
	// TODO
	return 0
}

func (product *Product) getMaterialEmissions() ([]Emission) {
	// TODO
	return nil
}

func (product *Product) getTransportEmissions() ([]Emission) {
	// TODO
	return nil
}

func (product *Product) getEnergyEmissions() ([]Emission) {
	// TODO
	return nil
}