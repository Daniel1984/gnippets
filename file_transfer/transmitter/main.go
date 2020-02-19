package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("./file.pdf")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("asset", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(part, file); err != nil {
		log.Fatal(err)
	}

	nameSegments := map[string]string{
		"hola": "HOLA",
		"name": "DAN",
	}

	for key, val := range nameSegments {
		if err := writer.WriteField(key, val); err != nil {
			log.Fatal(err)
		}
	}

	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8181", body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		log.Fatal(resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n### INFO: response from StorIQ:%s\n", respBody)
	//body := &bytes.Buffer{}
	//writer := multipart.NewWriter(body)
	//part, err := writer.CreateFormFile("uploadfile", filepath.Base(file.Name()))
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//if _, err := io.Copy(part, file); err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//if err := writer.WriteField("HEY", "HO"); err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//if err := writer.Close(); err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//req, err := http.NewRequest("POST", "http://localhost:8181", body)
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//fmt.Println(writer.FormDataContentType())
	//req.Header.Set("Content-Type", writer.FormDataContentType())

	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//defer resp.Body.Close()

	//if resp.StatusCode >= http.StatusBadRequest {
	//	fmt.Println("ERROR: ", resp.StatusCode)
	//	return
	//}

	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("ERROR: ", err)
	//	return
	//}

	//fmt.Printf("================================> %s\n", respBody)

	//err := postFile("./file.pdf", "http://localhost:8181")
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
