package main

import (
	"fmt"
)

func main() {

	a, b, c, d, e := 20, 20, 20, 20, 20
	a ^= 2
	b &= 2
	c |= 2
	d <<= 2
	e >>= 2
	fmt.Println(a, b, c, d, e)

	array := [...]int{1, 2, 3, 4, 5}
	array[2] = 10
	slice := array[2:4]
	fmt.Println(slice)
	fmt.Println((len(array)), array[2], array[3])
	sliceBaru := make([]string, 5, 10)
	sliceBaru[0], sliceBaru[1], sliceBaru[2] = "satu", "dua", "tiga"
	fmt.Println(sliceBaru)
	data_orang := map[string]string{
		"nama":      "Budi",
		"alamat":    "Jakarta",
		"pekerjaan": "Programmer",
		"status":    "Aktif",
	}
	fmt.Println(data_orang["nama"])
	delete(data_orang, "status")
	fmt.Println(data_orang)
	fmt.Println(len(data_orang))

	nilai32 := int32(10000)
	nilai64 := int64(nilai32)
	nilai8 := int8(nilai32)
	fmt.Println(nilai32, nilai64, nilai8)

	fmt.Println("Hello, World!")
	fmt.Println((1 + 2) * 3)

	if bilangan := 13; bilangan%2 == 0 {
		fmt.Println("a adalah bilangan genap")
	} else if bilangan%3 == 0 {
		fmt.Println("a adalah kelipatan tiga")
	} else {
		fmt.Println("a adalah bilangan ganjil")
	}

	switch bilangan := 12; {
	case bilangan%3 == 0:
		fmt.Println("a adalah kelipatan tiga")
	case bilangan%2 == 0:
		fmt.Println("a adalah bilangan genap")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}
} // tipe data number ada 2 yaitu integer dan floating point
