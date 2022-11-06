package kubernetes

import (
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/domain/kubernetes"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podImpl struct {
	cfg *config.PodConfig
}

func NewPodImpl(cfg *config.PodConfig) kubernetes.Pod {
	return &podImpl{cfg: cfg}
}

func (p *podImpl) GetPod(ctx context.Context, name string) (*corev1.Pod, error) {
	cli := GetClient()

	return cli.CoreV1().Pods(p.cfg.NameSpace).Get(ctx, name, metav1.GetOptions{})
}

func (p *podImpl) List(ctx context.Context) (*corev1.PodList, error) {
	cli := GetClient()

	return cli.CoreV1().Pods(p.cfg.NameSpace).List(ctx, metav1.ListOptions{})
}

func (p *podImpl) Create(ctx context.Context) error {
	_, err := GetClient().CoreV1().Pods(p.cfg.NameSpace).Create(ctx, p.getPodConf("text"), metav1.CreateOptions{})
	return err
}

func (p *podImpl) getPodConf(name string) *corev1.Pod {
	newPod := &corev1.Pod{}
	newPod.TypeMeta = metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	}
	newPod.ObjectMeta = metav1.ObjectMeta{
		Name:      name,
		Namespace: p.cfg.NameSpace,
	}

	spec := corev1.PodSpec{}
	newPod.Spec = spec

	spec.Containers = []corev1.Container{
		{
			Name:  name,
			Image: p.cfg.Image,
			Env: []corev1.EnvVar{
				{
					Name: "DB_USER",
					ValueFrom: &corev1.EnvVarSource{
						SecretKeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{Name: p.cfg.Secret},
							Key:                  "db-user",
						},
					},
				},
				{
					Name: "DB_PWD",
					ValueFrom: &corev1.EnvVarSource{
						SecretKeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{Name: p.cfg.Secret},
							Key:                  "db-password",
						},
					},
				},
			},
			ImagePullPolicy: corev1.PullAlways,
			Resources:       corev1.ResourceRequirements{},
			//LivenessProbe: &corev1.Probe{
			//	ProbeHandler: corev1.ProbeHandler{
			//		HTTPGet: &corev1.HTTPGetAction{
			//			Path: "callback/" + namespace + "/" + name,
			//			Port: intstr.IntOrString{
			//				Type:   intstr.Int,
			//				IntVal: 8080,
			//			},
			//		},
			//	},
			//	InitialDelaySeconds: 5,  //Pod容器启动多少时间后开始检测
			//	PeriodSeconds:       10, //探测间隔时间
			//	TimeoutSeconds:      3,  //超时时间
			//},
			//Lifecycle: &corev1.Lifecycle{
			//	PostStart: &corev1.LifecycleHandler{
			//		HTTPGet: &corev1.HTTPGetAction{
			//			Path: "callback/" + namespace + "/" + name,
			//			Port: intstr.IntOrString{
			//				Type:   intstr.Int,
			//				IntVal: 8080,
			//			},
			//		},
			//	},
			//},
		},
	}
	spec.RestartPolicy = corev1.RestartPolicyNever
	return newPod
}
