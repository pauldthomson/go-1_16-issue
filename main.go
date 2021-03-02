package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	l := prometheus.Labels{
		"version": "something",
	}

	fmt.Printf("Just making it feel used: %v", l)
}
