package util

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// The IsNum judges string is number or not.
func StringIsNumber(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

//date
func Md5String(str string) string {
	bytes := md5.Sum([]byte(str))
	return hex.EncodeToString(bytes[:])
}

func Md5Byte(src []byte) string {
	data := md5.Sum(src)
	return hex.EncodeToString(data[:])
}

//time
func UnixTimestamp() int64 {
	return time.Now().Unix()
}

func UnixTimestampStr() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func TimestampToDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func DateToTimestamp(date string, dateformat string) int64 {
	tm, _ := time.Parse(dateformat, date)
	return tm.Unix()
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

func RangeString(sort_slice []string, start string, end string) []string {
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

	return RangeString(alphabets, start, end)
}

func RangeABC(start string, end string) []string {
	if len(start) > 1 || len(end) > 1 {
		return []string{}
	}
	start = strings.ToUpper(start)
	end = strings.ToUpper(end)
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	return RangeString(alphabets, start, end)
}
