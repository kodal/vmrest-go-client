# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePowerState**](VMPowerManagementApi.md#ChangePowerState) | **Put** /vms/{id}/power | Changes the VM power state
[**GetPowerState**](VMPowerManagementApi.md#GetPowerState) | **Get** /vms/{id}/power | Returns the power state of the VM

# **ChangePowerState**
> VmPowerState ChangePowerState(ctx, body, id, optional)
Changes the VM power state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| VM power operation: on, off, shutdown, suspend, pause, unpause | 
  **id** | **string**| ID of VM | 
 **optional** | ***VMPowerManagementApiChangePowerStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMPowerManagementApiChangePowerStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | [**optional.Interface of string**](.md)| VM password for encrypted VM. If VM is already powered on password is not needed. | 

### Return type

[**VmPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPowerState**
> VmPowerState GetPowerState(ctx, id, optional)
Returns the power state of the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
 **optional** | ***VMPowerManagementApiGetPowerStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMPowerManagementApiGetPowerStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password for encrypted VM. If VM is already powered on password is not needed. | 

### Return type

[**VmPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

