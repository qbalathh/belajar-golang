package main

import (
	"fmt"
)

type Address struct {
	city, province, country string
}

// pointer pada method
type Man struct {
	name string
}

func (man *Man) maried() {
	man.name = "Mr. " + man.name
}

func main() {
	// pass by value
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1     // membuat salinan dari address1 by reference (pointer)
	address2.city = "Jakarta" // mengubah city pada address2

	addressA := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	addressB := &addressA                                 // membuat salinan dari addressA by reference (pointer)
	*addressB = Address{"Yogyakarta", "DIY", "Indonesia"} // membuat salinan dari addressA by value

	alamat := new(Address)

	fmt.Println("Address 1:", address1) // data tetap
	fmt.Println("Address 2:", address2)
	fmt.Println("Address A:", addressA)
	fmt.Println("Address B:", addressB)
	fmt.Println("Address Alamat:", alamat) // data kosong karena belum diinisialisasi

	eko := Man{"eko"}
	eko.maried()          // memanggil method mar
	fmt.Println(eko)      // Eko menjadi Mr. Eko
	fmt.Println(eko.name) // menampilkan nama Eko yang sudah berubah
}
