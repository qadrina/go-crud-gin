package initializers

// capital letters of a function means that it is a global function (can be used within other files)

import (
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
