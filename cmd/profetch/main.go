package main

import (
	"fmt"
	"os"

	"github.com/tejtex/profetch/internal/ascii"
	"github.com/tejtex/profetch/internal/display"
	"github.com/tejtex/profetch/internal/info"
)

func main() {
	logo, ok := ascii.FetchLogo();
	info, err := info.FetchInfo(".");
	if err != nil {
		fmt.Fprintf(os.Stderr, "profetch: %v\n", err);
		os.Exit(1);
	}
	display.Render(info, logo, ok);

}