package dp

type dpPage int

type Page interface {
	Page() int
}

type Size interface {
	Size() int
}

type dpSize int

func NewPage(page int) Page {
	if page <= 0 {
		page = 1
	}

	return dpPage(page)
}

func NewSize(size int) Size {
	if size <= 0 {
		size = 10
	}

	return dpSize(size)
}

func (p dpPage) Page() int {
	return int(p)
}

func (p dpSize) Size() int {
	return int(p)
}
