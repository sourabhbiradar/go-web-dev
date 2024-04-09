package main

import (
	"fmt"
	"os"
	"io"
	//"io/ioutil"
)

func main(){
	// io.Reader & io.Writer
	fmt.Println("\tio.Reader & io.Writer")

    // io.Reader
	of,err:=os.Open("README.md")
	if err!=nil{
		fmt.Println("cud not open file :",err)
	}
	defer of.Close()

	buffer:=make([]byte,1024)  // where read data will be stored

	for {
		n, err := of.Read(buffer) // read from 'of' & store in 'buffer'
		if err == io.EOF {
			fmt.Println("Reached end of file")
			break
		} else if err != nil {
			fmt.Println("Error reading from file:", err)
			break
		}
		// fmt.Println(string(buffer[:n]))
		fmt.Printf("number of letters read :%d\n",n)
	}

	// ioutil.ReadAll
	// utilR,err:=ioutil.ReadAll(of)
	// if err!=nil{
	// 	fmt.Println("Error reading from file:", err)
	// }
	// fmt.Println(string(utilR))

	// io.Writer
	fmt.Println("io.Writer")

	data:=[]byte("Write this data to write.txt file.")

	wf,err:=os.Create("write.txt")
	if err!=nil{
		panic(err)
	}
	defer wf.Close()

	n, err := wf.Write(data) // read from 'of' & store in 'buffer'
		if err!=nil{
			fmt.Println("Error writting to write.txt")
		}

	fmt.Printf("number of letters written :%d\n",n)

}
