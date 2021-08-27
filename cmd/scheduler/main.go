package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/humorliang/scheduler-framework/pkg/register"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	cmd := register.Register()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
