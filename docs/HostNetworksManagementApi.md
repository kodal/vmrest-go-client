# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](HostNetworksManagementApi.md#CreateNetwork) | **Post** /vmnets | Creates a virtual network
[**DeletePortforward**](HostNetworksManagementApi.md#DeletePortforward) | **Delete** /vmnet/{vmnet}/portforward/{protocol}/{port} | Deletes port forwarding
[**GetAllNetworks**](HostNetworksManagementApi.md#GetAllNetworks) | **Get** /vmnet | Returns all virtual networks
[**GetMACToIPs**](HostNetworksManagementApi.md#GetMACToIPs) | **Get** /vmnet/{vmnet}/mactoip | Returns all MAC-to-IP settings for DHCP service
[**GetPortforwards**](HostNetworksManagementApi.md#GetPortforwards) | **Get** /vmnet/{vmnet}/portforward | Returns all port forwardings
[**UpdateMacToIP**](HostNetworksManagementApi.md#UpdateMacToIP) | **Put** /vmnet/{vmnet}/mactoip/{mac} | Updates the MAC-to-IP binding
[**UpdatePortforward**](HostNetworksManagementApi.md#UpdatePortforward) | **Put** /vmnet/{vmnet}/portforward/{protocol}/{port} | Updates port forwarding

# **CreateNetwork**
> Network CreateNetwork(ctx, body)
Creates a virtual network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmnetParameter**](CreateVmnetParameter.md)| Host network to be created | 

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePortforward**
> DeletePortforward(ctx, vmnet, protocol, port)
Deletes port forwarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmnet** | **string**| NAT type of virtual network | 
  **protocol** | **string**| Protocol type: tcp, udp | 
  **port** | **int32**| Host port number | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNetworks**
> Networks GetAllNetworks(ctx, )
Returns all virtual networks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Networks**](Networks.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMACToIPs**
> MactoIps GetMACToIPs(ctx, vmnet)
Returns all MAC-to-IP settings for DHCP service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmnet** | **string**| Virtual network that has DHCP enabled | 

### Return type

[**MactoIps**](MACToIPs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortforwards**
> Portforwards GetPortforwards(ctx, vmnet)
Returns all port forwardings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmnet** | **string**| NAT type of virtual network | 

### Return type

[**Portforwards**](Portforwards.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMacToIP**
> ErrorModel UpdateMacToIP(ctx, body, vmnet, mac)
Updates the MAC-to-IP binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacToIpParameter**](MacToIpParameter.md)| IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted | 
  **vmnet** | **string**| Virtual network that enabled DHCP | 
  **mac** | **string**| Mac address that want to be mapped with a given IP | 

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortforward**
> ErrorModel UpdatePortforward(ctx, body, vmnet, protocol, port)
Updates port forwarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortforwardParameter**](PortforwardParameter.md)| Guest to forward to | 
  **vmnet** | **string**| NAT type of virtual network | 
  **protocol** | **string**| Protocol type: tcp, udp | 
  **port** | **int32**| Host port number | 

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

