package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address1 struct {
	name        string
	age_in_year int
	origin      string
	head_office string
}
type Address2 struct {
	streetNumber int
	landmark     string
	city         string
	pincode      int
	state        string
}

func main() {
	a1 := Address1{
		name:        "Tolexo Online Pvt.Ltd",
		age_in_year: 8,
		origin:      "Noida",
		head_office: "Noida,Uttar Pradesh",
	}
	j, err := json.Marshal(a1)
	if err != nil {
		log.Fatalf("Error occured during marshaling.Error:%s,%s,%s", err.Error())
	}
	fmt.Println("address1 JSON: %s,%s,%s\n", string(j))
	a2 := Address2{
		streetNumber: "91 Springboard",
		landmark:     "Axis Bank",
		city:         "Noida",
		pincode:      201301,
		state:        "Uttar Pradesh",
	}
	if err != nil {
		log.Fatalf("Error occured during marshaling.Error:%s,%s,%s,%s", err.Error())
	}
	fmt.Println("address2 JSON: %s,%s,%s,%s\n", string(j))
}
