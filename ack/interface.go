// ACK is a measure for the long connection program to maintain the message accuracy.
// When you don't want your message to be lost, it is very necessary to maintain an
// ACK plug-in in the server. The ACK principle is to maintain a storage of SID unique
// message flag kV type in the plug-in


// WarnCall ： The warning callback is due to the memory usage inside the ACK. A threshold
// must be set. This threshold can be set during instantiation. If it exceeds this
// threshold, the ACK needs to notify the program caller

// Handle : As ack is a basic plug-in, it should not execute the internal logic of
// other programs. When the contents of the internally maintained ack queue are taken
// out by the timer, a handle method should be passed into the ACK internal call as a parameter

package ack




type Acker interface {

	//Put user messages into ack queue
	AddMessage(token string,sid int64,content []byte)error

	// Delete ack queue data through sid
	DelMessage(sid int64)

	//Set Handle func
	Handle(func(token string, content []byte, sid int64)error)
}
