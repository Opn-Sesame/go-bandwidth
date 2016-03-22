package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"

func main6() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))

	fmt.Println("Errors")
	errors, _ := api.GetErrors()
	for _, err := range errors {
		fmt.Printf("%s\t%s\t%s\n", err.Code, err.Category, err.Time)
	}
}
