package authcontroller

import (
  "errors"
  "time"

  jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
  "user1": "password1",
  "user2": "password2",
}

// Claims ...
type Claims struct {
  Username string `json:"username"`
  jwt.StandardClaims
}

// Login ...
func Login(username string, password string) (string, error) {
  expectedPassword, ok := users[username]
  if !ok || expectedPassword != password {
    return "", errors.New("Wrong credentials")
  }
  expirationTime := time.Now().Add(1 * time.Minute)
  // Create claims with username and expiration
  claims := &Claims{
    Username: username,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(),
    },
  }
  // Declare the token with the algorithm used for signing, and the claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  // Create the JWT string
  return token.SignedString(jwtKey)
}

// Authenticate ...
func Authenticate(jwtTokenString string) (*jwt.Token, error) {
  claims := &Claims{}
  // Parse the JWT string and store the result in claims
  token, err := jwt.ParseWithClaims(jwtTokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return jwtKey, nil
  })
  return token, err
}

// Register ...
func Register(username string, password string) error {
  // Dummy sign up
  users[username] = password
  return nil
}
