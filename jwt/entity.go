package jwt

import (
	"github.com/ArpinFidel/go-utils/rest/restid"
	"github.com/dgrijalva/jwt-go"
)

// Claims claims
type Claims struct {
	jwt.StandardClaims
	MemberID       restid.ID `json:"id"`
	RoleID         restid.ID `json:"role_id"`
	MemberUsername string    `json:"username"`
	Type           string    `json:"type"`
}
