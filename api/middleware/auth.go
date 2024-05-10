package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var SECRET = []byte("supersecurepassword")

func GenerateJWT(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the Authorization header value
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			// If Authorization header is missing, return a 400 Bad Request error
			return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Authorization header is missing"})
		}

		// Check if the Authorization header has the Bearer scheme
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			// If the authorization format is invalid, return a 400 Bad Request error
			return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid authorization format"})
		}

		// Extract the token string by removing the Bearer prefix
		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		if tokenString == "" {
			// If the token string is empty, return a 401 Unauthorized error
			return c.Status(fiber.StatusUnauthorized).SendString("You're Unauthorized!")
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// If the signing method is unexpected, return a 401 Unauthorized error
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key for token verification
			return []byte(SECRET), nil
		})
		if err != nil {
			// If there's an error parsing the token, return a 401 Unauthorized error
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		if !token.Valid {
			// If the token is not valid (e.g., expired or tampered with), return a 401 Unauthorized error
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
		}

		// Store the claims in the context for use in subsequent handlers
		c.Locals("claims", token.Claims.(jwt.MapClaims))

		// Continue to the next handler
		return c.Next()
	}
}

// func ExtractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
// 	if request.Header["Token"] != nil {
// 		tokenString := request.Header["Token"][0]
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
// 				return nil, fmt.Errorf("there's an error with the signing method")
// 			}
// 			return rsaPrivateKey, nil
// 		})

// 		if err != nil {
// 			return "Error Parsing Token: ", err
// 		}
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if ok && token.Valid {
// 			email := claims["email"].(string)
// 			return email, nil
// 		}

// 	}
// 	return "unable to extract claims", nil
// }
