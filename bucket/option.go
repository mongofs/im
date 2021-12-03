package bucket

import (
	"context"
)

type OptionFunc func(h *hash)

func WithContext (ctx context.Context) OptionFunc {
	return func(h *hash) {
		h.ctx = ctx
	}
}


func WithCallBack (callback func ()) OptionFunc {
	return func(h *hash) {
		h.offline =callback
	}
}



func WithSize (size int32) OptionFunc {
	return func(h *hash) {
		h.size =size
	}
}



