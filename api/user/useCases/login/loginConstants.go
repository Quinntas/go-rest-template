package login

import "time"

const (
	TokenExpirationTime = time.Second * 3600
	TokenRedisKey       = "login:token:"
)
