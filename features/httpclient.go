package features

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchAndParse() {

	urlString := "https://jsonplaceholder.typicode.com/users/1"
	resp, err := http.Get(urlString)
	if err != nil {
		panic("request faild !!!")
	}

	respBytes, _ := io.ReadAll(resp.Body)
	parsedResp := make(map[string]any)

	err = json.Unmarshal(respBytes, &parsedResp)
	if err != nil {
		panic("parsing faild !!!")
	}
	fmt.Println(string(respBytes))
	fmt.Printf("%v", parsedResp)
}
