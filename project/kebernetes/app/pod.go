package app

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/qinsheng99/go-domain-web/project/kebernetes/domain/kubernetes"
)

type PodServiceImpl interface {
	GetPod(ctx context.Context, name string) (*corev1.Pod, error)
	PodList(ctx context.Context) (data []map[string]interface{}, _ error)
	Create(ctx context.Context) error
}

type podService struct {
	p kubernetes.Pod
}

func NewPodService(p kubernetes.Pod) PodServiceImpl {
	return &podService{p: p}
}

func (p *podService) GetPod(ctx context.Context, name string) (*corev1.Pod, error) {
	return p.p.GetPod(ctx, name)
}

func (p *podService) PodList(ctx context.Context) (data []map[string]any, _ error) {
	podlist, err := p.p.List(ctx)
	if err != nil {
		return nil, err
	}

	data = make([]map[string]any, 0, len(podlist.Items))

	for _, item := range podlist.Items {
		data = append(data, map[string]any{
			"name":      item.GetName(),
			"namespace": item.GetNamespace(),
		})
	}

	return data, err

}

func (p *podService) Create(ctx context.Context) error {
	return p.p.Create(ctx)
}
