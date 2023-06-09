package utils

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/rand"

	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

func ErrorNotFound(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func Label(name string) map[string]string {
	return map[string]string{"app": name}
}

func StrSliceToInterface(data []string) []interface{} {
	var res = make([]interface{}, 0, len(data))
	for k := range data {
		if len(data[k]) <= 0 {
			continue
		}
		res = append(res, data[k])
	}

	return res
}

func FilterRepeat(strs []string, str string) (repeat []string) {
	flag := len(str) > 0
	var item = make(map[string]struct{})
	for _, s := range strs {
		if flag && !strings.Contains(s, str) {
			continue
		}
		if _, ok := item[s]; !ok {
			item[s] = struct{}{}
			repeat = append(repeat, s)
		}
	}

	return
}

func GenerateCode(length int) string {
	rand.Seed(time.Now().Unix())
	l := len(_const.Code)
	var bys = new(bytes.Buffer)
	for i := 0; i < length; i++ {
		bys.Write([]byte{_const.Code[rand.Intn(l)]})
	}

	return bys.String()
}

func Now() int64 {
	return time.Now().Unix()
}

func ZeroNow() int64 {
	t := time.Now()

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

func ToDate(n int64) string {
	if n == 0 {
		n = Now()
	}

	return time.Unix(n, 0).Format("2006-01-02")
}

func Date() string {
	return ToDate(Now())
}
