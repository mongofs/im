package im

import (
	"github.com/zhenjl/cityhash"
)

func  Index(token string,size uint32) uint32 {
	return  cityhash.CityHash32([]byte(token), uint32(len(token))) % size
}