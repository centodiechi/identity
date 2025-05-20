package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func formatToken(tokenString string) string {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return tokenString
	}

	headerHash := sha256.Sum256([]byte(parts[0]))
	payloadHash := sha256.Sum256([]byte(parts[1]))
	signatureHash := sha256.Sum256([]byte(parts[2]))

	header := base64.RawURLEncoding.EncodeToString(headerHash[:])[:8]
	payload := base64.RawURLEncoding.EncodeToString(payloadHash[:])[:8]
	signature := base64.RawURLEncoding.EncodeToString(signatureHash[:])[:8]

	return fmt.Sprintf("%s-%s-%s", header, payload, signature)
}

func GetTokenPair(userId string) (string, string, error) {
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return
}

func GenerateAccessToken(userId string) (string, error) {
	return "", nil
}

func GenerateRefreshToken(userId string) (string, error) {
}
