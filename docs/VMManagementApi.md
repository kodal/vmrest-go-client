# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigVMParams**](VMManagementApi.md#ConfigVMParams) | **Put** /vms/{id}/params | update the vm config params
[**CreateVM**](VMManagementApi.md#CreateVM) | **Post** /vms | Creates a copy of the VM
[**DeleteVM**](VMManagementApi.md#DeleteVM) | **Delete** /vms/{id} | Deletes a VM
[**GetAllVMs**](VMManagementApi.md#GetAllVMs) | **Get** /vms | Returns a list of VM IDs and paths for all VMs
[**GetVM**](VMManagementApi.md#GetVM) | **Get** /vms/{id} | Returns the VM setting information of a VM
[**GetVMParams**](VMManagementApi.md#GetVMParams) | **Get** /vms/{id}/params/{name} | Get the VM config params
[**GetVMRestrictions**](VMManagementApi.md#GetVMRestrictions) | **Get** /vms/{id}/restrictions | Returns the restrictions information of the VM
[**RegisterVM**](VMManagementApi.md#RegisterVM) | **Post** /vms/registration | Register VM to VM Library
[**UpdateVM**](VMManagementApi.md#UpdateVM) | **Put** /vms/{id} | Updates the VM settings

# **ConfigVMParams**
> ErrorModel ConfigVMParams(ctx, body, id, optional)
update the vm config params

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigVmParamsParameter**](ConfigVmParamsParameter.md)| Parameters set to the VM | 
  **id** | **string**| ID of VM | 
 **optional** | ***VMManagementApiConfigVMParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiConfigVMParamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVM**
> VmInformation CreateVM(ctx, body, optional)
Creates a copy of the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VmCloneParameter**](VmCloneParameter.md)| Parameters of VM to create | 
 **optional** | ***VMManagementApiCreateVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiCreateVMOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**VmInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVM**
> DeleteVM(ctx, id, optional)
Deletes a VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
 **optional** | ***VMManagementApiDeleteVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiDeleteVMOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVMs**
> []Vmid GetAllVMs(ctx, )
Returns a list of VM IDs and paths for all VMs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Vmid**](VMID.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVM**
> VmInformation GetVM(ctx, id, optional)
Returns the VM setting information of a VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
 **optional** | ***VMManagementApiGetVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiGetVMOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**VmInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMParams**
> ConfigVmParamsParameter GetVMParams(ctx, id, name, optional)
Get the VM config params

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **name** | **string**| Name of the param | 
 **optional** | ***VMManagementApiGetVMParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiGetVMParamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**ConfigVmParamsParameter**](ConfigVMParamsParameter.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMRestrictions**
> VmRestrictionsInformation GetVMRestrictions(ctx, id, optional)
Returns the restrictions information of the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
 **optional** | ***VMManagementApiGetVMRestrictionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiGetVMRestrictionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**VmRestrictionsInformation**](VMRestrictionsInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterVM**
> VmRrgistrationInformation RegisterVM(ctx, body)
Register VM to VM Library

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VmRegisterParameter**](VmRegisterParameter.md)| Parameters of the VM to register | 

### Return type

[**VmRrgistrationInformation**](VMRrgistrationInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVM**
> VmInformation UpdateVM(ctx, body, id, optional)
Updates the VM settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VmParameter**](VmParameter.md)| VM definition | 
  **id** | **string**| ID of VM | 
 **optional** | ***VMManagementApiUpdateVMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMManagementApiUpdateVMOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**VmInformation**](VMInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

