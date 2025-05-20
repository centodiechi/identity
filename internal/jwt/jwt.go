package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	signKey              = "dojf!2kfodlmxcWWER@#@!@#"
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 30 * 24 * time.Hour
)

type TokenClaims struct {
	UserID string `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func formatToken(tokenString string) string {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return tokenString
	}

	return fmt.Sprintf("%s-%s-%s", parts[0], parts[1], parts[2])
}

func GetTokenPair(userId, role string) (string, string, error) {
	accessToken, err := GenerateAccessToken(userId, role)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := GenerateRefreshToken(userId, role)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func GenerateAccessToken(userId string, role string) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "identity-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return formatToken(tokenString), nil
}

func GenerateRefreshToken(userId string, role string) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "identity-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return formatToken(tokenString), nil
}
