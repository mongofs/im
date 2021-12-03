package validate

type Validater interface {
	Validate(token string)error
}
