package main

import "fmt"

//Bills is the structs that contains all neceassary variables to calculate the water bill price per person
type Bills struct {
	FixedCosts    float32
	VariableCosts float32
	NumDays       float32
	NumPeople     float32
	NumRobbieDays float32
}

func main() {

	b := Bills{
		FixedCosts:    0.0,
		VariableCosts: 0.0,
		NumDays:       0.0,
		NumRobbieDays: 0.0,
	}

	///Map of each person on the water bill
	guys := map[string]float32{"Robbie": 0.0, "Jackson": 0.0, "Chris": 0.0, "Daniel": 0.0}

	fmt.Println("Enter the total fixed cost: ")
	// Takes the fixed costs from user input
	fmt.Scanln(&b.FixedCosts)

	fmt.Println("Enter the total variable cost: ")
	// Takes the total variable cost from the users input
	fmt.Scanln(&b.VariableCosts)

	fmt.Println("Enter the total number of billing days: ")
	// Takes the total number of billing period days from the user
	fmt.Scanln(&b.NumDays)

	fmt.Println("Enter the number of days Robbie was there: ")
	// Takes the number of days Robbie was at the house from the user
	fmt.Scanln(&b.NumRobbieDays)

	//Calculates the fixed cost per person and returns each persons value.
	FixedCostPerPerson(b.FixedCosts, 4, guys)
	VariableCostPerPerson(b.VariableCosts, b.NumDays, b.NumRobbieDays, guys)
	fmt.Println(guys)

}

//FixedCostPerPerson prints the total fixed cost
func FixedCostPerPerson(fixed float32, NumPeople float32, guys map[string]float32) map[string]float32 {
	fmt.Println("Total Fixed cost is: $", fixed)
	guys["Robbie"] = fixed / NumPeople
	guys["Jackson"] = fixed / NumPeople
	guys["Chris"] = fixed / NumPeople
	guys["Daniel"] = fixed / NumPeople
	return guys
}

//VariableCostPerPerson calculates the amount each person pays of the variable water costs based on the amount of time spent at the house.
func VariableCostPerPerson(VariableCost float32, NumDays float32, NumRobbieDays float32, guys map[string]float32) map[string]float32 {
	ChargePerDay := VariableCost / NumDays
	PricePerDayRobbieThere := (ChargePerDay * NumRobbieDays) / 3
	PricePerDayRobbieNotThere := (ChargePerDay * (NumDays - NumRobbieDays)) / 2

	guys["Robbie"] = guys["Robbie"] + PricePerDayRobbieThere
	guys["Jackson"] = guys["Jackson"] + PricePerDayRobbieThere + PricePerDayRobbieNotThere
	guys["Chris"] = guys["Chris"] + PricePerDayRobbieThere + PricePerDayRobbieNotThere

	return guys
}
