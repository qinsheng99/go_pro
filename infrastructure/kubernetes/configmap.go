package kubernetes

import (
	"context"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/domain/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

type configImpl struct {
	cfg *config.KubernetesConfig
}

func NewConfigImpl(cfg *config.KubernetesConfig) kubernetes.ConfigMap {
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
