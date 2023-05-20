package main


import "fmt"


// Declaring Book Struct
type Book struct {
	isbn string
	title string
	price float64
}

// Declare Product Struct
type Product struct {
	name string
	unitPrice float32
	quantity int
}

func main() {

	// Initialize Book Struct 
	book := Book{"12345678", "Art of Programming", 456.78}

	// Initialize Book Struct 
	book1 := Book{ price: 456.78, isbn: "1234567890", title: "Art of Programming"}
	fmt.Println(book1.title)

	//	Declare empty Boom struct 
	book2 := Book{}
	fmt.Println(book2.title)

	//	Declare and initialize new Product
	p1 := Product{"T107", 56.78, 12}

	//  Pointers in Structures, declaring new structure by referencing values from another struct via Pointers
	ptr2p1 := &p1

	//	Accessing Product fields (Asterisk and (.) operator are used for dereferencing)
	fmt.Println((*ptr2p1).name)
	fmt.Println(ptr2p1.name)
	fmt.Println(ptr2p1)

	//	Accessing structure fields
	fmt.Println(book.price)

	//	Creating struct pointers using new keyword | 
	ptr := new(Product)
	fmt.Println(ptr)
}