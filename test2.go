package main

import "fmt"

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
}
