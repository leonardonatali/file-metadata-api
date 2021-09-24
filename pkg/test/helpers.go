package test

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

const filename = "./file.txt"

func FormData(values map[string]io.Reader) (content *bytes.Buffer, contentType string, err error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// File
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, "", err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, "", err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, "", err
		}

	}

	content = &b
	contentType = w.FormDataContentType()

	w.Close()

	return
}

func CreateFile() {
	ioutil.WriteFile(filename, []byte("this is a file!\n"), 0777)
}

func MustOpen() *os.File {
	f, _ := os.Open(filename)
	return f
}
