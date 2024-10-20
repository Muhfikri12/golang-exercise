package main

import "fmt"

/*

Latihan 1: Buatlah struct Car dengan field Brand dan Year, dan buat interface Drive() yang mengembalikan string, implementasikan pada struct Car.

Latihan 2: Buatlah struct Animal dengan field Species dan Sound. Implementasikan interface MakeSound() yang mengembalikan suara hewan tersebut.

Latihan 3: Buat struct Rectangle dengan field Length dan Width. Buat interface Area() yang menghitung dan mengembalikan luas persegi panjang.

Latihan 4: Buatlah interface Vehicle dengan metode Speed() yang mengembalikan kecepatan maksimum kendaraan. Implementasikan pada struct Bike dan Car dengan kecepatan yang berbeda.

Latihan 5: Buat struct Book dengan field Title dan Author. Buat interface GetDetails() yang mengembalikan detail buku tersebut. Implementasikan interface tersebut pada struct Book.

*/

// Jawaban No 1

type Car struct {
	Brand string
	Year int
}

type Driver interface {
	Drive() string
}

func (c Car) Drive() string {
	return fmt.Sprintf("%s is Brand that has been publish at %d", c.Brand, c.Year)
}

// Jawaban No 2

type Animal struct {
	Species string
	Sound string
}

type SoundMaker interface {
	MakeSound() string
}

func (a Animal) MakeSound() string {
	return fmt.Sprintf("the sound of %s is %s", a.Species, a.Sound)
}

// Jawaban No 3

type Rectangle struct {
	Length int
	Width int
}

type CountArea interface {
	Area() int
}

func (r Rectangle) Area() int {
	
	return r.Length * r.Width
}

// Jawaban NO 4

type Vehicle interface {
	Speed() int
}

type Bike struct {
	Brand string
	MaxSpeed int
}

type Boot struct {
	Brand string
	MaxSpeed int
}

func (b Bike) Speed() int {
	return b.MaxSpeed
}

func (bo Boot) Speed() int {
	return bo.MaxSpeed
}

// Jawaban No 5

type Book struct {
	Title string
	Author string
}

type Detail interface {
	getDetails() string
}

func (b Book) getDetails() string {
	return fmt.Sprintf("the %s is written by %s", b.Title, b.Author)
}