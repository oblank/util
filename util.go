package util

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	//	"fmt"
)

// The IsNum judges string is number or not.
func StringIsNumber(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

//date
func Md5String(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))

	//	bytes := md5.Sum([]byte(str))
	//	return hex.EncodeToString(bytes[:])
}

func Md5Byte(src []byte) string {
	data := md5.Sum(src)
	return hex.EncodeToString(data[:])
}

//time
func UnixTimestamp() int {
	return int(time.Now().Unix())
}

func UnixTimestampStr() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func TimestampToDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func DateToTimestamp(date string, dateformat string) int {
	tm, _ := time.Parse(dateformat, date)
	return int(tm.Unix())
}

//range like php range
func RangeNumber(start int, end int) []int {
	var ret []int
	if start <= end {
		len := end - start
		value := start
		for i := 0; i <= len; i++ {
			ret = append(ret, value)
			value++
		}
	} else {
		//start > end
		len := start - end
		value := start
		for i := 0; i <= len; i++ {
			ret = append(ret, value)
			value--
		}
	}
	return ret
}

func RangeStringWithSort(sort_slice []string, start string, end string) []string {
	sort.Strings(sort_slice)
	startIndex := sort.SearchStrings(sort_slice, start)
	endIndex := sort.SearchStrings(sort_slice, end)
	if (startIndex == 0 && start != sort_slice[0]) || (endIndex == 0 && end != sort_slice[0]) {
		return []string{}
	}

	if startIndex < endIndex {
		return sort_slice[startIndex : endIndex+1]
	}

	if startIndex == endIndex {
		return sort_slice[startIndex : endIndex+1]
	}

	return sort_slice[endIndex : startIndex+1]
}

func Rangeabc(start string, end string) []string {
	if len(start) > 1 || len(end) > 1 {
		return []string{}
	}
	start = strings.ToLower(start)
	end = strings.ToLower(end)
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	return RangeStringWithSort(alphabets, start, end)
}

func RangeABC(start string, end string) []string {
	if len(start) > 1 || len(end) > 1 {
		return []string{}
	}
	start = strings.ToUpper(start)
	end = strings.ToUpper(end)
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	return RangeStringWithSort(alphabets, start, end)
}

func RandomString() string {
	id := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		panic(err) // This shouldn't happen
	}
	return hex.EncodeToString(id)
}

func OrdStr(char string) int {
	if char == "" || len(char) > 1 {
		panic("is not corret string")
	}
	val := []byte(char)
	return int(val[0])
}

func Ord(char byte) int {
	return int(char)
}

func Char(ascii int) string {
	return string(IntToBytes(ascii))
}

func ByteToString(ascii byte) string {
	//ascii = bytes.Trim([]byte{ascii}, "\x00") //trim NUL character
	return string(ascii)
}

func IntToBytes(i int) []byte {
	var buf = make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	buf = bytes.Trim(buf, "\x00") //trim NUL character
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

func MbSubstr(source string, pos int, length int) string {
	runes := []rune(source)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func MbStrlen(source string) int {
	runes := []rune(source)
	return len(runes)
}

func Base64Encode(data string) string {
	return Base64EncodeBytes([]byte(data))
}

func Base64EncodeBytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) string {
	ret, _ := base64.StdEncoding.DecodeString(data)
	return string(ret)
}

func Base64Encode4Url(data string) string {
	return Base64Encode4UrlBytes([]byte(data))
}

func Base64Encode4UrlBytes(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base64Decode4Url(data string) string {
	ret, _ := base64.RawURLEncoding.DecodeString(data)
	return string(ret)
}
