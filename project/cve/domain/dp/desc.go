package dp

type description string

func NewDescription(desc string) Description {
	return description(desc)
}

type Description interface {
	Description() string
}

func (d description) Description() string {
	return string(d)
}
