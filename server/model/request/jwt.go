package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID     uint
	StuNum string `json:"stuNum"`
	jwt.StandardClaims
}
