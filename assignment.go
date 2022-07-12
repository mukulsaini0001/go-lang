package main

import (
	"fmt"
)

type Address1 struct {
	name        string
	age_in_year float64
	origin      string
	head_office string
}
type Address2 struct {
	streetNumber string
	landmark     string
	city         string
	pincode      int
	state        string
}

func main() {
	//address 1 struct
	a1 := Address1{
		name:        "Tolexo Online Pvt.Ltd",
		age_in_year: 8.5,
		origin:      "Noida",
		head_office: "Noida,Uttar Pradesh",
	}
	//address 2 struct
	a2 := Address2{
		streetNumber: "91 Springboard",
		landmark:     "Axis Bank",
		city:         "Noida",
		pincode:      201301,
		state:        "Uttar Pradesh",
	}
	fmt.Println(a1)
	fmt.Println(a2)

}
