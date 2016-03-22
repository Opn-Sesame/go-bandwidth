package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"

//run: sendSMS <from-number> <to-number> <text>
func main5() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))
	fromNumber := os.Args[1]
	toNumber := os.Args[2]
	text := os.Args[3]
	id, err := api.SendMessageTo(fromNumber, toNumber, text)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Message ID is %s", id)
}
