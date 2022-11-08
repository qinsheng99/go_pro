package utils

import (
	"errors"
	"github.com/qinsheng99/go-domain-web/api"
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
		page = Page
	} else {
		page = req.Page
	}
	if req.Size == 0 {
		size = Size
	} else {
		size = req.Size
	}
	return page, size
}

func Label(name string) map[string]string {
	return map[string]string{"app": name}
}
