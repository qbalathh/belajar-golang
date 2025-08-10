package main

import (
	"errors"
	"fmt"
)

// membuat struct sama seperti membuat class pada bahasa pemrograman lain
type Person struct { //membuat blueprnit
	Name, Address string
	Age           int
}

func (orang Person) sayHi(name string) {
	fmt.Println("hi", name, "my name is", orang.Name)
}
func sayHello(orang Person, name string) {
	fmt.Println("hello", name, "my name is", orang.Name)
}

// error interface
func pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b tidak boleh nol")
	}
	return a / b, nil
}

// interface
type Human interface {
	sayHuuu(name string)
}

// interface kosong
func ups() interface{} {
	return "Ups, something went wrong!"
}

type People struct {
	Name string
}

func (p People) sayHuuu(name string) {
	fmt.Println("Huuu", name, "from", p.Name)
}

// Nil adalah sebuah tipe data yang merepresentasikan ketiadaan nilai atau objek.
// Nil dapat digunakan untuk merepresentasikan variabel yang belum diinisialisasi atau objek yang
func newMap(name string) map[string]string {
	if name == "" {
		fmt.Println("data kosong") // mengembalikan nil jika nama kosong
		return nil
	} else {
		return map[string]string{"name": name} // mengembalikan map dengan nama
	}
}

// type assertions
func random() interface{} {
	return "ini adalah string"
}

func main() {
	hello := "hello world"
	fmt.Println(hello)

	person := Person{ //membuat objek dari struct Person
		Name:    "John Doe",
		Address: "123 Main St",
		Age:     30,
	}
	People := People{"Alice"}
	person2 := Person{"Jane Doe", "456 Elm St", 25} //posisi parameter harus sesuai dengan blueprint
	fmt.Println(person2)
	fmt.Print(person)
	sayHello(person, "Alice")
	person.sayHi("Bob")       //memanggil method dari struct Person
	People.sayHuuu("Charlie") //memanggil method dari People struct
	fmt.Println(ups())        //memanggil fungsi ups yang mengembalikan interface kosong

	fmt.Println(newMap(""))
	fmt.Println(newMap("GoLang")) //memanggil fungsi newMap dengan nama

	result, err := pembagian(10, 2) //memabuat 2 variable karena return valuenya 2
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}

	//resultString := random().(string) // melakukan type assertion untuk mengubah interface kosong menjadi string
	//resultInt := random().(int)       // akan panic jika random() tidak mengembalikan int
	// harus dipastikan type assertions harus sesuai dengan tipe data return dari fungsi random()
	// jika tidak sesuai, akan terjadi panic
	// untuk menghindari panic, bisa menggunakan _, ok := random().(string) untuk
	// melakukan pengecekan apakah type assertion berhasil atau tidak
	//fmt.Println(resultString)         // akan mencetak "ini adalah string"
	switch nilai := random().(type) { // melakukan type switch untuk menentukan tipe data dari hasil
	case string: // menggunakan switch expresssion agar tidak terjadi panic bila tipe data tidak sesuai
		fmt.Println("Hasil adalah string:", nilai)
	case int:
		fmt.Println("Hasil adalah int:", nilai)
	default:
		fmt.Println("Hasil adalah tipe yang tidak dikenal")
	}
}
