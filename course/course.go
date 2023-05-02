package course

import "fmt"

// * Identificadores exportados son todos aquellos que comienzan con una letra mayúscula, por ejemplo:
// * Estructuras: Course
// * Propiedades: Name, Price, IsFree, UserIDs, Classes
// * Método: PrintClasses

// * Identificadores no exportados son todos aquellos que comienzan con una letra minúscula, por ejemplo:
// * Método: changePrice


type Course struct {
	Name string
	Price float64
	IsFree bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) PrintClasses () {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c *Course) changePrice (price float64) {
	c.Price = price
}