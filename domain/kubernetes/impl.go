package kubernetes

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

type Pod interface {
	GetPod(ctx context.Context, name string) (*corev1.Pod, error)
	List(ctx context.Context) (*corev1.PodList, error)
	Create(ctx context.Context) error
}

type ConfigMap interface {
	Create(ctx context.Context) (*corev1.ConfigMap, error)
}
