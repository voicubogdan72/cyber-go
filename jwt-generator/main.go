package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my_secret_key")

// Generează token
func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Verifică token
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verifică algoritmul
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodă de semnare invalidă: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
}

func main() {
	username := "bogdan"
	tokenString, err := generateJWT(username)
	if err != nil {
		fmt.Println("❌ Eroare la generarea tokenului:", err)
		return
	}

	fmt.Println("✅ Token generat:", tokenString)

	// Verifică tokenul
	validatedToken, err := validateJWT(tokenString)
	if err != nil || !validatedToken.Valid {
		fmt.Println("❌ Token invalid:", err)
		return
	}

	// Extrage claims
	if claims, ok := validatedToken.Claims.(jwt.MapClaims); ok {
		fmt.Println("🔐 Token valid.")
		fmt.Println("👤 Username:", claims["username"])
		fmt.Println("⏳ Expiră la:", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("❌ Nu s-au putut citi claims.")
	}
}
