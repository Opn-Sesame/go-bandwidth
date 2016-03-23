package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"

func main1() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))
	id, _ := api.CreateApplication(&bandwidth.ApplicationData{Name: "Golang demo app"})
	fmt.Printf("App ID is %s\n", id)
	api.DeleteApplication(id)
}
