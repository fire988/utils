package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/saintfish/chardet"
)

var wi = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var tb2 = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

//IsValidPersonID -- 检验身份证合法性
func IsValidPersonID(idStr string) bool {

	if len(idStr) != 18 {
		return false
	}

	for i := 0; i < 17; i++ {
		if idStr[i] < '0' || idStr[i] > '9' {
			return false
		}
	}

	if idStr[6:8] != "19" && idStr[6:8] != "20" {
		return false
	}

	switch idStr[10:12] {
	case "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12":
	default:
		return false
	}

	d, _ := strconv.Atoi(idStr[12:14])

	if d < 1 || d > 31 {
		return false
	}

	id := []byte(idStr)

	var res int
	for i := 0; i < 17; i++ {
		res += int(id[i]-'0') * wi[i]
	}

	idx := res % 11

	ch := id[17]
	if ch == 'x' {
		ch = 'X'
	}

	return tb2[idx] == ch
}

//IsValidEmail --
func IsValidEmail(mail string) bool {
	s := strings.Split(mail, "@")
	if len(s) != 2 {
		return false
	}
	if strings.Contains(s[1], ".") == false {
		return false
	}

	return true
}

//IsUtf8File --
func IsUtf8File(file string) bool {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("open file error:", err.Error())
		return false
	}
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(buf)
	if err != nil {
		log.Println("detect error:", err.Error())
		return false
	}

	if result.Charset == "UTF-8" && result.Confidence > 50 {
		return true
	}

	return false
}
