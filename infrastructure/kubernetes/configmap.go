package kubernetes

import (
	"context"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/qinsheng99/go-domain-web/project/kubernetes/domain/kubernetes"
)

type configImpl struct {
	cfg *Config
}

func NewConfigImpl(cfg *Config) kubernetes.ConfigMap {
	return &configImpl{cfg: cfg}
}

func (c *configImpl) Create(ctx context.Context) (*corev1.ConfigMap, error) {
	return GetClient().
		CoreV1().
		ConfigMaps(c.cfg.NameSpace).
		Create(ctx, c.getConfigmap(), metav1.CreateOptions{})
}

func (c *configImpl) getConfigmap() *corev1.ConfigMap {
	configmap := &corev1.ConfigMap{}

	configmap.APIVersion = "v1"
	configmap.Kind = "ConfigMap"
	configmap.Name = "test-config"
	configmap.Namespace = "default"
	yamldata, err := os.ReadFile("./template/config.yaml")
	if err != nil {
		return nil
	}

	configmap.BinaryData = map[string][]byte{
		"config.yaml": yamldata,
	}

	return configmap
}
