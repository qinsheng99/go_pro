package kubernetes

import (
	"fmt"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"os/user"
)

var (
	k8sConfig *rest.Config
	k8sClient *kubernetes.Clientset
	dyna      dynamic.Interface
	restm     *restmapper.DeferredDiscoveryRESTMapper
)

func getHome() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}

	return u.HomeDir
}

func Init() (err error) {
	k8sConfig, err = clientcmd.BuildConfigFromFlags("", getHome()+"/.kube/config")
	if err != nil {
		return
	}

	k8sClient, err = kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return
	}
	dyna, err = dynamic.NewForConfig(k8sConfig)
	if err != nil || dyna == nil {
		err = fmt.Errorf("err is %v, %v", err, "dyna is nil")
		return
	}

	dis, err := discovery.NewDiscoveryClientForConfig(k8sConfig)
	if err != nil {
		return
	}

	restm = restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dis))
	return nil
}

func GetClient() *kubernetes.Clientset {
	return k8sClient
}

func GetDyna() dynamic.Interface {
	return dyna
}

func GetrestMapper() *restmapper.DeferredDiscoveryRESTMapper {
	return restm
}

func GetK8sConfig() *rest.Config {
	return k8sConfig
}
