// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// RegisterableService can register a service.
type RegisterableService interface{}

// RegisterableGRPCService can register gRPC services.
type RegisterableGRPCService interface {
	// RegisterGRPC registers a gRPC service
	RegisterGRPC(*grpc.Server)
}

// RegisterableRESTService can register REST services.
type RegisterableRESTService interface {
	// RegisterREST registers a REST service
	RegisterREST(mux *runtime.ServeMux, grpcEndpoint string) error
}
