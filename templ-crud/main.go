package main

import (
	"context"
	"log/slog"
	"os"
)

var dbURI = "file:data.dbmode=rwc"
var addr = "localhost:8080"

func main() {
	log := slog.Default()
	ctx := context.Background()
	if err := run(ctx, log); err != nil {
		log.Error("Failed to run server", slog.Any("error", err))
		os.Exit(1)
	}
}
