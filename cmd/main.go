package main

import (
	"fmt"

	"github.com/v1bh475u/LibMan_MVC/pkg/api"
)

func main() {
	fmt.Println("Server running on port 8080...")
	api.StartApi()
}
