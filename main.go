package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name     string
	Location string
	Gender   string
}

func main() {
	// Menerima input jumlah data dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan jumlah data yang ingin dimasukkan: ")
	var numData int
	fmt.Scanln(&numData)

	var people []Person

	// Mengisi data sendiri
	for i := 1; i <= numData; i++ {
		var name, location, gender string

		fmt.Printf("Masukkan nama data ke-%d: ", i)
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Printf("Masukkan lokasi data ke-%d: ", i)
		location, _ = reader.ReadString('\n')
		location = strings.TrimSpace(location)

		fmt.Printf("Masukkan gender data ke-%d (M/F): ", i)
		gender, _ = reader.ReadString('\n')
		gender = strings.TrimSpace(gender)

		people = append(people, Person{Name: name, Location: location, Gender: gender})
	}

	// Menampilkan hasil data yang telah dimasukkan
	fmt.Println("\nData yang telah dimasukkan:")
	for _, person := range people {
		fmt.Printf("Name: %s, Location: %s, Gender: %s\n", person.Name, person.Location, person.Gender)
	}
}
