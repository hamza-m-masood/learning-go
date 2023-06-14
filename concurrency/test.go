package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdOutBuff bytes.Buffer
	fmt.Fprintln(&stdOutBuff, "Producer Done")
	stdOutBuff.WriteTo(os.Stdout)
}
