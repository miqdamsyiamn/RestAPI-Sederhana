package main

import "fmt"

// Struct User
type User struct {
 ID   int
 Name string
 Age  int
}

// Fungsi untuk menampilkan data user
func printUser(user User) {
 fmt.Printf("ID: %d\n", user.ID)
 fmt.Printf("Name: %s\n", user.Name)
 fmt.Printf("Age: %d\n", user.Age)
}

func main() {
 // Inisialisasi struct User dengan data
 user := User{ID: 1, Name: "Miqdam", Age: 21}

 // Memanggil fungsi untuk mencetak data user
 printUser(user)
}