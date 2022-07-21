// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package config

import (
	"github.com/gitpod-io/gitpod/common-go/grpc"
)

type ServiceConfiguration struct {
	Server struct {
		Port int `json:"port"`
		// TLS  struct {
		// 	CA          string `json:"ca"`
		// 	Certificate string `json:"crt"`
		// 	PrivateKey  string `json:"key"`
		// } `json:"tls"`
		RateLimits map[string]grpc.RateLimit `json:"ratelimits"`
	} `json:"server"`

	PProf struct {
		Addr string `json:"addr"`
	} `json:"pprof"`

	Prometheus struct {
		Addr string `json:"addr"`
	} `json:"prometheus"`
}
