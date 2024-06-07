package main

import (
	"encoding/json"
	"fmt"
	"github.com/willpw/GoDmsoftWillpwr"
	"os"
)

func main() {

	DllJson := make(map[string]string)
	_ = json.Unmarshal(dmsoft.DllJson, &DllJson)
	DllJson["AiEnableFindPicWindow"] = "AiEnableFindPicWindow"
	DllJson["AiFindPic"] = "AiFindPic"
	DllJson["AiFindPicEx"] = "AiFindPicEx"
	DllJson["AiFindPicMem"] = "AiFindPicMem"
	DllJson["AiFindPicMemEx"] = "AiFindPicMemEx"

	Jsonfile, _ := json.MarshalIndent(DllJson, "", "\t")

	_ = os.WriteFile("Ddll.json", Jsonfile, 0644)

	fmt.Println(DllJson["Willpwr"])
}
