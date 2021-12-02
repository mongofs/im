package cache

import (
	"github.com/garyburd/redigo/redis"
	"websocket/internal/basic/validate"
)

type cache struct {
	redis *redis.Pool
}

const (
	TOKEN_APP_PREFIX ="app:token:user_id:"
)


func NewCache (redis  *redis.Pool)validate.Validater{
	return &cache{redis: nil}
}



func (c *cache)Validate(token string)error{
	conn := c.redis.Get()
	defer conn.Close()
	_ ,err:=redis.String(conn.Do("get", TOKEN_APP_PREFIX+token))
	if err!=nil {
		return err
	}
	return nil
}


