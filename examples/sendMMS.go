package examples

import "github.com/bandwidthcom/go-bandwidth"
import "os"
import "fmt"
import "math/rand"


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}


//run: sendMMS <from-number> <to-number> <text> <path-to-jpg-file-to-attach>
func main() {
	api, _ := bandwidth.New(os.Getenv("CATAPULT_USER_ID"), os.Getenv("CATAPULT_API_TOKEN"), os.Getenv("CATAPULT_API_SECRET"))
	fromNumber := os.Args[1]
	toNumber := os.Args[2]
	text := os.Args[3]
	jpgPath := os.Args[4]
	mediaName := randomString(64)
	err := api.UploadMediaFile(mediaName, jpgPath, "image/jpeg")
	if err != nil {
		fmt.Printf("Error on uploading file: %s", err.Error())
		return
	}
	id, err := api.SendMessageTo(fromNumber, toNumber, text,
		map[string]interface{}{"media": []string{
			fmt.Sprintf("https://api.catapult.inetwork.com/v1/users/%s/media/%s", os.Getenv("CATAPULT_USER_ID"), mediaName)}})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Message ID is %s", id)
}
