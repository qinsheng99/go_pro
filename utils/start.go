package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

var (
	m *manager
	l sync.Mutex
)

type manager struct {
	c      *sync.Cond
	signal bool
	wg     sync.WaitGroup
}

type ListenAndServer interface {
	Shutdown(ctx context.Context) error
	ListenAndServe() error
}

func init() {
	m = &manager{
		c:  sync.NewCond(&sync.Mutex{}),
		wg: sync.WaitGroup{},
	}

	go listen()
}

func listen() {
	l.Lock()
	s := signals()
	l.Unlock()
	<-s
	m.c.L.Lock()
	m.signal = true
	m.c.Broadcast()
	m.c.L.Unlock()
}

var signals = func() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	return sig
}

func wait(f func()) {
	m.c.L.Lock()
	if !m.signal {
		m.c.Wait()
	}
	m.c.L.Unlock()
	f()
}

func Start(port int, handler http.Handler) {
	lis := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	go func() {
		m.wg.Wait()
	}()

	defer wait(func() {
		logrus.Info("interrupt")
	})
	listenRun(lis)
}

func listenRun(s ListenAndServer) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		logrus.WithError(s.ListenAndServe()).Info("server exit")
	}()

	go wait(func() {
		logrus.WithError(s.Shutdown(context.Background())).Info("server shutdown")
	})
}
