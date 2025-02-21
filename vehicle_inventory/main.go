package main

import (
	"errors"
	"fmt"
)

type VehicleType interface {
	ShowDetails() string
}

type Vehicle struct {
	make string
	model string
	year string
}

type Car struct {
	Vehicle
	doors int
	seats int
}

func (c Car) ShowDetails() string {
	return fmt.Sprintf("Car: %s %s %s Doors: %d Seats: %d", c.make, c.model, c.year, c.doors, c.seats)
}

type Truck struct {
	Vehicle
	max_load_kg float32
}

func (t Truck) ShowDetails() string  {
	return fmt.Sprintf("Truck: %s %s %s Load: %.2f Kg", t.make, t.model, t.year, t.max_load_kg)
}

type Inventory struct {
	vehicles []VehicleType
}

func (i *Inventory) AddVehicle(v VehicleType){
	i.vehicles = append(i.vehicles, v)
}

func (i *Inventory) RemoveVehicle(index int) error{
	if index < 0 || index >= len(i.vehicles) {
		return errors.New("invalid index")
	}
	i.vehicles = append(i.vehicles[:index], i.vehicles[index+1:]...)
	return nil
}

func (i *Inventory) ShowInventory() {
	for i, v := range i.vehicles {
		fmt.Printf("%d %s\n", i+1, v.ShowDetails())
	}
}


func main() {
	inventry := Inventory{}

	car1 := Car{Vehicle{"Skoda","Fabia","2023"}, 5, 5}
	car2 := Car{Vehicle{"Tesla","Model Y","2025"}, 5, 5}
	truck1 := Truck{Vehicle{"Mercedes","SomeModel", "2000"},1200.6}

	inventry.AddVehicle(car1)
	inventry.AddVehicle(car2)
	inventry.AddVehicle(truck1)

	fmt.Println("Current Inventory:")
	inventry.ShowInventory()

	inventry.RemoveVehicle(1)

	fmt.Println("Current Inventory:")
	inventry.ShowInventory()
}
