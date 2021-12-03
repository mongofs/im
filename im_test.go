package im

import (
	"log"
)

func ExampleNew() {
	bs := New()
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
}