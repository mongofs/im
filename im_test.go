package im

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNew(t *testing.T) {
	bs := New(DefaultOption())
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
}

func ExampleNew() {
	bs := New(DefaultOption())
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
}