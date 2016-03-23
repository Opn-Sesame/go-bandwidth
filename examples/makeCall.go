package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"

//run: makeCall <from-number> <to-number>
func main3() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))
	fromNumber := os.Args[1]
	toNumber := os.Args[2]
	id, err := api.CreateCall(&bandwidth.CreateCallData{From: fromNumber, To: toNumber})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Call ID is %s", id)
}
