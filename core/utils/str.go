package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ver1 的版本是否大于ver2
func CheckVersion(ver1, ver2 string) bool {
	vs1 := strings.Split(ver1, ".")
	vs2 := strings.Split(ver2, ".")
	ls := max(len(vs1), len(vs2))
	v1 := float64(0)
	v2 := float64(0)
	for i, v := range vs1 {
		vv, _ := strconv.ParseFloat(v, 10)
		v1 += math.Pow10((ls-i-1)*2) * vv
	}
	for i, v := range vs2 {
		vv, _ := strconv.ParseFloat(v, 10)
		v2 += math.Pow10((ls-i-1)*2) * vv
	}
	return v1 > v2
}

func BuildIntsToString(ids ...uint64) string {
	ret := []string{}
	for _, i := range ids {
		ret = append(ret, fmt.Sprint(i))
	}
	return strings.Join(ret, ",")
}

// 生成随机字符串
func RandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// 隐藏手机号
func HideMobile(mobile string) string {
	return fmt.Sprintf("%s****%s", mobile[0:3], mobile[7:])
}

// 隐藏邮箱
func HideEmail(email string) string {
	i := strings.Index(email, "@")
	name := email[:i]
	host := email[i:]
	ln := len(name)
	if ln == 1 {
		return fmt.Sprintf("%s****%s", name, host)
	} else if ln > 5 {
		return fmt.Sprintf("%s****%s%s", name[:3], name[ln-3:], host)
	} else {
		return fmt.Sprintf("%s****%s%s", name[:ln-2], host)
	}

}
