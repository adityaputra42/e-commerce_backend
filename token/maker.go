package token

import "time"

type Maker interface {
	CreateToken(username, uid, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
