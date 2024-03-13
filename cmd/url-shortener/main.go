package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog -> log/slog package

	// TODO: init storage: sqlite

	// TODO: init router: chi -> net/http "chi render"

	// TODO: tun server

}
