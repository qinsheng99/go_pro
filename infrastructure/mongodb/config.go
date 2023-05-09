package mongodb

type Config struct {
	Host       string `json:"host"`
	Port       int64  `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}
