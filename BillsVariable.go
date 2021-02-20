package main

import "fmt"

//Bills is the struct that contains the general variable in order to calculate bill price per person
type Bills struct {
	// Amount of Fixed costs for billing
	FixedCosts    float32
	// Amount of Variable costs for billing
	VariableCosts float32
	// Total number of days in the billing period
	TotalNumDays  int8
	// Number of people splitting the bill
	NumPerson     int
	// List of People spliting the bill
	People        []Person
	// Person struct that holds all the info of each person splitting the bill
	Person        Person
}

//Person is a struct that contains the required variables for each person to calculate the water bill
type Person struct {
	// Name of person
	Name        string
	// Number of days the person was at the property 
	NumDays     int8
	// Amount the person has to pay for the billing period
	AmountToPay float32
	// number of days the person was at the property longer then the lowest person
	DaysPastMinPerson int8
}

func main() {
	b := Bills{
		FixedCosts:    0.0,
		VariableCosts: 0.0,
		TotalNumDays:  0,
		NumPerson:     1,
		People:        []Person{},
		Person: Person{
			Name:        "",
			NumDays:     0,
			AmountToPay: 0.0,
			DaysPastMinPerson: 0,
		},
	}

	fmt.Println("Enter the total number of people: ")
	// Takes the total number of billing period days from the user
	fmt.Scanln(&b.NumPerson)

	b = GetData(b)
	
	for i := 0; i < b.NumPerson; i++ {
		b = GetPersonData(b)
	}

	fmt.Println("Total Fixed cost is: $", b.FixedCosts)
	for i := 0; i < b.NumPerson; i++ {
		b = FixedCostPerPerson(b.FixedCosts, b.NumPerson, b, i)
	}

	fmt.Println("Total Variable Cost is: $", b.VariableCosts)
	for i := 0; i < b.NumPerson; i++ {
		b = VariableBaseCostPerPerson(b.VariableCosts, b.TotalNumDays, b, i)
	}

	DaysSpentEqual := CheckDaysSpentEqual(b)

	if DaysSpentEqual == false{
		var lowest = FindLowest(b)
		lowest.DaysPastMinPerson = 0
		b = GetDaysPastLowest(b, lowest)
	}

	fmt.Println(b.People)

}

//FixedCostPerPerson calculates the amount each person pays of the fixed cost
func FixedCostPerPerson(fixed float32, NumPerson int, bills Bills, i int) Bills{
	pricePerPerson := fixed / float32(bills.NumPerson)

	bills.People[i].AmountToPay = pricePerPerson
	return bills
}

// CheckDaysSpentEqual checks if everyone was at the house for the full billing period.
func CheckDaysSpentEqual(bills Bills) bool{
	check := true
	for i := 1; i < bills.NumPerson; i++ {
		if bills.People[i].NumDays != bills.TotalNumDays{
			check = false
		}
	}
	return check
}

// FindLowest finds the person who spent the lowest amount of time at the house.
func FindLowest(bills Bills) Person{
	lowest := bills.People[0]
	for i := 1; i < bills.NumPerson; i++{
		if bills.People[i].NumDays < lowest.NumDays{
			lowest = bills.People[i]
		}
	}
	return lowest
}

// GetDaysPastLowest gets the amount of days each person spent
// at the house longer then the lowest person
func GetDaysPastLowest(bills Bills, lowestPerson Person) Bills{
	var lowest int8 = bills.TotalNumDays - lowestPerson.NumDays
	for i := 0; i < bills.NumPerson; i++{
		bills.People[i].DaysPastMinPerson = bills.People[i].NumDays - lowest
	}
	return bills
}

//VariableBaseCostPerPerson calculates the amount each person pays of the variable costs if time spent at the house is 
// all the same if not it requires the helper func VariableCostDiffDays()
func VariableBaseCostPerPerson(VariableCost float32, TotalNumDays int8, bills Bills, i int) Bills {
	ChargePerDay := VariableCost / float32(TotalNumDays)
	ChargePerPerson := ChargePerDay / float32(bills.NumPerson)

	bills.People[i].AmountToPay += ChargePerPerson * float32(bills.People[i].NumDays)

	return bills
}

// VariableCostDiffDays is a helper func to calculate the amount each person needs to pay if they dont all spend the same amount of time at the house
// func VariableCostDiffDays(VariableCost float32, TotalNumDays int8, bills Bills, i int) Bills{
	
//}


//GetData gets the data from user input and assigns to the fields of Bills struct
func GetData(bills Bills) Bills {
	fmt.Println("Enter the total fixed cost: ")
	// Takes the fixed costs from user input
	fmt.Scanln(&bills.FixedCosts)

	fmt.Println("Enter the total variable cost: ")
	// Takes the total variable cost from the users input
	fmt.Scanln(&bills.VariableCosts)

	fmt.Println("Enter the total number of billing days: ")
	// Takes the total number of billing period days from the user
	fmt.Scanln(&bills.TotalNumDays)

	return bills
}

// GetPersonData gets the data of each person
func GetPersonData(bills Bills) Bills {
	
	j := Person {
		Name:        "",
		NumDays:     0,
		AmountToPay: 0.0,
		DaysPastMinPerson: 0,
	}
	
	fmt.Println("Enter name: ")
	// Takes the user input for persons name
	fmt.Scanln(&j.Name)

	fmt.Println("Enter total number of days person was at house during pay period: ")
	// Takes number of days person was at the house during the pay period
	fmt.Scanln(&j.NumDays)

	if (j.NumDays > bills.TotalNumDays) {
		fmt.Printf("Please re-enter number of days less than or equal to %d: \n", bills.TotalNumDays)
		// Takes number of days person was at the house during the pay period
		fmt.Scanln(&j.NumDays)
	}

	bills.People = append(bills.People, j)

	return bills
}