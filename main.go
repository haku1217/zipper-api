package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func top(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("view/top.html"))
	str := "Sample Message"
	if err := t.ExecuteTemplate(w, "top.html", str); err != nil {
		log.Fatal(err)
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			fmt.Println(err)
		}
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)

		reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
		count := 0
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			if count == 0 {
				fmt.Println(line)
			}
			count++
			output := strings.Join(line[:], ",")
			fmt.Fprintf(w, output)
		}
		s := strconv.Itoa(count)
		fmt.Fprintf(w, s)
		// f, err := os.Create("./test/" + handler.Filename)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// defer f.Close()
		// o, err := io.Copy(f, file)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(o, "**o")
	}
	if r.Method == "GET" {
		// crutime := time.Now().Unix()
		// h := md5.New()
		// io.WriteString(h, strconv.FormatInt(crutime, 10))
		// token := fmt.Sprintf("%x", h.Sum(nil))
		// fmt.Println(token, "**token")
		token := ""
		t := template.Must(template.ParseFiles("view/upload.html"))
		t.Execute(w, token)
	}
}

func logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T", h)
		fmt.Println("hoge")
		h(w, r)
	})
}

func main() {
	fmt.Println("Print")
	http.HandleFunc("/", logger(top))
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
