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

	adminfake "github.com/microsoft/fabric-sdk-go/fabric/admin/fake"
	apacheairflowjobfake "github.com/microsoft/fabric-sdk-go/fabric/apacheairflowjob/fake"
	copyjobfake "github.com/microsoft/fabric-sdk-go/fabric/copyjob/fake"
	corefake "github.com/microsoft/fabric-sdk-go/fabric/core/fake"
	dashboardfake "github.com/microsoft/fabric-sdk-go/fabric/dashboard/fake"
	dataflowfake "github.com/microsoft/fabric-sdk-go/fabric/dataflow/fake"
	datamartfake "github.com/microsoft/fabric-sdk-go/fabric/datamart/fake"
	datapipelinefake "github.com/microsoft/fabric-sdk-go/fabric/datapipeline/fake"
	digitaltwinbuilderfake "github.com/microsoft/fabric-sdk-go/fabric/digitaltwinbuilder/fake"
	digitaltwinbuilderflowfake "github.com/microsoft/fabric-sdk-go/fabric/digitaltwinbuilderflow/fake"
	environmentfake "github.com/microsoft/fabric-sdk-go/fabric/environment/fake"
	eventhousefake "github.com/microsoft/fabric-sdk-go/fabric/eventhouse/fake"
	eventstreamfake "github.com/microsoft/fabric-sdk-go/fabric/eventstream/fake"
	graphqlapifake "github.com/microsoft/fabric-sdk-go/fabric/graphqlapi/fake"
	kqldashboardfake "github.com/microsoft/fabric-sdk-go/fabric/kqldashboard/fake"
	kqldatabasefake "github.com/microsoft/fabric-sdk-go/fabric/kqldatabase/fake"
	kqlquerysetfake "github.com/microsoft/fabric-sdk-go/fabric/kqlqueryset/fake"
	lakehousefake "github.com/microsoft/fabric-sdk-go/fabric/lakehouse/fake"
	mirroredazuredatabrickscatalogfake "github.com/microsoft/fabric-sdk-go/fabric/mirroredazuredatabrickscatalog/fake"
	mirroreddatabasefake "github.com/microsoft/fabric-sdk-go/fabric/mirroreddatabase/fake"
	mirroredwarehousefake "github.com/microsoft/fabric-sdk-go/fabric/mirroredwarehouse/fake"
	mlexperimentfake "github.com/microsoft/fabric-sdk-go/fabric/mlexperiment/fake"
	mlmodelfake "github.com/microsoft/fabric-sdk-go/fabric/mlmodel/fake"
	mounteddatafactoryfake "github.com/microsoft/fabric-sdk-go/fabric/mounteddatafactory/fake"
	notebookfake "github.com/microsoft/fabric-sdk-go/fabric/notebook/fake"
	paginatedreportfake "github.com/microsoft/fabric-sdk-go/fabric/paginatedreport/fake"
	reflexfake "github.com/microsoft/fabric-sdk-go/fabric/reflex/fake"
	reportfake "github.com/microsoft/fabric-sdk-go/fabric/report/fake"
	semanticmodelfake "github.com/microsoft/fabric-sdk-go/fabric/semanticmodel/fake"
	sparkfake "github.com/microsoft/fabric-sdk-go/fabric/spark/fake"
	sparkjobdefinitionfake "github.com/microsoft/fabric-sdk-go/fabric/sparkjobdefinition/fake"
	sqldatabasefake "github.com/microsoft/fabric-sdk-go/fabric/sqldatabase/fake"
	sqlendpointfake "github.com/microsoft/fabric-sdk-go/fabric/sqlendpoint/fake"
	variablelibraryfake "github.com/microsoft/fabric-sdk-go/fabric/variablelibrary/fake"
	warehousefake "github.com/microsoft/fabric-sdk-go/fabric/warehouse/fake"
	warehousesnapshotfake "github.com/microsoft/fabric-sdk-go/fabric/warehousesnapshot/fake"
)

// ServerFactory is a fake server for instance of the fabric.Client type.
type ServerFactory struct {
	Admin                          adminfake.ServerFactory
	Apache                         apacheairflowjobfake.ServerFactory
	CopyJob                        copyjobfake.ServerFactory
	Core                           corefake.ServerFactory
	Dashboard                      dashboardfake.ServerFactory
	Dataflow                       dataflowfake.ServerFactory
	Datamart                       datamartfake.ServerFactory
	DataPipeline                   datapipelinefake.ServerFactory
	DigitalTwinBuilder             digitaltwinbuilderfake.ServerFactory
	DigitalTwinBuilderFlow         digitaltwinbuilderflowfake.ServerFactory
	Environment                    environmentfake.ServerFactory
	Eventhouse                     eventhousefake.ServerFactory
	Eventstream                    eventstreamfake.ServerFactory
	GraphQLApi                     graphqlapifake.ServerFactory
	KQLDashboard                   kqldashboardfake.ServerFactory
	KQLDatabase                    kqldatabasefake.ServerFactory
	KQLQueryset                    kqlquerysetfake.ServerFactory
	Lakehouse                      lakehousefake.ServerFactory
	MirroredAzureDatabricksCatalog mirroredazuredatabrickscatalogfake.ServerFactory
	MirroredDatabase               mirroreddatabasefake.ServerFactory
	MirroredWarehouse              mirroredwarehousefake.ServerFactory
	MLExperiment                   mlexperimentfake.ServerFactory
	MLModel                        mlmodelfake.ServerFactory
	MountedDataFactory             mounteddatafactoryfake.ServerFactory
	Notebook                       notebookfake.ServerFactory
	PaginatedReport                paginatedreportfake.ServerFactory
	Reflex                         reflexfake.ServerFactory
	Report                         reportfake.ServerFactory
	SemanticModel                  semanticmodelfake.ServerFactory
	Spark                          sparkfake.ServerFactory
	SparkJobDefinition             sparkjobdefinitionfake.ServerFactory
	SQLDatabase                    sqldatabasefake.ServerFactory
	SQLEndpoint                    sqlendpointfake.ServerFactory
	VariableLibrary                variablelibraryfake.ServerFactory
	Warehouse                      warehousefake.ServerFactory
	WarehouseSnapshot              warehousesnapshotfake.ServerFactory
}

// ServerFactoryTransport connects instance of fabric.Client to instance of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                              *ServerFactory
	trMu                             sync.Mutex
	trAdmin                          *adminfake.ServerFactoryTransport
	trApache                         *apacheairflowjobfake.ServerFactoryTransport
	trCopyJob                        *copyjobfake.ServerFactoryTransport
	trCore                           *corefake.ServerFactoryTransport
	trDashboard                      *dashboardfake.ServerFactoryTransport
	trDataflow                       *dataflowfake.ServerFactoryTransport
	trDatamart                       *datamartfake.ServerFactoryTransport
	trDataPipeline                   *datapipelinefake.ServerFactoryTransport
	trDigitalTwinBuilder             *digitaltwinbuilderfake.ServerFactoryTransport
	trDigitalTwinBuilderFlow         *digitaltwinbuilderflowfake.ServerFactoryTransport
	trEnvironment                    *environmentfake.ServerFactoryTransport
	trEventhouse                     *eventhousefake.ServerFactoryTransport
	trEventstream                    *eventstreamfake.ServerFactoryTransport
	trGraphQLApi                     *graphqlapifake.ServerFactoryTransport
	trKQLDashboard                   *kqldashboardfake.ServerFactoryTransport
	trKQLDatabase                    *kqldatabasefake.ServerFactoryTransport
	trKQLQueryset                    *kqlquerysetfake.ServerFactoryTransport
	trLakehouse                      *lakehousefake.ServerFactoryTransport
	trMirroredAzureDatabricksCatalog *mirroredazuredatabrickscatalogfake.ServerFactoryTransport
	trMirroredDatabase               *mirroreddatabasefake.ServerFactoryTransport
	trMirroredWarehouse              *mirroredwarehousefake.ServerFactoryTransport
	trMLExperiment                   *mlexperimentfake.ServerFactoryTransport
	trMLModel                        *mlmodelfake.ServerFactoryTransport
	trMountedDataFactory             *mounteddatafactoryfake.ServerFactoryTransport
	trNotebook                       *notebookfake.ServerFactoryTransport
	trPaginatedReport                *paginatedreportfake.ServerFactoryTransport
	trReflex                         *reflexfake.ServerFactoryTransport
	trReport                         *reportfake.ServerFactoryTransport
	trSemanticModel                  *semanticmodelfake.ServerFactoryTransport
	trSpark                          *sparkfake.ServerFactoryTransport
	trSparkJobDefinition             *sparkjobdefinitionfake.ServerFactoryTransport
	trSQLDatabase                    *sqldatabasefake.ServerFactoryTransport
	trSQLEndpoint                    *sqlendpointfake.ServerFactoryTransport
	trVariableLibrary                *variablelibraryfake.ServerFactoryTransport
	trWarehouse                      *warehousefake.ServerFactoryTransport
	trWarehouseSnapshot              *warehousesnapshotfake.ServerFactoryTransport
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of fabric.Client via the
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
	client := strings.Split(method, ".")[0]
	var resp *http.Response
	var err error

	switch client {
	case "admin":
		initServer(s, &s.trAdmin, func() *adminfake.ServerFactoryTransport { return adminfake.NewServerFactoryTransport(&s.srv.Admin) })
		resp, err = s.trAdmin.Do(req)
	case "apacheairflowjob":
		initServer(s, &s.trApache, func() *apacheairflowjobfake.ServerFactoryTransport {
			return apacheairflowjobfake.NewServerFactoryTransport(&s.srv.Apache)
		})
		resp, err = s.trApache.Do(req)
	case "copyjob":
		initServer(s, &s.trCopyJob, func() *copyjobfake.ServerFactoryTransport {
			return copyjobfake.NewServerFactoryTransport(&s.srv.CopyJob)
		})
		resp, err = s.trCopyJob.Do(req)
	case "core":
		initServer(s, &s.trCore, func() *corefake.ServerFactoryTransport { return corefake.NewServerFactoryTransport(&s.srv.Core) })
		resp, err = s.trCore.Do(req)
	case "dashboard":
		initServer(s, &s.trDashboard, func() *dashboardfake.ServerFactoryTransport {
			return dashboardfake.NewServerFactoryTransport(&s.srv.Dashboard)
		})
		resp, err = s.trDashboard.Do(req)
	case "dataflow":
		initServer(s, &s.trDataflow, func() *dataflowfake.ServerFactoryTransport {
			return dataflowfake.NewServerFactoryTransport(&s.srv.Dataflow)
		})
		resp, err = s.trDataflow.Do(req)
	case "datamart":
		initServer(s, &s.trDatamart, func() *datamartfake.ServerFactoryTransport {
			return datamartfake.NewServerFactoryTransport(&s.srv.Datamart)
		})
		resp, err = s.trDatamart.Do(req)
	case "datapipeline":
		initServer(s, &s.trDataPipeline, func() *datapipelinefake.ServerFactoryTransport {
			return datapipelinefake.NewServerFactoryTransport(&s.srv.DataPipeline)
		})
		resp, err = s.trDataPipeline.Do(req)
	case "digitaltwinbuilder":
		initServer(s, &s.trDigitalTwinBuilder, func() *digitaltwinbuilderfake.ServerFactoryTransport {
			return digitaltwinbuilderfake.NewServerFactoryTransport(&s.srv.DigitalTwinBuilder)
		})
		resp, err = s.trDigitalTwinBuilder.Do(req)
	case "digitaltwinbuilderflow":
		initServer(s, &s.trDigitalTwinBuilderFlow, func() *digitaltwinbuilderflowfake.ServerFactoryTransport {
			return digitaltwinbuilderflowfake.NewServerFactoryTransport(&s.srv.DigitalTwinBuilderFlow)
		})
		resp, err = s.trDigitalTwinBuilderFlow.Do(req)
	case "environment":
		initServer(s, &s.trEnvironment, func() *environmentfake.ServerFactoryTransport {
			return environmentfake.NewServerFactoryTransport(&s.srv.Environment)
		})
		resp, err = s.trEnvironment.Do(req)
	case "eventhouse":
		initServer(s, &s.trEventhouse, func() *eventhousefake.ServerFactoryTransport {
			return eventhousefake.NewServerFactoryTransport(&s.srv.Eventhouse)
		})
		resp, err = s.trEventhouse.Do(req)
	case "eventstream":
		initServer(s, &s.trEventstream, func() *eventstreamfake.ServerFactoryTransport {
			return eventstreamfake.NewServerFactoryTransport(&s.srv.Eventstream)
		})
		resp, err = s.trEventstream.Do(req)
	case "graphqlapi":
		initServer(s, &s.trGraphQLApi, func() *graphqlapifake.ServerFactoryTransport {
			return graphqlapifake.NewServerFactoryTransport(&s.srv.GraphQLApi)
		})
		resp, err = s.trGraphQLApi.Do(req)
	case "kqldashboard":
		initServer(s, &s.trKQLDashboard, func() *kqldashboardfake.ServerFactoryTransport {
			return kqldashboardfake.NewServerFactoryTransport(&s.srv.KQLDashboard)
		})
		resp, err = s.trKQLDashboard.Do(req)
	case "kqldatabase":
		initServer(s, &s.trKQLDatabase, func() *kqldatabasefake.ServerFactoryTransport {
			return kqldatabasefake.NewServerFactoryTransport(&s.srv.KQLDatabase)
		})
		resp, err = s.trKQLDatabase.Do(req)
	case "kqlqueryset":
		initServer(s, &s.trKQLQueryset, func() *kqlquerysetfake.ServerFactoryTransport {
			return kqlquerysetfake.NewServerFactoryTransport(&s.srv.KQLQueryset)
		})
		resp, err = s.trKQLQueryset.Do(req)
	case "lakehouse":
		initServer(s, &s.trLakehouse, func() *lakehousefake.ServerFactoryTransport {
			return lakehousefake.NewServerFactoryTransport(&s.srv.Lakehouse)
		})
		resp, err = s.trLakehouse.Do(req)
	case "mirroredazuredatabrickscatalog":
		initServer(s, &s.trMirroredAzureDatabricksCatalog, func() *mirroredazuredatabrickscatalogfake.ServerFactoryTransport {
			return mirroredazuredatabrickscatalogfake.NewServerFactoryTransport(&s.srv.MirroredAzureDatabricksCatalog)
		})
		resp, err = s.trMirroredAzureDatabricksCatalog.Do(req)
	case "mirroreddatabase":
		initServer(s, &s.trMirroredDatabase, func() *mirroreddatabasefake.ServerFactoryTransport {
			return mirroreddatabasefake.NewServerFactoryTransport(&s.srv.MirroredDatabase)
		})
		resp, err = s.trMirroredDatabase.Do(req)
	case "mirroredwarehouse":
		initServer(s, &s.trMirroredWarehouse, func() *mirroredwarehousefake.ServerFactoryTransport {
			return mirroredwarehousefake.NewServerFactoryTransport(&s.srv.MirroredWarehouse)
		})
		resp, err = s.trMirroredWarehouse.Do(req)
	case "mlexperiment":
		initServer(s, &s.trMLExperiment, func() *mlexperimentfake.ServerFactoryTransport {
			return mlexperimentfake.NewServerFactoryTransport(&s.srv.MLExperiment)
		})
		resp, err = s.trMLExperiment.Do(req)
	case "mlmodel":
		initServer(s, &s.trMLModel, func() *mlmodelfake.ServerFactoryTransport {
			return mlmodelfake.NewServerFactoryTransport(&s.srv.MLModel)
		})
		resp, err = s.trMLModel.Do(req)
	case "mounteddatafactory":
		initServer(s, &s.trMountedDataFactory, func() *mounteddatafactoryfake.ServerFactoryTransport {
			return mounteddatafactoryfake.NewServerFactoryTransport(&s.srv.MountedDataFactory)
		})
		resp, err = s.trMountedDataFactory.Do(req)
	case "notebook":
		initServer(s, &s.trNotebook, func() *notebookfake.ServerFactoryTransport {
			return notebookfake.NewServerFactoryTransport(&s.srv.Notebook)
		})
		resp, err = s.trNotebook.Do(req)
	case "paginatedreport":
		initServer(s, &s.trPaginatedReport, func() *paginatedreportfake.ServerFactoryTransport {
			return paginatedreportfake.NewServerFactoryTransport(&s.srv.PaginatedReport)
		})
		resp, err = s.trPaginatedReport.Do(req)
	case "reflex":
		initServer(s, &s.trReflex, func() *reflexfake.ServerFactoryTransport { return reflexfake.NewServerFactoryTransport(&s.srv.Reflex) })
		resp, err = s.trReflex.Do(req)
	case "report":
		initServer(s, &s.trReport, func() *reportfake.ServerFactoryTransport { return reportfake.NewServerFactoryTransport(&s.srv.Report) })
		resp, err = s.trReport.Do(req)
	case "semanticmodel":
		initServer(s, &s.trSemanticModel, func() *semanticmodelfake.ServerFactoryTransport {
			return semanticmodelfake.NewServerFactoryTransport(&s.srv.SemanticModel)
		})
		resp, err = s.trSemanticModel.Do(req)
	case "spark":
		initServer(s, &s.trSpark, func() *sparkfake.ServerFactoryTransport { return sparkfake.NewServerFactoryTransport(&s.srv.Spark) })
		resp, err = s.trSpark.Do(req)
	case "sparkjobdefinition":
		initServer(s, &s.trSparkJobDefinition, func() *sparkjobdefinitionfake.ServerFactoryTransport {
			return sparkjobdefinitionfake.NewServerFactoryTransport(&s.srv.SparkJobDefinition)
		})
		resp, err = s.trSparkJobDefinition.Do(req)
	case "sqldatabase":
		initServer(s, &s.trSQLDatabase, func() *sqldatabasefake.ServerFactoryTransport {
			return sqldatabasefake.NewServerFactoryTransport(&s.srv.SQLDatabase)
		})
		resp, err = s.trSQLDatabase.Do(req)
	case "sqlendpoint":
		initServer(s, &s.trSQLEndpoint, func() *sqlendpointfake.ServerFactoryTransport {
			return sqlendpointfake.NewServerFactoryTransport(&s.srv.SQLEndpoint)
		})
		resp, err = s.trSQLEndpoint.Do(req)
	case "variablelibrary":
		initServer(s, &s.trVariableLibrary, func() *variablelibraryfake.ServerFactoryTransport {
			return variablelibraryfake.NewServerFactoryTransport(&s.srv.VariableLibrary)
		})
		resp, err = s.trVariableLibrary.Do(req)
	case "warehouse":
		initServer(s, &s.trWarehouse, func() *warehousefake.ServerFactoryTransport {
			return warehousefake.NewServerFactoryTransport(&s.srv.Warehouse)
		})
		resp, err = s.trWarehouse.Do(req)
	case "warehousesnapshot":
		initServer(s, &s.trWarehouseSnapshot, func() *warehousesnapshotfake.ServerFactoryTransport {
			return warehousesnapshotfake.NewServerFactoryTransport(&s.srv.WarehouseSnapshot)
		})
		resp, err = s.trWarehouseSnapshot.Do(req)
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
