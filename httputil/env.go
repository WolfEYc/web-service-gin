package httputil

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const SCHEME = "SCHEME"
const MONGODB_URI = "MONGODB_URI"
const ENV = ".env"
const DEFAULT_SCHEME = "https"
const JWT_SECRET = "JWT_PRIVATE_KEY"

func GetURLProtocol() []string {
	scheme, hasScheme := os.LookupEnv(SCHEME)

	if hasScheme {
		return []string{scheme}
	} else {
		return []string{DEFAULT_SCHEME}
	}
}

var IsSecure bool = GetURLProtocol()[0] == DEFAULT_SCHEME

func GetMongoURI() string {
	uri, exists := os.LookupEnv(MONGODB_URI)
	if !exists {
		err := godotenv.Load(ENV)

		if err != nil {
			log.Fatal(err)
		}

		return os.Getenv(MONGODB_URI)
	} else {
		return uri
	}
}

func GetJWTPrivateKey() []byte {
	return []byte(os.Getenv(JWT_SECRET))
}

var JWTPrivateKey []byte = GetJWTPrivateKey()
