package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"

func main7() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))

	fmt.Println("Messages")
	messages, _ := api.GetMessages()
	for _, m := range messages {
		fmt.Printf("==========\n%s\t%s -> %s\n%s\n==========\n\n", m.Time, m.From, m.To, m.Text)
	}
}
