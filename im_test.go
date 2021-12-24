package im

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
	_ "net/http/pprof"
)

func TestNew(t *testing.T) {
	bs := New(DefaultOption())

	go func() {
		http.ListenAndServe(":6060",nil)
	}()
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