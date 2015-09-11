package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

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
