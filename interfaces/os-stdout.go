package main

import "os"

func writeData(w *os.File, data []byte) {

	w.Write(data)

}
func main() {
	writeData(os.Stdout, []byte("hello hello"))
}
