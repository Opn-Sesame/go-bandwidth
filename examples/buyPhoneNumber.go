package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"
import "strings"

func main2() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))
	numbersResult, _ := api.GetAvailableNumbers(bandwidth.AvailableNumberTypeLocal, &bandwidth.GetAvailableNumberQuery{
		City: "Cary",
		State: "NC",
		Quantity: 3})
	l := len(numbersResult)
	numbers := make([]string, l)
	for i := 0;  i < l;  i++{
		numbers[i] = numbersResult[i].Number
	}
	fmt.Printf("Found numbers: %s\n", strings.Join(numbers, ", "))
	api.CreatePhoneNumber(&bandwidth.CreatePhoneNumberData{Number: numbers[0]})
	fmt.Printf("Number %s is yours now\n", numbers[0])
}
