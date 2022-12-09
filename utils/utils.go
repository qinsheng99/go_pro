package utils

import (
	"errors"
	"strings"

	"github.com/qinsheng99/go-domain-web/api"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
	"gorm.io/gorm"
)

func ErrorNotFound(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func GetPage(req api.Pages) (int, int) {
	var page, size int
	if req.Page == 0 {
		page = _const.Page
	} else {
		page = req.Page
	}
	if req.Size == 0 {
		size = _const.Size
	} else {
		size = req.Size
	}
	return page, size
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
