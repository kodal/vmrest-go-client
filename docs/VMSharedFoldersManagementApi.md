# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedFolder**](VMSharedFoldersManagementApi.md#CreateSharedFolder) | **Post** /vms/{id}/sharedfolders | Mounts a new shared folder in the VM
[**DeleteSharedFolder**](VMSharedFoldersManagementApi.md#DeleteSharedFolder) | **Delete** /vms/{id}/sharedfolders/{folder id} | Deletes a shared folder
[**GetAllSharedFolders**](VMSharedFoldersManagementApi.md#GetAllSharedFolders) | **Get** /vms/{id}/sharedfolders | Returns all shared folders mounted in the VM
[**UpdataSharedFolder**](VMSharedFoldersManagementApi.md#UpdataSharedFolder) | **Put** /vms/{id}/sharedfolders/{folder id} | Updates a shared folder mounted in the VM

# **CreateSharedFolder**
> []SharedFolder CreateSharedFolder(ctx, body, id, optional)
Mounts a new shared folder in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedFolder**](SharedFolder.md)| Parameters of the shared folder to mount | 
  **id** | **string**| ID of VM | 
 **optional** | ***VMSharedFoldersManagementApiCreateSharedFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMSharedFoldersManagementApiCreateSharedFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**[]SharedFolder**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSharedFolder**
> DeleteSharedFolder(ctx, id, folderId, optional)
Deletes a shared folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **folderId** | **string**| ID of shared folder | 
 **optional** | ***VMSharedFoldersManagementApiDeleteSharedFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMSharedFoldersManagementApiDeleteSharedFolderOpts struct
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

# **GetAllSharedFolders**
> []SharedFolder GetAllSharedFolders(ctx, id, optional)
Returns all shared folders mounted in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
 **optional** | ***VMSharedFoldersManagementApiGetAllSharedFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMSharedFoldersManagementApiGetAllSharedFoldersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**[]SharedFolder**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdataSharedFolder**
> []SharedFolder UpdataSharedFolder(ctx, body, id, folderId, optional)
Updates a shared folder mounted in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedFolderParameter**](SharedFolderParameter.md)| Parameters of the shared folder to update to | 
  **id** | **string**| ID of VM | 
  **folderId** | **string**| ID of VM shared folder | 
 **optional** | ***VMSharedFoldersManagementApiUpdataSharedFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VMSharedFoldersManagementApiUpdataSharedFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vmPassword** | [**optional.Interface of string**](.md)| VM password | 

### Return type

[**[]SharedFolder**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

