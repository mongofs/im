package im

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)


var r = rand.New(rand.NewSource(time.Now().Unix()))

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetSliceOfStrings(len int)[]string{
	strs := make([]string,len)
	for i:=0;i<len ; i++{
		strs[i]= RandString(20)
	}
	return strs
}

func TestIndex(t *testing.T) {
	tests := []struct{
		token string
		size uint32
		want uint32
	}{
		{
			token: "abcdefg",
			size:  10,
			want:  Index("abcdefg",10),
		},

		{
			token: "123432dadsffadsacsf",
			size:  10,
			want:  Index("123432dadsffadsacsf",10),
		},
		{
			token: "QgegWDADSSDASD",
			size:  10,
			want:  Index("QgegWDADSSDASD",10),
		},
		{
			token: "DFSFDSFS12445",
			size:  10,
			want:  Index("DFSFDSFS12445",10),
		},
	}

	for _,v:= range tests {
		assert.Equal(t,v.want, Index(v.token,v.size))
	}
}


func TestIndex_1000distribute(t *testing.T){
	var size = 10
	mp := make(map[uint32]int,size)
	strs := GetSliceOfStrings(1000)
	for i:= 0;i<len(strs);i++{
		idx := Index(strs[i],uint32(size))
		mp[idx]++
	}
	fmt.Println(mp)
	// output : Hash function only needs the relative balance of values
}


func TestIndex_10000distribute(t *testing.T){
	var size = 10
	mp := make(map[uint32]int,size)
	strs := GetSliceOfStrings(10000)
	fmt.Println(len(strs))
	for i:= 0;i<len(strs);i++{
		idx := Index(strs[i],uint32(size))
		mp[idx]++
	}
	fmt.Println(mp)
	// output : Hash function only needs the relative balance of values
}



func TestIndex_100000distribute(t *testing.T){
	var size = 10
	mp := make(map[uint32]int,size)
	strs := GetSliceOfStrings(100000)
	fmt.Println(len(strs))
	for i:= 0;i<len(strs);i++{
		idx := Index(strs[i],uint32(size))
		mp[idx]++
	}
	fmt.Println(mp)
	// output : Hash function only needs the relative balance of values
}


func Test_main(t *testing.T){
	fmt.Println(2 << 10)
}
