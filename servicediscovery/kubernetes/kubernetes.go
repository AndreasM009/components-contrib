// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kubernetes

import (
	"fmt"

	"github.com/dapr/components-contrib/servicediscovery"
	"github.com/dapr/dapr/pkg/logger"
)

type resolver struct {
	logger logger.Logger
}

func NewKubernetesResolver(logger logger.Logger) servicediscovery.Resolver {
	return &resolver{logger: logger}
}

func (z *resolver) ResolveID(req servicediscovery.ResolveRequest) (string, error) {
	// Dapr requires this formatting for Kubernetes services
	return fmt.Sprintf("%s-dapr.%s.svc.cluster.local:%d", req.ID, req.Namespace, req.Port), nil
}
