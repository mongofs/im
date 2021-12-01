package main


import (
	"log"
	"websocket/internal/basic"
)

func main (){


	bs := basic.New()

	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()

}