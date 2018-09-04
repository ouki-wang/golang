package main
/*
import (
    "io"
    "net/http"
    "os"
)

func main() {
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
    http.ListenAndServe(":5000", nil)
}
*/
 
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var stDir string

func getCurDir() string {
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func FileServer(w http.ResponseWriter, r *http.Request) {

	fmt.Println("stDir=", stDir)
	fmt.Println("r.URL.Path=",r.URL.Path)
	if strings.HasPrefix(r.URL.Path, "/") {
		file := stDir + r.URL.Path
		fmt.Println(file)
		f, err := os.Open(file)
		defer f.Close()
 
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintln(w, "File not exist")
			return
		}
		http.ServeFile(w, r, file)
		return
	} else {
		fmt.Fprintln(w, "Hello world")
	}
}
func main() {
	stDir = getCurDir()
	http.HandleFunc("/", FileServer)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
