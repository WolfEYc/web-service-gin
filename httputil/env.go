package httputil

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const SECURE = "SECURE"
const MONGODB_URI = "MONGODB_URI"
const ENV = ".env"
const INSECURE_SCEME = "http"
const SECURE_SCHEME = "https"
const JWT_SECRET = "JWT_PRIVATE_KEY"

func GetSchemes() []string {
	if IsSecure {
		return []string{SECURE_SCHEME}
	} else {
		return []string{INSECURE_SCEME}
	}
}

func GetIsSecure() bool {
	_ = godotenv.Load(ENV)
	securityStr, hasSecureStr := os.LookupEnv(SECURE)

	return !hasSecureStr || securityStr != "false"
}

var IsSecure bool = GetIsSecure()

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
