# \VMPowerManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePowerState**](VMPowerManagementAPI.md#ChangePowerState) | **Put** /vms/{id}/power | Changes the VM power state
[**GetPowerState**](VMPowerManagementAPI.md#GetPowerState) | **Get** /vms/{id}/power | Returns the power state of the VM



## ChangePowerState

> VMPowerState ChangePowerState(ctx, id).Body(body).VmPassword(vmPassword).Execute()

Changes the VM power state

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
	body := string(987) // string | VM power operation: on, off, shutdown, suspend, pause, unpause
	vmPassword := "vmPassword_example" // string | VM password for encrypted VM. If VM is already powered on password is not needed. (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPowerManagementAPI.ChangePowerState(context.Background(), id).Body(body).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPowerManagementAPI.ChangePowerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePowerState`: VMPowerState
	fmt.Fprintf(os.Stdout, "Response from `VMPowerManagementAPI.ChangePowerState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePowerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | VM power operation: on, off, shutdown, suspend, pause, unpause | 
 **vmPassword** | **string** | VM password for encrypted VM. If VM is already powered on password is not needed. | 

### Return type

[**VMPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerState

> VMPowerState GetPowerState(ctx, id).VmPassword(vmPassword).Execute()

Returns the power state of the VM

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
	vmPassword := "vmPassword_example" // string | VM password for encrypted VM. If VM is already powered on password is not needed. (optional)

	configuration := vmrest.NewConfiguration()
	apiClient := vmrest.NewAPIClient(configuration)
	resp, r, err := apiClient.VMPowerManagementAPI.GetPowerState(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMPowerManagementAPI.GetPowerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPowerState`: VMPowerState
	fmt.Fprintf(os.Stdout, "Response from `VMPowerManagementAPI.GetPowerState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPowerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password for encrypted VM. If VM is already powered on password is not needed. | 

### Return type

[**VMPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

