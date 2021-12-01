package basic

import "testing"


//BenchmarkIndex-6   	86346265	        14.1 ns/op
func BenchmarkIndex(b *testing.B) {
	for i := 0;i<b.N;i ++{
		Index("sadfadsf",20)
	}
}