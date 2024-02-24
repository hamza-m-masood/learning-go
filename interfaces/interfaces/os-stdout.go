package interfaces

import "os"

func writeData(w *os.File, data []byte) {

	w.Write(data)

}
func RunOs() {
	writeData(os.Stdout, []byte("hello hello"))
}
