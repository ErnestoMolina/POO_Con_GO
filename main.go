package main

import (
	"library/animal"
)

func main() {
	// //encapsulamiento Mayuscula publico, Minuscula Privado
	// // Book := book.Book("Moby Dick", "Herman Melville", 300) // no se puede ejecutar porque los atributos son privados
	// // Crear un objeto utilizando un constructor
	// myBook := book.NewBook("Moby Dick", "Herman Melville", 300)
	// myBook.PrintInfo()

	// // Modificar un atributo de un objeto utilizando funcion Set
	// myBook.SetTitle("Mody Dick (Edicion especial)")
	// // Imprimiendo con Get dato modificado con Set
	// fmt.Printf("Nuevo nombre: %v\n", myBook.GetTitle())
	// // Book.PrintInfo() // objeto que no se puede ejecutar porque no se pudo crear correctamente.
	// // Creando un objeto con herencia
	// mytextBook := book.NewTextBook("Cien años de soledad", "Gabriel Garcia Marquez", 658, "PrintSecund", "Decimo")
	// mytextBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(mytextBook)

	// // Polimorfismo
	// // Creando un objeto sin constructor
	// miPerro := animal.Perro{Nombre: "Max"}
	// miGato := animal.Gato{Nombre: "Tom"}

	// animal.HacerSonido(&miPerro)
	// animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
