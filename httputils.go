package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

//MultiPartFile -- file
type MultiPartFile struct {
	FileName string
	FileData []byte
}

//MultiPartForm -- form
type MultiPartForm struct {
	FormName  string
	FormValue string
}

//MultiPart --
type MultiPart struct {
	Files []MultiPartFile
	Forms []MultiPartForm
}

//GetMultiParts --
func GetMultiParts(r *http.Request) (MultiPart, error) {
	mp := MultiPart{}

	mr, err := r.MultipartReader()
	if err != nil {
		return mp, err
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			return mp, nil
		}

		if err != nil {
			return mp, err
		}

		formName := p.FormName()
		fileName := p.FileName()
		if formName != "" && fileName == "" {
			formValue, _ := ioutil.ReadAll(p)
			mp.Forms = append(mp.Forms, MultiPartForm{FormName: formName, FormValue: string(formValue)})
		}
		if fileName != "" {
			fileData, _ := ioutil.ReadAll(p)
			mp.Files = append(mp.Files, MultiPartFile{FileName: fileName, FileData: fileData})
		}
	}
}
