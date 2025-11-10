package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/tejtex/profetch/internal/ascii"
	"github.com/tejtex/profetch/internal/display"
	"github.com/tejtex/profetch/internal/info"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed with current timestamp

	logo, ok := ascii.FetchLogo(".");
	info, err := info.FetchInfo(".", rand.Intn(37 - 31) + 31);
	if err != nil {
		fmt.Fprintf(os.Stderr, "profetch: %v\n", err);
		os.Exit(1);
	}
	display.Render(info, logo, ok);

}