// // There is only one main function in the application which is the starting point your app.

// // whatever you importing from the package, the first letter should be capital.

// // run command : go run main.go
package main

import (
	"fmt"
	// "math"
	// "strings"
	// "sort"
	// "strings"
)

//// Initial Function for data types and String Formatting

// func main() {
// 	//Strings
// 	var nameOne string = "Piyush" // Way-1
// 	var nameTwo = "Piyush1" //way-2

// 	fmt.Println(nameOne)
// 	fmt.Println(nameTwo)

// 	nameThree := "Joshi" // THis can be used in inside the function only. (: is basically replacing the var keyword) also it is used at the initialization

// 	fmt.Println(nameThree)

// 	// For Integers its the same thing

// 	//nits and memory

// 	var num1 int8 = 35
// 	var num2 uint = 5  // Negative values are not allowed

// 	fmt.Println(num1)
// 	fmt.Println(num2)

// 	//Formatted Strings
// 	//Format Specifier => %v
// 	// %v -> Any variable value will be printed
// 	// %q -> Quote -> only string values will be printed in quotes
// 	// %T -> It will print the Type of the variable.
// 	// %f -> It will print the flat value (0.1f,0.2f to limit the decimal points)

// 	fmt.Printf("My name us %v and my age is %v \n",nameOne,num1)

// 	//Sprintf() -> Saves the formatted strind into the variable

// 	var str = fmt.Sprintf("My name us %v and my age is %v",nameOne,num1)

// 	fmt.Println("Saved String",str)

// }

//// Function For Slices and Arrays

// func main(){
// 	// Arrays

// 	var ages [3]int = [3]int{20,25,30}

// 	names := [4]string{"Piyush","Joshi","Hello","Legend"}

// 	fmt.Println(ages, len(ages))
// 	fmt.Println(names)

// 	//Slices

// 	var scores =  []int{100,50,60}

// 	scores = append(scores, 85) // => It will not update the scores slice instead it will create the new slice to update the scores we use the given method

// 	fmt.Println(scores)

// }

//// Function to see Standard Libraries

// func main(){
// 	greeting :="This is my new Function"

// 	//Import packages here it is a strings package
// 	//Contains is the method in strings package. which checks that in string given string is there or not? The statement returns boolean.
// 	//https://www.codecademy.com/resources/docs/go/strings/count => All the methods for strings

// 	// fmt.Println(strings.Count("this is new world is it?","i"))

// 	// fmt.Println((strings.Index("Hii How are you?","Ho")))

// 	var arr = strings.Split(greeting," ")

// 	fmt.Println(strings.Split(arr[0],"i"))

// 	//sort package

// 	ages := []int{19,23,34,45,35,64,73,83}

// 	names:= []string{"My","Iss", "Is","Piyush"}
// 	sort.Ints(ages)
// 	sort.Strings(names)
// 	index := sort.SearchInts(ages,11)

// 	fmt.Println(index)
// 	fmt.Println(names)
// 	fmt.Println(ages)

// }

//// Loops

// func main(){
// 	// x := 0

// 	// for x < 5{
// 	// 	fmt.Println("Value of x is:", x)
// 	// 	x++
// 	// }

// 	// for i := 0; i < 5; i++{
// 	// 	fmt.Println("Value of i is:" ,i)
// 	// }

// 	//Loop through Slice

// 	names:= []string{"My","Iss", "Is","Piyush"}

// 	// for i :=0; i < len(names); i++{
// 	// 	fmt.Println(names[i])
// 	// }

// 	for index, value := range names{
// 		fmt.Printf("position  at index %v is %v:",index,value )
// 	}
// }

//// Conditional and Booleans

// func main(){
// 	age:= 45
// 	names := []int{12,243,5453,654,235,653}

// 	fmt.Println(age <= 50);

// if age <= 40 {
// 	for i :=0; i < len(names); i++{
// 	fmt.Println(names[i])
// }
// }else if age >=30{
// 	for i :=0; i < len(names)-1; i++{
// 		fmt.Println(names[i+1])
// 	}
// }

// 	for index, value := range names{
// 		if index == 1{
// 			fmt.Printf("position  at index %v is %v:",index,value )
// 			continue
// 		}

// 		fmt.Printf("position  at index %v is %v:",index,value )

// }

// }

//// Functions in Go

// func sayGreeting(n string){
// 	fmt.Printf("Good Morning %v \n",n)
// }

// func sayBye(n string){
// 	fmt.Printf("Good Bye %v \n",n)
// }

// func countSheep(x int, y string){
// 	fmt.Printf("There are %v %v in the jungle", x,y)
// }

// //Passing a function as a argument

// func countNames(n []string, f func(string)){
// 	for _,v := range n{
// 		f(v)
// 	}
// }
// // We have to define a return type after the parent thesis
// func circleArea(r float64) float64{
// 	return math.Pi *r * r
// }

// func main(){
// 	// names:= []string{"Piyush","Joshi", "",""}
// 	// countNames(names, sayBye)

// 	circleArea(11.11)

// 	fmt.Print(circleArea(11.11))
// }

////Multiple return

// func getInitials(n string ) (string,string){
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s," ")

// 	var initials []string

// 	for _,v := range names{

// 		initials = append(initials, v[:1])

// 	}

// 	if len(initials) >1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

// func main() {
// 	fni,sni := getInitials("piyush joshi")

// 	fmt.Println(fni)
// 	fmt.Println(sni)
// }

////Maps

// func main(){
// 	menu := map[string]float64{
// 		"soup":4.99,
// 		"pie":7.99,
// 		"salad": 6.99,
// 	}

// 	fmt.Println(menu)
// 	fmt.Println(menu["pie"])

// 	//Looping Maps

// 	for k, v := range menu{
// 		fmt.Println(k,"-",v)
// 	}
// }

////Pointers
// func updateName(x *string) {
// 	*x = "Joshi"
// }

// func main(){
// 	name := "Piyush"

// 	// updateName(name)

// 	//&name will point out the memory location of name variable
// 	// fmt.Println("Memory address of name is:",&name)

// 	m := &name

// 	//* ahead of a pointer shows the value which it is pointing to

// 	// fmt.Println("Value Of a Pointer is:", *m)

// 	// The important learning is that you cannot update the original value by using any function as GO creates a copy of variable but you can change the variable from a memory location using pointers. E.g

// 	fmt.Println(name) // before Update
// 	updateName(m)
// 	fmt.Println(name) // After Changing the value from memory location

// }

//// Structs

// Blueprint which descibes the type of data (It is kind of interface in typescript)

// Remember to run the both files

func main(){
	mybill := newBill("Piyush's Bill")

	mybill.updateTip(10)

	mybill.addItem("Banana",5.44)

	fmt.Println(mybill.format())


}


//// Receiver Function


