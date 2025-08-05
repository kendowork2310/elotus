package jwt

import (
	"elotus/cmd/common/errs"
	"elotus/pkg/cfg"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	SecretKey        = "JWT_SECRET_KEY"
	RefreshSecretKey = "JWT_REFRESH_SECRET_KEY"
)

type Claims struct {
	Username  string `json:"username"`
	TokenType string `json:"token_type"` // "access" || "refresh" token
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` //  expiration time in seconds
}

// GenerateTokenPair creates both access and refresh tokens
func GenerateTokenPair(username string) (*TokenPair, error) {
	// Access token (15 minutes)
	accessClaims := Claims{
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.Reader().MustGetString(SecretKey)))
	if err != nil {
		return nil, errs.NewCustomError(errs.ErrInternalServer)
	}

	// Refresh token (7 days)
	refreshClaims := Claims{
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Reader().MustGetString(RefreshSecretKey)))
	if err != nil {
		return nil, errs.NewCustomError(errs.ErrInternalServer)
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    15 * 60, // 15 minutes
	}, nil
}

// ValidateAccessToken validates the access token and returns the claims
func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewCustomError(errs.ErrInvalidToken)
		}
		return []byte(cfg.Reader().MustGetString(SecretKey)), nil
	})

	if err != nil {
		return nil, errs.NewCustomError(errs.ErrInvalidToken)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.TokenType != "access" {
			return nil, errs.NewCustomError(errs.ErrInvalidToken)
		}
		return claims, nil
	}

	return nil, errs.NewCustomError(errs.ErrInvalidToken)
}

// ValidateRefreshToken validates the refresh token and returns the claims
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewCustomError(errs.ErrInvalidToken)
		}
		return []byte(cfg.Reader().MustGetString(RefreshSecretKey)), nil
	})

	if err != nil {
		return nil, errs.NewCustomError(errs.ErrInvalidToken)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.TokenType != "refresh" {
			return nil, errs.NewCustomError(errs.ErrInvalidToken)
		}
		return claims, nil
	}

	return nil, errs.NewCustomError(errs.ErrInvalidToken)
}

// RefreshAccessToken generates a new access token using a valid refresh token
func RefreshAccessToken(refreshTokenString string) (string, error) {
	claims, err := ValidateRefreshToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	// Generate new access token
	accessClaims := Claims{
		Username:  claims.Username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	tokenString, err := token.SignedString([]byte(cfg.Reader().MustGetString(SecretKey)))
	if err != nil {
		return "", errs.NewCustomError(errs.ErrInternalServer)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	return ValidateAccessToken(tokenString)
}
