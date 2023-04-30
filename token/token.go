package token

import "time"

type Maker interface {
	CreateToken(UserID int64, username string, role string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
