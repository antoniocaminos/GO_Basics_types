package main

import "fmt"

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
//Structs
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type Car struct {
	NumberOfTyres int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// basics types (numeros strings y booleanos)
/* var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUint uint
var myFloat float32
var myFloat64 float64 */

//	aggregate type ( arrays y structs)

//	reference type ( pointers, slices, maps, functions, channels)

//	interface type

func main() {
	/* myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Tony"
	log.Println(myString)
	myString = "jhon"

	var myBool = true
	myBool = false
	log.Println(myBool) */
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//ARRAY
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	/* var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println(myStrings) */

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//Structs
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	myCar := Car{
		NumberOfTyres: 4,
		Luxury:        false,
		Make:          "Suran",
		Model:         "A90",
		Year:          2010,
	}
	fmt.Printf("mi auto es un %s %s %d", myCar.Make, myCar.Model, myCar.Year)
}
