package kubernetes

type Config struct {
	NameSpace string    `json:"namespace"`
	Pod       pod       `json:"pod"`
	ConfigMap configMap `json:"config_map"`
	Crd       crd       `json:"crd"`
}

type pod struct {
	Image  string   `json:"image"`
	Secret string   `json:"secret"`
	Name   string   `json:"name" required:"true"`
	Args   []string `json:"args" required:"true"`
	Port   int32    `json:"port" required:"true"`
}

type configMap struct {
	ConfigMapName string `json:"config_map_name"`
	ConfigName    string `json:"config_name"`
	MounthPath    string `json:"mounth_path"`
}

type crd struct {
	Group   string `json:"group"`
	Kind    string `json:"kind"`
	Version string `json:"version"`
}
