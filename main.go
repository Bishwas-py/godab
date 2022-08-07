package main

import (
	"encoding/json"
	"fmt"
)

const Version = "0.0.1"

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Address Address
}
type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Pincode string `json:"pincode,omitempty"`
}

func main() {
	dir := "./"
	db, err = New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
	employees := []User{
		{"John", "30", "Orolando, FL", Address{"123 Main St", "Orolando", "FL", "32122"}},
		{"Jane", "25", ", FL", Address{"456 Main St", "New York", "Florida", "21809"}},
		{"Jack", "35", "Orolando, FL", Address{"789 Main St", "Kolkata", "West Bengal", "70009"}},
		{"Sameer", "35", "Orolando, FL", Address{"789 Main St", "Kathmandu", "Nepal", "70809"}},
		//	fill the rest of the array with dummy data and address
		{"Bishwas", "22", "", Address{"2B Wasington City", "Wasington City", "DC", "20001"}},
		{"Pritam", "22", "", Address{"2B Wasington City", "Wasington City", "DC", "20001"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(records)

	allUsers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal(f, &employeeFound); err != nil {
			fmt.Println("Error:", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	//if err := db.Delete("users", "John"); err != nil {
	//	fmt.Println("Error:", err)
	//}
	//
	//if err := db.Delete("users", ""); err != nil {
	//	fmt.Println("Error:", err)
	//}
}
