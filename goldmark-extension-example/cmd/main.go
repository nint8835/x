package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"

	goldmarkextensionexample "github.com/nint8835/x/goldmark-extension-example"
)

var testingInput = `# Example

The following should be evaluated: ` + "`math:1+1`"

func main() {
	md := goldmark.New(goldmark.WithExtensions(goldmarkextensionexample.New()))

	var buf bytes.Buffer
	if err := md.Convert([]byte(testingInput), &buf); err != nil {
		fmt.Printf("Failed to convert markdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Rendered HTML:\n\n%s", buf.String())
}
