package constant

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ErrBodyCannotBeEmpty = errors.New("body cannot be empty")
var PORT = "8080"

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	
	PORT = os.Getenv("PORT")
}