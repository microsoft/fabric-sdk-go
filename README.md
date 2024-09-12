# Microsoft Fabric SDK for Go

[![GoDoc](https://godoc.org/github.com/microsoft/fabric-sdk-go/?status.svg)](https://pkg.go.dev/github.com/microsoft/fabric-sdk-go)

This project provides a data-plane Go SDK for [Microsoft Fabric](https://www.microsoft.com/microsoft-fabric). Microsoft Fabric is an all-in-one analytics solution for enterprises that covers everything from data movement to data science, Real-Time Analytics, and business intelligence. It offers a comprehensive suite of services, including data lake, data engineering, and data integration, all in one place.

> [!WARNING]
> This code is experimental and provided solely for evaluation purposes. It is **NOT** intended for production use and may contain bugs, incomplete features, or other issues. Use at your own risk, as it may undergo significant changes without notice, and no guarantees or support are provided. By using this code, you acknowledge and agree to these conditions. Consult the documentation or contact the maintainer if you have questions or concerns.

## Getting started

### Prerequisites

* Go, version `1.22` or higher - [Install Go](https://go.dev/doc/install)
* [Microsoft Fabric](https://learn.microsoft.com/fabric/get-started/microsoft-fabric-overview)

### Install modules

Install the `fabric` and `azidentity` modules for Go with `go get`:

```shell
go get -u github.com/microsoft/fabric-sdk-go
go get -u github.com/Azure/azure-sdk-for-go/sdk/azidentity
```

The [azidentity](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity) module is used for [Microsoft Entra ID](https://www.microsoft.com/security/business/microsoft-entra) authentication with Microsoft Fabric.

## Examples

Examples for various scenarios can be found in the `*_example_test.go` files under each package.

## Authentication

To work with Fabric SDK, like many other Microsoft services, first you need to authenticate and get authorization to Fabric services using the Microsoft Entra ID.

The `azidentity` module provides Microsoft Entra ID (formerly [Azure Active Directory](https://learn.microsoft.com/entra/fundamentals/new-name)) token authentication support across the Microsoft SDKs. It includes a set of `TokenCredential` implementations, which can be used with Microsoft SDKs clients supporting token authentication, include Fabric SDK.

This example demonstrates authenticating with `DefaultAzureCredential`.

```go
import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
```

```go
cred, err := azidentity.NewDefaultAzureCredential(nil)
if err != nil {
  // handle error
}
```

Explore all [Credential Types](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#readme-credential-types) from `azidentity` module to choose the ones that best suit your needs.

> [!NOTE]
> `DefaultAzureCredential` is intended to simplify getting started with the SDK by handling common scenarios with reasonable default behaviors. Developers who want more control or whose scenario isn't served by the default settings should use other credential types. Read more in the [Key concept](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#readme-key-concepts) section.

### Microsoft Entra ID app

A Microsoft Entra ID app controls access levels for your Fabric services. Before you can make any calls, you'll have to register a Microsoft Entra ID app. The app allows you to:

* Establish an identity for your app
* Let your app access the Microsoft Fabric
* Specify your app's Fabric permissions

Please follow [Create a Microsoft Entra ID app](https://learn.microsoft.com/rest/api/fabric/articles/get-started/create-entra-app) guide how to register your app and setup delegated [scopes](https://learn.microsoft.com/rest/api/fabric/articles/scopes).

#### Azure CLI usage

To use [Azure CLI](https://learn.microsoft.com/cli/azure/) cached credentials that will be consumed by `azidentity`, additionally to basic configuration described above, you have to Expose API and pre-authorize Azure CLI to access your Microsoft Entra ID app.

In the `Expose an API` menu of your app, you need to define your application ID URI:

* Application ID URI: `api://<client_id>` - you can use default client id, or provide own name, for example:

```text
api://my_fabric_app
```

* Add required scope in the `Scopes defined by this API` section:

1. Scope name: `default`
1. Who can consent: `Admins and users`
1. Admin consent display name: `My Fabric App`
1. Admin consent description: `Allows connection to backend services for My Fabric App`
1. User consent display name: `My Fabric App`
1. User consent description: `Allows connection to backend services for My Fabric App`
1. State: `Enabled`

* You will finally need to pre-authorize Azure CLI to access your API by adding Azure CLI's client application `04b07795-8ddb-461a-bbee-02f9e1bf7b46` in the `Authorized client applications` section.

After above steps you should be able to authenticate using Azure CLI:

```shell
az login --allow-no-subscriptions --scope api://my_fabric_app/default
```

## Client usage

The Fabric Go SDK utilizes a client factory pattern that allows users to interact with various Fabric services via their respective clients.

### Fabric Client

The top level client is Fabric client, which can be used to created sub clients via the Factor pattern.

```go
import "github.com/microsoft/fabric-sdk-go/fabric"
```

```go
fabClient, err := fabric.NewClient(cred, nil, nil)
if err != nil {
  // handle error
}
```

### Client Factory with Fabric Client

The Fabric client can be used to create a client for each service. This can be used to get a specific item client in the corresponding service. In the example below, we are creating a core service client that will allow us to create clients for specific items under the core service.

```go
import "github.com/microsoft/fabric-sdk-go/fabric/core"
```

```go
coreCF := core.NewClientFactoryWithClient(*fabClient)
```

### Client Factory

Another way to create a service client is directly using the credentials from `azidentity`

```go
import "github.com/microsoft/fabric-sdk-go/fabric/core"
```

```go
coreCF, err := core.NewClientFactory(cred, nil, nil)
if err != nil {
  // handle error
}
```

### Item-level Client

An item level client can be created with the factory client for the corresponding service. For example, using the core service factory, a new workspaces client is created blow.

```go
workspacesClient := coreCF.NewWorkspacesClient()
```

## Long Running Operations (LRO)

Fabric uses the LRO pattern in many request to handle operations that take an amount of time to complete. These operations are typically non-blocking and allow the client to continue with other tasks without waiting for the operation to complete.

To learn more about how Fabric API behaves with LRO, read [Fabric API - Long Running Operations](https://learn.microsoft.com/rest/api/fabric/articles/long-running-operation) article.

The SDK expose two ways of handling results from the LROs requests:

* with sync method - where polling is managed with return of the final result
* with Poller - where you can control behavior of polling and getting results

### LRO with sync method

The SDK provides utility Sync methods that handle waiting for a longer running operation on the behalf of a user. This will result in polling being handled via internal components and the method blocks until the operation is complete.

In the example below a lakehouse is created using the CreateLakeHouse method.

```go
lakehouseCF := lakehouse.NewClientFactoryWithClient(*fabricClient)
lakehouseItemsClient := lakehouseCF.NewItemsClient()
req := lakehouse.CreateLakehouseRequest{
 DisplayName: to.Ptr(lakehouseName),
 Description: to.Ptr("Created by Go SDK"),
}
resp, err := lakehouseItemsClient.CreateLakehouse(ctx, *myWs.Workspace.ID, req, nil)

```

### LRO using Poller

For situations that require more granular control. The SDK exposes native poller objects which can be used to poll for the completion of a long running operations.

Each long running operation includes a `Begin*` method to start an LRO. It will return a poller that can used to keep polling for the result until LRO is done.

For example, to create a lakehouse using the native poller constructs:

```go
lakehouseCF := lakehouse.NewClientFactoryWithClient(*fabricClient)
lakehouseItemsClient := lakehouseCF.NewItemsClient()
req := lakehouse.CreateLakehouseRequest{
  DisplayName: to.Ptr("lh_mylakehouse"),
}
poller, err := lakehouseItemsClient.BeginCreateLakehouse(ctx, "workspace_id", req, nil)
if err != nil {
  // handle error...
}
resp, err = poller.PollUntilDone(ctx, nil)
if err != nil {
  // handle error...
}
// dealing with `resp`
```

#### Concurrency for long running operations via go channels

Since SDK provides sync methods to handle long running operations, go channels can be utilized to support concurrent programming patterns.

```go
ch := make(chan ItemsClientCreateLakehouseResponse)

go func(){
  lakehouseCF := lakehouse.NewClientFactoryWithClient(*fabricClient)
  lakehouseItemsClient := lakehouseCF.NewItemsClient()
  req := lakehouse.CreateLakehouseRequest{
    DisplayName: to.Ptr("lh_mylakehouse"),
    Description: to.Ptr("Created by Go SDK"),
  }()

  resp, err := lakehouseItemsClient.CreateLakehouse(ctx, *myWs.Workspace.ID, req, nil)
  if err!= nil{
  //handle error
  }
  ch <- resp
}

resp := <- ch
```

## Pagination

Pagination refers to the practice of breaking up a large set of data into smaller, manageable chunks or pages when delivering the data to a client application. It's a common technique used to improve the performance and efficiency of API requests, especially when dealing with a large amount of data. Pagination is also used to prevent data loss when there's too much data to display in one chunk.

To learn more about how Fabric API behaves with Pagination, read [Fabric API - Pagination](https://learn.microsoft.com/rest/api/fabric/articles/pagination) article.

The SDK expose two ways of handling results from the Pageable requests:

* with result method - where iterating is managed with return of the final result
* with Pager - where you can control behavior of iterating and getting results

### Pagination with result method

Each pagable operation in the Go SDK exposes an result (or helper) method that will handle iterating through all pages of data and returning back results.

Note: This is a convenience utility for scenarios that do not have a lot of data underneath. For large sets of data, it's recommended to use the native pager object to avoid fetching all the data.

```go
// create a workspace client in the core service
workspaceClient := coreCF.NewWorkspacesClient()

// call the ListWorkspaces method to fetch all workspaces
workspaces, err := workspaceClient.ListWorkspaces(ctx, nil)
```

### Pagination using Pager

For situations that require more granular control. The SDK exposes native pager objects which can be used to poll for the completion of a long running operations.

The `New*Pager` methods are used to create a pager helper for all pageable operations. With the returned `*runtime.Pager[T]`, you can fetch pages and determine if there are more pages to fetch.

For example, the following code lists workspaces by using the native pager object.

```go
// create a workspace client in the core service
workspacesClient := coreCF.NewWorkspacesClient()

// create a new Workspaces pager object
pager := workspacesClient.NewListWorkspacesPager(nil)
var workspaces []*core.Workspace

// while the pager has more items (pager.More() == true), fetch the next page.
for pager.More() {
  page, err := pager.NextPage(ctx)
  if err != nil {
  }
  if page.Value != nil {
    workspaces = append(workspaces, page.Value...)
  }
}
```

More info about [AzCore pageable operations](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore#hdr-Pageable_Operations)

## Testing

Each package contains sub-package called `fake` that provides implementations for fake servers that can be used for creating unit tests.

> [!NOTE]
> While the examples use `lakehouse`, the patterns are applicable to any package with a fake sub-package.

To create a fake server, declare an instance of fake server factory with required fake server type(s).

```go
import "github.com/microsoft/fabric-sdk-go/fabric/lakehouse/fake"
```

```go
lakehouseFakeSF := fake.ServerFactory{
  ItemsServer: fake.ItemsServer{},
}
```

Next, provide func implementations for the client methods you wish to fake. The named return variables can be used to simplify return value construction.

```go
import azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
```

```go
lakehouseFakeSF.ItemsServer.GetLakehouse(ctx context.Context, workspaceID string, lakehouseID string, options *lakehouse.ItemsClientGetLakehouseOptions) (resp azfake.Responder[lakehouse.ItemsClientGetLakehouseResponse], errResp azfake.ErrorResponder) {
  // resp.SetResponse(/* your fake lakehouse.ItemsClientGetLakehouseResponse response */)
  return
}
```

You connect the fake server factory to a client factory instance during its construction through the optional transport.

Use the `TokenCredential` type from `azcore/fake` to create a fake `azcore.TokenCredential`.

```go
import azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
```

```go
lakehouseCF, err := lakehouse.NewClientFactory(&azfake.TokenCredential{}, nil, &azcore.ClientOptions{
  Transport: fake.NewServerFactoryTransport(&lakehouseFakeSF),
})
if err != nil {
  // handle error
}
```

Calling methods on the client will pass the provided values to the matching fake implementation. The values can be arbitrary, including the zero-value for any/all parameters.

```go
itemsClient := lakehouseCF.NewItemsClient()
resp, err := itemsClient.GetLakehouse(context.TODO(), "fake-workspace-id", "fake-lakehouse-id", nil)
if err != nil {
  // handle error
}
```

More info about [AzCore fakes](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore#hdr-Fakes)

## Tips & Tricks

`azcore` provides several important items and utilities. One such utility is the `to.Ptr()` method which can take a golang value and return a point to the value. This also works with literals.

```go
import "github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
```

```go
to.Ptr()
```

## Troubleshooting

### Error Handling

All methods that send HTTP requests return `error` what you can unwrap. For teh Fabric related errors use `*core.ResponseError`. It has error details and the raw response from the service.

```go
// Handle Fabric error response
var errRespFabric *core.ResponseError
errors.As(err, &errRespFabric)

// Handle AzIdentity error response
var errRespAzIdentity *azidentity.AuthenticationFailedError
errors.As(err, &errRespAzIdentity)

// Handle AzCore error response
var errRespAzCore *azcore.ResponseError
errors.As(err, &errRespAzCore)
```

## Contributing

This project welcomes suggestions via GitHub Issues. Please do **NOT** submit PRs - the code in this repo is auto generated.

Most contributions require you to agree to a Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us the rights to use your contribution. For details, visit <https://cla.opensource.microsoft.com>.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft trademarks or logos is subject to and must follow[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general). Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship. Any use of third-party trademarks or logos are subject to those third-party's policies.
