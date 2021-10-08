package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}
//store fn will save a file with product detail in it
func (prod *Product) store() {
	//Create a file with 'id' name
	file,_ := os.Create(prod.id+".txt")
	//Sprintf will return String, not log to cmd
	content := fmt.Sprintf("ID:%v\nTitle:%v\nDesc:%v\nPrice:%v\n",prod.id,prod.title,prod.description,prod.price)
	//Write content with string
	file.WriteString(content)
	//MUST close file after writing
	file.Close()
	
}

//a fn that tied to Product struct; print all product's data to StdIn
func (prod *Product) printData() {
	fmt.Printf("ID:%v\n",prod.id)
	fmt.Printf("Title:%v\n",prod.title)
	fmt.Printf("Desc:%v\n",prod.description)
	fmt.Printf("Price:%v\n",prod.price)
}

func main() {
	newProduct:=getProductInput()
	newProduct.printData()
	newProduct.store()
}

func getProductInput() *Product {
	var reader = bufio.NewReader(os.Stdin)
	id:=readUserInput(reader,"id:")
	title:=readUserInput(reader,"title:")
	description:=readUserInput(reader,"description:")
	price,_ := strconv.ParseFloat(readUserInput(reader,"price:"),64) 

	return &Product{id,title,description,price}
}

func readUserInput(reader *bufio.Reader ,promtText string) string {
	fmt.Print(promtText)
	userInput,_ := reader.ReadString('\n')
	if(runtime.GOOS == "windows"){
		userInput = strings.Replace(userInput,"\r\n","",-1)
		}else{
		userInput = strings.Replace(userInput,"\n","",-1)	
	}
	return userInput
}

func NewProduct(id string, title string, description string, price float64) *Product {
	return &Product{id: id, title: title, description: description, price: price}
}


// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.