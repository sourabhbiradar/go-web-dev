package main

import(
	"fmt"
	"encoding/json"
	"os"
)

type Person struct{
	Name string  `json:"name,omitempty"` // Name --> name in json data
	Age int 	 `json:"age"`
	Gender string `json:"-"` // JSON ignores this
	Address []string `json:"address"`
}

func main(){
	// marshalling & unmarshalling
	fmt.Println("\tmarshalling & unmarshalling")

	// marshalling 
	p:=Person{
		Name:"Abc",
		Age:18,
		Address:[]string{"flat 201","Bluru"},
	}

	 b,err:=json.Marshal(p)
	if err!=nil{
		fmt.Println("Error Marshalling JSON",err)
		panic(err)
	}
	fmt.Println("JSON data :",string(b))

	// unmarshalling
	var um Person

	err=json.Unmarshal(b,&um)  // ([]byte,where to store data)
	if err!=nil{
		panic("Error while unmarshalling")
	}
	fmt.Println("unmarshalled :",um)
	fmt.Println("Name :",um.Name)

	// encoding
	fmt.Println("encoding")

	f,err:=os.Create("file.json")  // created file.json
	if err!=nil{
		fmt.Println("Failed to create file")
		panic(err)
	}
	defer f.Close()  // recommanded to close file

	err = json.NewEncoder(f).Encode(p)  // README.md
	if err!=nil{
		fmt.Println("Error while encoding",err)
	}

	// decoding
	fmt.Println("decoding")

    f1,err:= os.Open("file.json")  // open file.json to read from it
    if err!=nil{
		fmt.Println("Failed to open file")
		panic(err)
	}
	defer f1.Close()

	var newP Person

	err = json.NewDecoder(f1).Decode(&newP)
	if err!=nil{
		fmt.Println("Error while decoding")
		panic(err)
	}

	fmt.Println("newD :",newP)
	
}