package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lunny/log"
)

//DispHex 十六进制显示数据
func DispHex(msg string, buf []byte) {
	if len(msg) > 0 {
		fmt.Println(msg)
	}
	Length := len(buf)
	for i := 0; i < Length/16; i++ {
		fmt.Printf("%04X", i*16)
		for j := 0; j < 16; j++ {
			if j == 8 {
				fmt.Printf("-%02X", buf[i*16+j])
			} else {
				fmt.Printf(" %02X", buf[i*16+j])
			}
		}
		fmt.Printf("  ")
		for j := 0; j < 16; j++ {
			if buf[i*16+j] >= 32 && buf[i*16+j] < 128 {
				fmt.Printf("%c", buf[i*16+j])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
	if (Length % 16) > 0 {
		fmt.Printf("%04X", Length/16*16)
	}
	for j := 0; j < 16; j++ {
		if j < (Length % 16) {
			if j == 8 {
				fmt.Printf("-%02X", buf[(Length/16)*16+j])
			} else {
				fmt.Printf(" %02X", buf[(Length/16)*16+j])
			}
		} else {
			fmt.Printf("   ")
		}
	}
	fmt.Printf("  ")
	for j := 0; j < (Length % 16); j++ {
		if buf[(Length/16)*16+j] >= 32 && buf[(Length/16)*16+j] < 128 {
			fmt.Printf("%c", buf[(Length/16)*16+j])
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println("")
}

//WriteToFile 写buf到一个数组里
func WriteToFile(buf []byte, filename string) error {
	fo, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fo.Write(buf)
	fo.Close()
	return nil
}

//IsHexStr 判定一个字符串是不是由16进制字符组成
func IsHexStr(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9' || str[i] >= 'a' && str[i] <= 'f' || str[i] >= 'A' && str[i] <= 'F') {
			return false
		}
	}
	return true
}

//ReadToArray 读文本文件到一个string array里，每个string已经去掉了'\r','\n'字符
func ReadToArray(file string) ([]string, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	strs := strings.Split(string(buf), "\n")
	return strs, nil
}

//ReadToArrayAndTrim 读文本文件到一个string array里，每个string已经去掉了'\r','\n'字符
//并且去掉了空行
func ReadToArrayAndTrim(file string) ([]string, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	strs := strings.Split(string(buf), "\n")
	r := []string{}
	for _, v := range strs {
		t := strings.TrimSpace(v)
		if t != "" {
			r = append(r, t)
		}
	}
	return r, nil
}

//IsLocalIP 判断是不是本地IP
func IsLocalIP(IP string) bool {
	pos := strings.LastIndex(IP, ":")
	if pos == -1 {
		return IP == "127.0.0.1"
	}

	str := IP[:pos]
	if str == "127.0.0.1" {
		return true
	}

	//[::1]:47764
	if str == "[::1]" {
		return true
	}

	return false
}

//IPData --
type IPData struct {
	Country   string `json:"country"`    //国家
	CountryID string `json:"country_id"` //代码
	Area      string `json:"area"`       //
	AreaID    string `json:"area_id"`    //
	Region    string `json:"region"`     //
	RegionID  string `json:"region_id"`  //
	City      string `json:"city"`       //
	CityID    string `json:"city_id"`    //
	County    string `json:"county"`     //
	CountyID  string `json:"county_id"`  //
	ISP       string `json:"isp"`        //
	ISPID     string `json:"isp_id"`     //
	IP        string `json:"ip"`         //
}

/*
{"code":0,"data":{"country":"\u9a6c\u6765\u897f\u4e9a","country_id":"MY","area":"","area_id":"","region":"","region_id":"","city":"","city_id":"","county":"","county_id":"","isp":"","isp_id":"","ip":"121.122.132.122"}}
*/
type IPInfo struct {
	Code int    `json:"code"` //返回的错误代码，0--OK
	Data IPData `json:"data"` //IP数据
}

//GetIPInfo 根据IP获取地理信息（国家、省、市等）
//返回一个json
func GetIPInfo(IP string) *IPInfo {
	pos := strings.LastIndex(IP, ":")
	var queryIP string
	if pos == -1 {
		queryIP = IP
	} else {
		queryIP = IP[:pos]
	}

	resp, err := http.Get(fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", queryIP))
	if err != nil {
		return nil
	}

	buf, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil
	}

	r := &IPInfo{}

	err = json.Unmarshal(buf, r)
	if err != nil {
		return nil
	}

	return r
}

//IsFileExist 判断一个文件是不是存在
func IsFileExist(FileName string) bool {
	_, err := os.Stat(FileName)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}

//FileMustExist 确保一个文件存在
func FileMustExist(FileName string) {
	_, err := os.Stat(FileName)
	if err == nil || os.IsExist(err) {
		return
	}
	tmpFile, err := os.Create(FileName)
	if err != nil {
		log.Error("create file error:" + err.Error())
	}
	defer tmpFile.Close()
}

//GetFileSize 取文件长度
func GetFileSize(FileName string) (int64, error) {
	info, err := os.Stat(FileName)
	if err == nil || os.IsExist(err) {
		return info.Size(), nil
	}
	return -1, errors.New("file does not exist")
}

//Md5 计算一个bufffer的MD5值
func Md5(buf []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(buf)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//FileMd5 计算一个文件的MD5
func FileMd5(file string) (string, error) {
	size, err := GetFileSize(file)
	if err != nil {
		return "", err
	}

	n := int(size / (1024 * 1024))
	r := int(size % (1024 * 1024))

	fp, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	buf := make([]byte, 1024*1024)

	md5Ctx := md5.New()
	for i := 0; i < n; i++ {
		nRead, err := fp.Read(buf)
		if err != nil || nRead < 1024*1024 {
			return "", err
		}
		md5Ctx.Write(buf)
	}

	_, err = fp.Read(buf[:r])
	md5Ctx.Write(buf[:r])
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr), nil
}

//JSONSort sort json string by the key.
func JSONSort(jsonStr string) string {
	mapJSON := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &mapJSON)
	if err != nil {
		log.Println("err 01:", err)
		return jsonStr
	}

	buf, err := json.Marshal(mapJSON)
	if err != nil {
		log.Println("err 02:", err)
		return jsonStr
	}
	return string(buf)
}

//IsNotExist windows下的os.IsNotExist有问题，不能识别"not found"
func IsNotExist(err error) bool {
	if os.IsNotExist(err) {
		return true
	}
	if strings.Contains(err.Error(), "not found") {
		return true
	}
	if strings.Contains(err.Error(), "does not exist") {
		return true
	}

	return false
}


const charTbl = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ00112233445566778899"

var rnd *rand.Rand

//GenRandomPass --
func GenRandomPass(length int) string {

	buf := bytes.Buffer{}
	tblLength := len(charTbl)
	for i := 0; i < length; i++ {
		buf.WriteByte(charTbl[rnd.Intn(tblLength)])
	}

	return buf.String()
}

const digiTbl = "0123456789"

//GenRandomDigiCode -- 生成全数字的串，一般是做验证码用
func GenRandomDigiCode(length int) string {

	buf := bytes.Buffer{}
	tblLength := len(digiTbl)
	for i := 0; i < length; i++ {
		buf.WriteByte(digiTbl[rnd.Intn(tblLength)])
	}

	return buf.String()
}

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}
