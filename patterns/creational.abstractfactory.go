package main

import "fmt"

type Vehicle struct {
	make  string
	model string
	parts []string
}

type Factory interface {
	Build(make string, model string) string
}

type VehicleChassisFactory struct{}
type VehicleEngineFactory struct{}
type VehicleBodyFactory struct{}

func (factory VehicleChassisFactory) Build(make string, model string) string {
	return "Chassis for a " + make + " " + model
}

func (factory VehicleEngineFactory) Build(make string, model string) string {
	return "Engine for a " + make + " " + model
}

func (factory VehicleBodyFactory) Build(make string, model string) string {
	return "Body for a " + make + " " + model
}

type VehicleFactory struct {
	chassis_factory Factory
	engine_factory  Factory
	body_factory    Factory
}

func NewVehicleFactory() VehicleFactory {
	factory := VehicleFactory{}
	factory.chassis_factory = VehicleChassisFactory{}
	factory.engine_factory = VehicleEngineFactory{}
	factory.body_factory = VehicleBodyFactory{}
	return factory
}

func (factory VehicleFactory) Build(make string, model string) Vehicle {
	vehicle := Vehicle{}
	vehicle.make = make
	vehicle.model = model

	childFactories := []Factory{
		factory.chassis_factory,
		factory.engine_factory,
		factory.body_factory,
	}

	for _, childFactory := range childFactories {
		vehicle.parts = append(vehicle.parts, childFactory.Build(make, model))
	}

	return vehicle
}

func main() {
	factory := NewVehicleFactory()

	var car Vehicle
	car = factory.Build("Ford", "Mustang")
	fmt.Println(car) // {Ford Mustang [Chassis for a Ford Mustang Engine for a Ford Mustang Body for a Ford Mustang]}

	car = factory.Build("Nissan", "Xtrail")
	fmt.Println(car) // {Nissan Xtrail [Chassis for a Nissan Xtrail Engine for a Nissan Xtrail Body for a Nissan Xtrail]}
}
