package client

import "fmt"

type Example struct {}


func ( e *Example) Handle (cli Clienter,data []byte){
	fmt.Println(string(data))
}
