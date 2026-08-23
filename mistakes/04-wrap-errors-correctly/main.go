package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	err := sql.ErrNoRows

	// No
	wrapped := fmt.Errorf("getUser 32: %v", err)
	fmt.Println(errors.Is(wrapped, sql.ErrNoRows))

	// Yes
	wrapped = fmt.Errorf("getUser 32: %w", err)
	fmt.Println(errors.Is(wrapped, sql.ErrNoRows))
}
