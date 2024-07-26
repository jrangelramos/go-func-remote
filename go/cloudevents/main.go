package main

import (
	"fmt"
	"os"

	ce "knative.dev/func-go/cloudevents"
)

func main() {
	if err := ce.Start(ce.DefaultHandler{Handler: Handle}); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
