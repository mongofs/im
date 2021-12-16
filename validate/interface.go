package validate

import "github.com/mongofs/im/client"

type Validater interface {
	Validate(token string)error
	ValidateFailed(err error,cli client.Clienter)
}
