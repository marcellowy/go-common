// Package server
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package server

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	defaultServerStopWaitDuration = 500 * time.Millisecond
)

func init() {
	if v := config.GetInt64("server.stopWaitMillisecond"); v > 0 {
		defaultServerStopWaitDuration = time.Millisecond * time.Duration(v)
	}
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		quit: make(chan os.Signal),
	}
}

type Server struct {
	http.Server
	quit chan os.Signal
}

func (s *Server) start() {
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (s *Server) wait() {
	signal.Notify(s.quit, os.Interrupt,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGTERM)

exit:
	for {
		select {
		case <-s.quit:
			fmt.Println("cc")
			s.waitStop()
			break exit
		}
	}
}

func (s *Server) waitStop() {
	ctx, cancel := context.WithTimeout(context.TODO(), defaultServerStopWaitDuration)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func (s *Server) Start() {
	go s.start()
	s.wait()
}
