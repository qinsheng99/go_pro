package kubernetes

import (
	"context"
	"log"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func NewListen(res *kubernetes.Clientset, dym dynamic.Interface, resource schema.GroupVersionResource, listen bool) ListenInter {
	return &Listen{res: res, wg: &sync.WaitGroup{}, mux: &sync.Mutex{}, dym: dym, resource: resource, listen: listen}
}

type ListenInter interface {
	ListenResource()
}

type Listen struct {
	wg       *sync.WaitGroup
	res      *kubernetes.Clientset
	resync   time.Duration
	mux      *sync.Mutex
	dym      dynamic.Interface
	resource schema.GroupVersionResource
	listen   bool
}

func (l *Listen) ListenResource() {
	if !l.listen {
		return
	}
	log.Println("listen k8s resource for crd")
	infor := l.crdConfig()
	infor.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: l.Update,
		DeleteFunc: l.Delete,
		AddFunc:    l.Add,
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	infor.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, infor.HasSynced) {
		log.Println("cache sync err")
		return
	}

	//l.infor = infor

	<-stopCh
}

func (l *Listen) Update(oldObj, newObj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(newObj)
	if err != nil {
		log.Println("update func err: ", err.Error())
	}
	log.Println("update func key: ", key)
}

func (l *Listen) Delete(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		log.Println("delete func err: ", err.Error())
	}
	log.Println("delete func key: ", key)
}

func (l *Listen) Add(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		log.Println("add func err: ", err.Error())
	}
	log.Println("add func key: ", key)
}

func (l *Listen) crdConfig() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return l.dym.Resource(l.resource).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return l.dym.Resource(l.resource).Watch(context.TODO(), options)
			},
		},
		&unstructured.Unstructured{},
		0,
		cache.Indexers{},
	)
}

func (l *Listen) infor() {
	inforFac := informers.NewSharedInformerFactory(l.res, 0)
	inforFac.Core().V1().Pods().Informer().AddEventHandler(nil)
}
