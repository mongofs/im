package client


type RoundOptions struct {
	ReadBufSize  int
	WriteBufSize int
	MessageBuffer 	uint
	Protocol 	 uint8  // json /proto
	MessageType  uint8
}