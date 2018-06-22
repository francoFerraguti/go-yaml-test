package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TaxType string

const (
	NoVat               TaxType = "NONE"
	VatOnExpenses20     TaxType = "INPUT2"
	ZeroRatedECServices TaxType = "ECZROUTPUTSERVICES"
)

type Consultant struct {
	Name         string
	Enabled      bool
	XeroName     string
	PricePerHour float64
	TaxType      TaxType
}

type Storage struct {
	Consultants map[string]Consultant
}

func main() {
	filename := "consultants.yaml"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	storage := Storage{}

	err = yaml.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println(err.Error())
	}

	storage.Consultants["franco"] = Consultant{}
	storage.Consultants["gonzalo"] = Consultant{}

	newYaml, err := yaml.Marshal(storage)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	err = ioutil.WriteFile("consultants.yaml", newYaml, 0644)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println("Finished")
}
