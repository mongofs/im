package bucket

import (
	"context"
)

type Option interface {
	Apply (h *hash)
}


type optionFunc func(h *hash)


func (o optionFunc)Apply(h *hash){
	o(h)
}


func WithContext (ctx context.Context) Option {
	return optionFunc( func(h *hash)  {
		h.ctx =ctx
	})
}


func WithCallBack (callback func ()) Option {
	return optionFunc(func(h *hash) {
		h.offline =callback
	})
}

func WithSize (size int8) Option {
	return optionFunc(func(h *hash) {
		h.size =size
	})
}



