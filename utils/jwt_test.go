//path jwt_test.go

package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Test GenerateJWT
	tokenString, err := GenerateJWT("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Test ValidateToken
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	assert.NoError(t, err)
	claims, ok := token.Claims.(*JWTClaim)
	assert.True(t, ok)
	assert.Equal(t, "test", claims.Id)
	assert.True(t, claims.ExpiresAt > time.Now().Local().Unix())
}

func TestValidateToken(t *testing.T) {
	tokenString, err := GenerateJWT("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	err = ValidateToken(tokenString)
	assert.NoError(t, err)

	// Test expired token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	assert.NoError(t, err)
	claims, ok := token.Claims.(*JWTClaim)
	assert.True(t, ok)
	claims.ExpiresAt = time.Now().Local().Unix() - 1
	tokenString, err = token.SignedString(jwtKey)
	assert.NoError(t, err)
	err = ValidateToken(tokenString)
	assert.Error(t, err)

}
