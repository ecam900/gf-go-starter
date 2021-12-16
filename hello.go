package hello

import (
	"fmt"
	"net/http"
	"os"
)

func init() {

	if os.Getenv("ENVIRONMENT") == "dev" {
		fmt.Println("Running in dev mode.")

	} else {
		fmt.Println("Running in production mode.")
	}

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello World 👋")
}
