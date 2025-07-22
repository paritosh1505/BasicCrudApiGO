package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Userid string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	dataToSend := Data{"100", "This is hello", "Go outside and take sun bath"}
	urlName := "https://jsonplaceholder.typicode.com"

	clientReq := &http.Client{}
	endPointPost := "posts"
	var reqType int

	for {
		fmt.Println("Enter option 1.GET \n 2.Update \n 3.Delete\n 4.Add \n 5.Exit")
		fmt.Scanln(&reqType)
		switch reqType {
		case 1:
			makeGetCall("GET", urlName, endPointPost, clientReq)
		case 2:
			makeUpdateCall("PUT", urlName, endPointPost, dataToSend, clientReq)
		case 3:
			//makeDeleteCall("Delete", urlName, endPointPost)
		case 4:
			//makeAddCall("PUT", urlName.endPointPost)
		case 5:
			fmt.Println("Exiting the scenario")
			break
		default:
			fmt.Println("Enter valid scenario")
		}
	}

}

func makeGetCall(method string, urlName string, endPoint string, clientReq *http.Client) {

	newurl := urlName + "/" + endPoint
	fmt.Println(newurl)
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		fmt.Println("Get request error")
		return
	}
	resp, err := clientReq.Do(req)
	if err != nil {
		fmt.Println("resposne errorr")
		return
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in printing the body")
	}
	fmt.Println(string(content))

}
func makeUpdateCall(method string, url string, endPoint string, datasend Data, clientReq *http.Client) error {
	//newurl := url + "/" + endPoint
	newurl := "https://jsonplaceholder.typicode.com/posts/1"
	marshallData, _ := json.Marshal(datasend)

	req, err := http.NewRequest(method, newurl, bytes.NewBuffer(marshallData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error while updating the entry")
		return nil
	}
	resp, _ := clientReq.Do(req)

	contentCheck, _ := io.ReadAll(resp.Body)
	fmt.Println(string(contentCheck))
	return nil
}
