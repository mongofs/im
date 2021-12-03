package im

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	bs := New()
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
}

func ExampleNew() {
	bs := New()
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
}