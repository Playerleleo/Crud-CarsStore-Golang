package main

import (
	"github.com/labstack/echo"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Corsa", Price: 1000})
	cars = append(cars, Car{Name: "Celta", Price: 2000})
	cars = append(cars, Car{Name: "Sivic", Price: 10000})
}

func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, car)
}
