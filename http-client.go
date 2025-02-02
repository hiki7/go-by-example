package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//test
	//sending http request
	resp, err := http.Get("https://gobyexample.com")
	//if there is an error then the program will be terminated
	if err != nil {
		panic(err)
	}
	//defer works like finally in other languages, so some code will always be executed
	//we need to close the body of the response in order to prevent the memory leakage
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body) //we are using bufio.NewScanner in order to read the response line by line
	for i := 0; scanner.Scan() && i < 5; i++ { //for example we are reading only first 5 lines
		fmt.Println(scanner.Text())
	}

	//generally, why do wee need to use bufio.Scanner()
	//because, it is efficient for reading large responses beacause it reads line by line
	//instead of loading everything into memory

	//if error occurs while scanning, it will panic and terminate the program
	if err := scanner.Err(); err != nil { //Scanner.Err() detects unexpected errors
		//for example errors might happen with network failure while reading the response
		panic(err)
	}
}