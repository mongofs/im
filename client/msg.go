package client

import "encoding/json"

type Basic struct {
	// Sid is the unique identifier of the message
	Sid int64
	// MSG is the main body of information transmission
	Msg interface{}
}



func( b*Basic)Marshal()([]byte,error){
	return json.Marshal(b)
}
