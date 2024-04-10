package auth_type

import "github.com/dgrijalva/jwt-go/v4"

type AuthTokenClaims struct {
	TokenUUID  string   `json:"tid"`        // 토큰 UUID
	UserID     int      `json:"userId"`     // 유저 ID
	Name       *string  `json:"name"`       // 유저 이름
	Email      string   `json:"mail"`       // 유저 메일
	Permission []string `json:"permission"` // 유저 역할
	jwt.ClaimStrings
}
