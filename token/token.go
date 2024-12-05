package token

import "fmt"

var tokenMaker Maker

func InitTokenMaker(SecretKey string) error {
	var err error

	tokenMaker, err = NewJWTMaker(SecretKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker %w", err)
	}
	return nil
}

func GetTokenMaker() Maker {
	return tokenMaker
}
