package dp

type description string

func NewDescription(desc string) PackageDescription {
	return description(desc)
}

func NewCveDescription(desc string) CveDescription {
	return description(desc)
}

type PackageDescription interface {
	PackageDescription() string
}

type CveDescription interface {
	CveDescription() string
}

func (d description) PackageDescription() string {
	return string(d)
}

func (d description) CveDescription() string {
	return string(d)
}
