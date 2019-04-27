package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

//MultiPartFile --
type MultiPartFile struct {
	FileName string
	FileData []byte
}

//MultiPart --
type MultiPart struct {
	Files []MultiPartFile
	Forms map[string]string
}

//GetMultiParts --
func GetMultiParts(r *http.Request) (MultiPart, error) {
	mp := MultiPart{Files: []MultiPartFile{}, Forms: map[string]string{}}

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
			mp.Forms[formName] = string(formValue)
		}
		if fileName != "" {
			fileData, _ := ioutil.ReadAll(p)
			mp.Files = append(mp.Files, MultiPartFile{FileName: fileName, FileData: fileData})
		}
	}
}
