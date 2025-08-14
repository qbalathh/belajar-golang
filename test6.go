package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	NamaSaya := "Muhammad Iqbal Athhar"

	fmt.Println(
		strings.Contains("joko widodo", "o"),
		strings.Split(NamaSaya, " "),
		strings.ToLower(NamaSaya),
		strings.ToUpper(NamaSaya),
		strings.Trim("Msaya suka makanM", "M"),
	)

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}
	fmt.Println(strconv.ParseInt("205003", 10, 64))
	fmt.Println(
		strconv.FormatInt(2025, 10),
		strconv.FormatInt(2025, 2),
		strconv.FormatInt(2025, 8),
		strconv.FormatInt(2025, 16),
		strconv.Itoa(1000),
	)

	fmt.Println(
		math.Max(10, 20),
		math.Min(10, 20),
		math.Abs(-10.5),
		math.Pow(2, 10),
		math.Sqrt(100),
		math.Round(10.5),
		math.Floor(10.5),
		math.Ceil(10.5),
		math.Trunc(10.5),
		math.Sin(math.Pi/2),
		math.Cos(math.Pi/2),
		math.Tan(math.Pi/4),
		math.Log(100),
		math.Log10(100),
	)

	dataNama := list.New()
	dataNama.PushBack("Muhammad") // menambahkan data dari paling ujung belakang (tail)
	dataNama.PushBack("Iqbal")
	dataNama.PushBack("Athhar")
	dataNama.PushFront("St.") // menambahkan data dari paling ujung depan (head)

	for e := dataNama.Front(); e != nil; e = e.Next() { // .Front() untuk mendapatkan elemen pertama dan .Next() untuk mendapatkan elemen berikutnya
		fmt.Println(e.Value)
		fmt.Println(dataNama.Front().Prev()) // .Prev() untuk mendapatkan elemen sebelumnya
		fmt.Println(dataNama.Back().Next())  // .Back() untuk mendapatkan elemen terakhir
	}

	dataRing := ring.New(5) // membuat ring dengan kapasitas 5

	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = "Data ke-" + strconv.Itoa(i)
		dataRing = dataRing.Next() // berpindah ke elemen berikutnya
	}
	dataRing.Do(func(value interface{}) {
		fmt.Println(value) // mencetak nilai dari setiap elemen dalam ring
	})
}
