package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("got here", r.Body)

		r.ParseMultipartForm(32 << 20) // limit your max input length!
		var buf bytes.Buffer
		// in your case file would be fileupload
		file, _, err := r.FormFile("asset")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(&buf, file)
		contents := buf.String()

		f, err := os.Create("./sample.pdf")
		if err != nil {
			fmt.Println(err)
			return
		}

		l, err := f.WriteString(contents)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		fmt.Println(l, "bytes written successfully")

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		//return

		//err := r.ParseMultipartForm(0)
		//if err != nil {
		//	fmt.Println("-------> ", err)
		//}

		//for k, v := range r.Form {
		//	fmt.Println(k, "<====>", v)
		//}

		//fmt.Println("====> ", r.Form)
	})

	log.Fatal(http.ListenAndServe(":8181", router))
}
