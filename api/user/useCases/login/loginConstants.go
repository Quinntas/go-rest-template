package login

import "time"

const (
	TOKEN_EXPIRATION_TIME = time.Second * 3600
	TOKEN_REDIS_KEY       = "login:token:"
)
