package main

import (
	"github.com/CristFlores/2_encapsulacion/course"
)

func main () {
	Go := &course.Course{
		Name: "Go desde cero",
		Price: 12.34,
		IsFree: false,
		UserIDs: []uint{12,56,89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
	Go.changePrice(15.5)
}

