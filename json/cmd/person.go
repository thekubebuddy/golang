package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	FirstName    string `json:"firstName"`
	LastName     string `jsoson:"lastName"`
	Age          int    `json:"age,omitempty"` // Omit if its empty or not initialized
	CreditCard   string `json:"-"`
	MobileNumber int    `json:"mobileNumber"`
}

func (p Person) ShowPersonInfo() {
	fmt.Printf("%+v\n", p) // Show value and type of person struct
}

func (p Person) MarshalPerson() string {
	pBytes, _ := json.Marshal(p)
	fmt.Printf("[######## Marshaling Person info into Bytes ########]\n")
	err := ioutil.WriteFile("personInfo.json", pBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}
	return string(pBytes)
}

func UnmarshalPersonFromFile(fileName string) Person {
	fmt.Printf("[######## Unmarshing Person from provide file: %s ########]\n", fileName)
	pBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var p Person
	json.Unmarshal(pBytes, &p)
	return p
}
