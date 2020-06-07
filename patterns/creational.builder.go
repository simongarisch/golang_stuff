package main

import "fmt"

type Builder interface {
	BuildFoundation() string
	BuildLevels() string
}

type FlatBuilder struct{}
type HouseBuilder struct{}

func (builder FlatBuilder) BuildFoundation() string {
	return "small"
}

func (builder FlatBuilder) BuildLevels() string {
	return "many"
}

func (builder HouseBuilder) BuildFoundation() string {
	return "large"
}

func (builder HouseBuilder) BuildLevels() string {
	return "one"
}

type EngineerDirector struct{}

func (engineer EngineerDirector) ConstructBuilding(builder Builder) string {
	foundation := builder.BuildFoundation()
	levels := builder.BuildLevels()
	return foundation + " foundation and " + levels + " levels"
}

func main() {
	engineer := EngineerDirector{}
	fmt.Println(engineer.ConstructBuilding(FlatBuilder{}))  // small foundation and many levels
	fmt.Println(engineer.ConstructBuilding(HouseBuilder{})) // large foundation and one levels
}
