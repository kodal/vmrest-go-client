# \VMManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigVMParams**](VMManagementAPI.md#ConfigVMParams) | **Put** /vms/{id}/params | update the vm config params
[**CreateVM**](VMManagementAPI.md#CreateVM) | **Post** /vms | Creates a copy of the VM
[**DeleteVM**](VMManagementAPI.md#DeleteVM) | **Delete** /vms/{id} | Deletes a VM
[**GetAllVMs**](VMManagementAPI.md#GetAllVMs) | **Get** /vms | Returns a list of VM IDs and paths for all VMs
[**GetVM**](VMManagementAPI.md#GetVM) | **Get** /vms/{id} | Returns the VM setting information of a VM
[**GetVMParams**](VMManagementAPI.md#GetVMParams) | **Get** /vms/{id}/params/{name} | Get the VM config params
[**GetVMRestrictions**](VMManagementAPI.md#GetVMRestrictions) | **Get** /vms/{id}/restrictions | Returns the restrictions information of the VM
[**RegisterVM**](VMManagementAPI.md#RegisterVM) | **Post** /vms/registration | Register VM to VM Library
[**UpdateVM**](VMManagementAPI.md#UpdateVM) | **Put** /vms/{id} | Updates the VM settings



## ConfigVMParams

> ErrorModel ConfigVMParams(ctx, id).ConfigVMParamsParameter(configVMParamsParameter).VmPassword(vmPassword).Execute()

update the vm config params

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	configVMParamsParameter := *vmrest.NewConfigVMParamsParameter() // ConfigVMParamsParameter | Parameters set to the VM
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.ConfigVMParams(context.Background(), id).ConfigVMParamsParameter(configVMParamsParameter).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.ConfigVMParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigVMParams`: ErrorModel
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.ConfigVMParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigVMParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configVMParamsParameter** | [**ConfigVMParamsParameter**](ConfigVMParamsParameter.md) | Parameters set to the VM | 
 **vmPassword** | **string** | VM password | 

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVM

> VMInformation CreateVM(ctx).VMCloneParameter(vMCloneParameter).VmPassword(vmPassword).Execute()

Creates a copy of the VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	vMCloneParameter := *vmrest.NewVMCloneParameter("Name_example", "ParentId_example") // VMCloneParameter | Parameters of VM to create
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.CreateVM(context.Background()).VMCloneParameter(vMCloneParameter).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.CreateVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVM`: VMInformation
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.CreateVM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMCloneParameter** | [**VMCloneParameter**](VMCloneParameter.md) | Parameters of VM to create | 
 **vmPassword** | **string** | VM password | 

### Return type

[**VMInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVM

> DeleteVM(ctx, id).VmPassword(vmPassword).Execute()

Deletes a VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	r, err := apiClient.VMManagementAPI.DeleteVM(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.DeleteVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVMs

> []VMID GetAllVMs(ctx).Execute()

Returns a list of VM IDs and paths for all VMs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.GetAllVMs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.GetAllVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllVMs`: []VMID
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.GetAllVMs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVMsRequest struct via the builder pattern


### Return type

[**[]VMID**](VMID.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVM

> VMInformation GetVM(ctx, id).VmPassword(vmPassword).Execute()

Returns the VM setting information of a VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.GetVM(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.GetVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVM`: VMInformation
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.GetVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password | 

### Return type

[**VMInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMParams

> ConfigVMParamsParameter GetVMParams(ctx, id, name).VmPassword(vmPassword).Execute()

Get the VM config params

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	name := "name_example" // string | Name of the param
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.GetVMParams(context.Background(), id, name).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.GetVMParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMParams`: ConfigVMParamsParameter
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.GetVMParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 
**name** | **string** | Name of the param | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | **string** | VM password | 

### Return type

[**ConfigVMParamsParameter**](ConfigVMParamsParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMRestrictions

> VMRestrictionsInformation GetVMRestrictions(ctx, id).VmPassword(vmPassword).Execute()

Returns the restrictions information of the VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.GetVMRestrictions(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.GetVMRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMRestrictions`: VMRestrictionsInformation
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.GetVMRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password | 

### Return type

[**VMRestrictionsInformation**](VMRestrictionsInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVM

> VMRrgistrationInformation RegisterVM(ctx).VMRegisterParameter(vMRegisterParameter).Execute()

Register VM to VM Library

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	vMRegisterParameter := *vmrest.NewVMRegisterParameter() // VMRegisterParameter | Parameters of the VM to register

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.RegisterVM(context.Background()).VMRegisterParameter(vMRegisterParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.RegisterVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterVM`: VMRrgistrationInformation
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.RegisterVM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMRegisterParameter** | [**VMRegisterParameter**](VMRegisterParameter.md) | Parameters of the VM to register | 

### Return type

[**VMRrgistrationInformation**](VMRrgistrationInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVM

> VMInformation UpdateVM(ctx, id).VMParameter(vMParameter).VmPassword(vmPassword).Execute()

Updates the VM settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	vmrest "github.com/kodal/vmrest-go-client"
)

func main() {
	id := "id_example" // string | ID of VM
	vMParameter := *vmrest.NewVMParameter() // VMParameter | VM definition
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMManagementAPI.UpdateVM(context.Background(), id).VMParameter(vMParameter).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMManagementAPI.UpdateVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVM`: VMInformation
	fmt.Fprintf(os.Stdout, "Response from `VMManagementAPI.UpdateVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vMParameter** | [**VMParameter**](VMParameter.md) | VM definition | 
 **vmPassword** | **string** | VM password | 

### Return type

[**VMInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

