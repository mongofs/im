package bucket

/*

var r = rand.New(rand.NewSource(time.Now().Unix()))

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}



type cli struct {
	token string
}

func newCil()*cli {
	return &cli{token: RandString(20)}
}

func (c *cli) Send(bytes []byte, i ...int64) {

}

func (c *cli) Offline() {
	panic("implement me")
}



func getUserMap(nums int)(Bucketer){
	bucket := New(DefaultOption())
	for i:=0 ;i< nums;i++ {
		user := newCil()
	}
	return bucket
}

func getUsersWithRand(nums int)(Bucketer, string){
	bucket := New()
	token := ""
	for i:=0 ;i< nums;i++ {
		user := newCil()
		bucket.Register(user,user.token)
		if i ==100{
			token =user.token
		}
	}
	return bucket,token
}

//BenchmarkHash_1000Send-6     	   62710	     18341 ns/op
func BenchmarkHash_1000SendAll(b *testing.B) {
	bu:= getUserMap(1000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bu.SendAll([]byte("content"),false)
	}
}

//BenchmarkHash_10000Send-6    	    6572	    178934 ns/op
func BenchmarkHash_10000SendAll(b *testing.B) {
	bu:= getUserMap(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bu.SendAll([]byte("content"),false)
	}
}

//BenchmarkHash_1000000SendAll-6   	      63	  19182915 ns/op
// 19ms下推到1000000 goroutine,每秒预计下推能力：1000000*500 =5.2千万消息下推
func BenchmarkHash_1000000SendAll(b *testing.B) {
	bu:= getUserMap(1000000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bu.SendAll([]byte("content"),false)
	}
}



// BenchmarkHash_1000Send-6   	30086384	        38.1 ns/op
func BenchmarkHash_1000Send(b *testing.B) {
	bu,token:= getUsersWithRand(1000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bu.Send([]byte("content"),token,false)
	}
}


//BenchmarkHash_100000Send-6   	30029653	        37.6 ns/op
func BenchmarkHash_100000Send(b *testing.B) {
	bu,token:= getUsersWithRand(100000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		bu.Send([]byte("content"),token,false)
	}
}


*/