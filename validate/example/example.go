package example

import "errors"

type  DefaultValidate struct {

}


func (d *DefaultValidate) Validate(token string)error{

	if token == "" {return errors.New("token is not good ")}
	return nil
}
