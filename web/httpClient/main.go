package main

import(
	"net/http"
	"fmt"
	"encoding/json"
	"time"
)

// {
// 	"activity": "Start a band",
// 	"type": "music",
// 	"participants": 4,
// 	"price": 0.3,
// 	"link": "",
// 	"key": "5675880",
// 	"accessibility": 0.8
// }

type boringApi struct{
	Activity string `json:"activity"`
	Type string `json:"type"`
	Participants int `json:"participants"`
	Price float32 `json:"price"`
	Link string `json:"link"`
	Key string `json:"key"`
	Accessibility float32 `json:"accessibility"`
}

var url = "https://www.boredapi.com/api/activity/"

func main(){
	// HTTP Client
	fmt.Println("\tHTTP Client")

	client:=&http.Client{  // creating new http client
		Timeout:10*time.Second,
	}

	req,err:=http.NewRequest("GET",url,nil) // creating request
	if err!=nil{
		fmt.Println("Error creating request :",err)
		return
	}

	res,err:=client.Do(req) // sending request via client
	if err!=nil{
		fmt.Println("Error sending request ;",err)
		return
	}
	defer res.Body.Close()

	var resBody boringApi 

	if err:=json.NewDecoder(res.Body).Decode(&resBody);err!=nil{ // reading response body
		fmt.Println("Error while reading :",err)
		return
	}

	fmt.Printf("%+v",resBody)
}