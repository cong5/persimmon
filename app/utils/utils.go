package utils

import (
	"io"
	"os"
	"fmt"
	"net"
	"math"
	"time"
	"bytes"
	"regexp"
	"strconv"
	"reflect"
	"strings"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	mathRand "math/rand"
	"github.com/cong5/persimmon/app/models"
)

/**
Date Time Format:
	"2006-01-02"
	"2006-01-02 15:04:05"
	"Mon, 2 Jan 2006 15:04:05"
*/

// md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Guid
func NewGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func HashFile(path string) (hashSum string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hashVal := h.Sum(nil)
	return string(hashVal), nil
}

func ToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}

func Urlencode(data map[string]string) string {
	var buf bytes.Buffer
	for k, v := range data {
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0 : len(s)-1]
}

func SQLFilters(sql string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return re.MatchString(sql)
}

func IntJoin(intArr []int, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(intArr), " ", delimiter, -1), "[]")
}

func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func GetTotalPage(count int, limit int) int {
	return int(math.Ceil(float64(count) / float64(limit)))
}

func IntDefault(check bool, inputVal int, defaultVal int) int {
	newVal := defaultVal
	if check {
		newVal = inputVal
	}
	return newVal
}

func StringDefault(check bool, inputVal string, defaultVal string) string {
	newVal := defaultVal
	if check {
		newVal = inputVal
	}
	return newVal
}

func Date(format string, timestamp int64) (string) {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func SliceIntArr(arr []int, start int, length int, total int) []int {
	intArr := make([]int, 0)
	if total > start && total > length {
		intArr = arr[start:length]
	} else if total <= length {
		intArr = arr[start:total]
	}

	return intArr
}

func CacheKey(keys ...interface{}) string {
	strArr := make([]string, len(keys))
	for k, v := range keys {
		vtype := reflect.TypeOf(v).String()
		switch vtype {
		case "string":
			strArr[k] = v.(string)
		case "int":
			strArr[k] = strconv.Itoa(v.(int))
		default:
		}
	}
	newStr := strings.Join(strArr, ":")
	return fmt.Sprintf("persimmon:%s", newStr)
}

func Random(max int) int {
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func Substr(src string, length int) string {
	if length <= 0 {
		length = 125
	}
	newSrc := TrimHtml(src)
	newStr := []rune(newSrc)
	strLen := len(newStr)
	var desc string
	if strLen > length {
		desc = string(newStr[:length])
	} else {
		desc = newSrc
	}
	return desc
}

func SubstrContent(posts []models.Posts) []models.Posts {
	for key, val := range posts {
		posts[key].Content = Substr(val.Content, 130)
	}
	return posts
}

func CommentRelolver(comments []models.Comments) []models.Comments {
	pat := `@[^\s]+\s?`
	re, _ := regexp.Compile(pat)
	reFun := func(s string) string {
		return fmt.Sprintf("<a href='[anchor]' rel='external nofollow'>%s</a>", s)
	}

	for key, val := range comments {
		//email md5
		comments[key].Md5 = Md5(val.Email)
		comments[key].Email = "Privacy"
		//解析@username，暂时只支持@一个人
		newContent := re.ReplaceAllStringFunc(val.Content, reFun)
		anchor := fmt.Sprintf("#%d", val.Pid)
		comments[key].Content = strings.Replace(newContent, "[anchor]", anchor, 1)
	}
	return comments
}
