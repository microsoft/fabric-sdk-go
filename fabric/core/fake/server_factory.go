// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package fake

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ServerFactory is a fake server for instances of the core.ClientFactory type.
type ServerFactory struct {
	CapacitiesServer                CapacitiesServer
	DeploymentPipelinesServer       DeploymentPipelinesServer
	ExternalDataSharesServer        ExternalDataSharesServer
	GitServer                       GitServer
	ItemsServer                     ItemsServer
	JobSchedulerServer              JobSchedulerServer
	LongRunningOperationsServer     LongRunningOperationsServer
	ManagedPrivateEndpointsServer   ManagedPrivateEndpointsServer
	OneLakeDataAccessSecurityServer OneLakeDataAccessSecurityServer
	OneLakeShortcutsServer          OneLakeShortcutsServer
	WorkspacesServer                WorkspacesServer
}

// ServerFactoryTransport connects instances of core.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                               *ServerFactory
	trMu                              sync.Mutex
	trCapacitiesServer                *CapacitiesServerTransport
	trDeploymentPipelinesServer       *DeploymentPipelinesServerTransport
	trExternalDataSharesServer        *ExternalDataSharesServerTransport
	trGitServer                       *GitServerTransport
	trItemsServer                     *ItemsServerTransport
	trJobSchedulerServer              *JobSchedulerServerTransport
	trLongRunningOperationsServer     *LongRunningOperationsServerTransport
	trManagedPrivateEndpointsServer   *ManagedPrivateEndpointsServerTransport
	trOneLakeDataAccessSecurityServer *OneLakeDataAccessSecurityServerTransport
	trOneLakeShortcutsServer          *OneLakeShortcutsServerTransport
	trWorkspacesServer                *WorkspacesServerTransport
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of core.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	// client := method[:strings.Index(method, ".")]
	client := strings.Split(method, ".")[1]
	var resp *http.Response
	var err error

	switch client {
	case "CapacitiesClient":
		initServer(s, &s.trCapacitiesServer, func() *CapacitiesServerTransport { return NewCapacitiesServerTransport(&s.srv.CapacitiesServer) })
		resp, err = s.trCapacitiesServer.Do(req)
	case "DeploymentPipelinesClient":
		initServer(s, &s.trDeploymentPipelinesServer, func() *DeploymentPipelinesServerTransport {
			return NewDeploymentPipelinesServerTransport(&s.srv.DeploymentPipelinesServer)
		})
		resp, err = s.trDeploymentPipelinesServer.Do(req)
	case "ExternalDataSharesClient":
		initServer(s, &s.trExternalDataSharesServer, func() *ExternalDataSharesServerTransport {
			return NewExternalDataSharesServerTransport(&s.srv.ExternalDataSharesServer)
		})
		resp, err = s.trExternalDataSharesServer.Do(req)
	case "GitClient":
		initServer(s, &s.trGitServer, func() *GitServerTransport { return NewGitServerTransport(&s.srv.GitServer) })
		resp, err = s.trGitServer.Do(req)
	case "ItemsClient":
		initServer(s, &s.trItemsServer, func() *ItemsServerTransport { return NewItemsServerTransport(&s.srv.ItemsServer) })
		resp, err = s.trItemsServer.Do(req)
	case "JobSchedulerClient":
		initServer(s, &s.trJobSchedulerServer, func() *JobSchedulerServerTransport { return NewJobSchedulerServerTransport(&s.srv.JobSchedulerServer) })
		resp, err = s.trJobSchedulerServer.Do(req)
	case "LongRunningOperationsClient":
		initServer(s, &s.trLongRunningOperationsServer, func() *LongRunningOperationsServerTransport {
			return NewLongRunningOperationsServerTransport(&s.srv.LongRunningOperationsServer)
		})
		resp, err = s.trLongRunningOperationsServer.Do(req)
	case "ManagedPrivateEndpointsClient":
		initServer(s, &s.trManagedPrivateEndpointsServer, func() *ManagedPrivateEndpointsServerTransport {
			return NewManagedPrivateEndpointsServerTransport(&s.srv.ManagedPrivateEndpointsServer)
		})
		resp, err = s.trManagedPrivateEndpointsServer.Do(req)
	case "OneLakeDataAccessSecurityClient":
		initServer(s, &s.trOneLakeDataAccessSecurityServer, func() *OneLakeDataAccessSecurityServerTransport {
			return NewOneLakeDataAccessSecurityServerTransport(&s.srv.OneLakeDataAccessSecurityServer)
		})
		resp, err = s.trOneLakeDataAccessSecurityServer.Do(req)
	case "OneLakeShortcutsClient":
		initServer(s, &s.trOneLakeShortcutsServer, func() *OneLakeShortcutsServerTransport {
			return NewOneLakeShortcutsServerTransport(&s.srv.OneLakeShortcutsServer)
		})
		resp, err = s.trOneLakeShortcutsServer.Do(req)
	case "WorkspacesClient":
		initServer(s, &s.trWorkspacesServer, func() *WorkspacesServerTransport { return NewWorkspacesServerTransport(&s.srv.WorkspacesServer) })
		resp, err = s.trWorkspacesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
