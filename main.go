package main

import (
	"github.com/labstack/echo"
)

type Car struct {
	Nome  string
	Price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Nome: "Corsa", Price: 1000})
	cars = append(cars, Car{Nome: "Celta", Price: 2000})
	cars = append(cars, Car{Nome: "Sivic", Price: 10000})
}

func main() {
	createCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}
