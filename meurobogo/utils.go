package main
import (
	"fmt"
)
func logDebug(msg string) {
	if modoDebug {
		fmt.Println("[DEBUG]", msg)
	}
}