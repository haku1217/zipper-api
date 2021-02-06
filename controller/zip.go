package controller

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/haku1217/zipper/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type ZipController struct {
	repo model.ZipRepository
}

func newZipController(db *sqlx.DB) *ZipController {
	repo, _ := model.NewZipRepository(db)
	return &ZipController{repo}
}

func (c *ZipController) upload(w http.ResponseWriter, r *http.Request) {
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
	}
	if r.Method == "GET" {
		token := ""
		t := template.Must(template.ParseFiles("view/upload.html"))
		err := t.Execute(w, token)
		if err != nil {
			fmt.Println(err)
		}
	}
}
