package configs

//todo
import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("loading env from .env success")
	}
	// fmt.Println()
}
