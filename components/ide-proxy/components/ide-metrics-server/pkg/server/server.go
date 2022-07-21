// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/ide-metrics-api/config"
	"github.com/gorilla/websocket"
	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type IDEMetricsServer struct {
}

func New() *IDEMetricsServer {
	return &IDEMetricsServer{}
}

func (s *IDEMetricsServer) Run(cfg *config.ServiceConfiguration) {
	apiServices := []RegisterableService{}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.WithError(err).Fatal("cannot start api endpoint")
	}
	m := cmux.New(l)
	restMux := grpcruntime.NewServeMux()
	grpcMux := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcEndpoint := fmt.Sprintf("localhost:%d", cfg.Server.Port)
	for _, reg := range apiServices {
		if reg, ok := reg.(RegisterableGRPCService); ok {
			reg.RegisterGRPC(grpcServer)
		}
		if reg, ok := reg.(RegisterableRESTService); ok {
			err := reg.RegisterREST(restMux, grpcEndpoint)
			if err != nil {
				log.WithError(err).Fatal("cannot register REST service")
			}
		}
	}
	go grpcServer.Serve(grpcMux)

	httpMux := m.Match(cmux.HTTP1Fast())
	routes := http.NewServeMux()
	grpcWebServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true), grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}))

	routes.Handle("/metrics-api/", http.StripPrefix("/metrics-api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") || websocket.IsWebSocketUpgrade(r) {
			grpcWebServer.ServeHTTP(w, r)
		} else {
			restMux.ServeHTTP(w, r)
		}
	})))
	go http.Serve(httpMux, routes)

	go m.Serve()
}
