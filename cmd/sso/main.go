package sso

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
