# Go API client for swagger

vmrest 1.3.1 build-24585314

## GENERATION
```
swagger-codegen generate -i http://localhost:8697/json/swagger.json -l go
```

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.3.1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HostNetworksManagementApi* | [**CreateNetwork**](docs/HostNetworksManagementApi.md#createnetwork) | **Post** /vmnets | Creates a virtual network
*HostNetworksManagementApi* | [**DeletePortforward**](docs/HostNetworksManagementApi.md#deleteportforward) | **Delete** /vmnet/{vmnet}/portforward/{protocol}/{port} | Deletes port forwarding
*HostNetworksManagementApi* | [**GetAllNetworks**](docs/HostNetworksManagementApi.md#getallnetworks) | **Get** /vmnet | Returns all virtual networks
*HostNetworksManagementApi* | [**GetMACToIPs**](docs/HostNetworksManagementApi.md#getmactoips) | **Get** /vmnet/{vmnet}/mactoip | Returns all MAC-to-IP settings for DHCP service
*HostNetworksManagementApi* | [**GetPortforwards**](docs/HostNetworksManagementApi.md#getportforwards) | **Get** /vmnet/{vmnet}/portforward | Returns all port forwardings
*HostNetworksManagementApi* | [**UpdateMacToIP**](docs/HostNetworksManagementApi.md#updatemactoip) | **Put** /vmnet/{vmnet}/mactoip/{mac} | Updates the MAC-to-IP binding
*HostNetworksManagementApi* | [**UpdatePortforward**](docs/HostNetworksManagementApi.md#updateportforward) | **Put** /vmnet/{vmnet}/portforward/{protocol}/{port} | Updates port forwarding
*VMManagementApi* | [**ConfigVMParams**](docs/VMManagementApi.md#configvmparams) | **Put** /vms/{id}/params | update the vm config params
*VMManagementApi* | [**CreateVM**](docs/VMManagementApi.md#createvm) | **Post** /vms | Creates a copy of the VM
*VMManagementApi* | [**DeleteVM**](docs/VMManagementApi.md#deletevm) | **Delete** /vms/{id} | Deletes a VM
*VMManagementApi* | [**GetAllVMs**](docs/VMManagementApi.md#getallvms) | **Get** /vms | Returns a list of VM IDs and paths for all VMs
*VMManagementApi* | [**GetVM**](docs/VMManagementApi.md#getvm) | **Get** /vms/{id} | Returns the VM setting information of a VM
*VMManagementApi* | [**GetVMParams**](docs/VMManagementApi.md#getvmparams) | **Get** /vms/{id}/params/{name} | Get the VM config params
*VMManagementApi* | [**GetVMRestrictions**](docs/VMManagementApi.md#getvmrestrictions) | **Get** /vms/{id}/restrictions | Returns the restrictions information of the VM
*VMManagementApi* | [**RegisterVM**](docs/VMManagementApi.md#registervm) | **Post** /vms/registration | Register VM to VM Library
*VMManagementApi* | [**UpdateVM**](docs/VMManagementApi.md#updatevm) | **Put** /vms/{id} | Updates the VM settings
*VMNetworkAdaptersManagementApi* | [**CreateNICDevice**](docs/VMNetworkAdaptersManagementApi.md#createnicdevice) | **Post** /vms/{id}/nic | Creates a network adapter in the VM
*VMNetworkAdaptersManagementApi* | [**DeleteNICDevice**](docs/VMNetworkAdaptersManagementApi.md#deletenicdevice) | **Delete** /vms/{id}/nic/{index} | Deletes a VM network adapter
*VMNetworkAdaptersManagementApi* | [**GetAllNICDevices**](docs/VMNetworkAdaptersManagementApi.md#getallnicdevices) | **Get** /vms/{id}/nic | Returns all network adapters in the VM
*VMNetworkAdaptersManagementApi* | [**GetIPAddress**](docs/VMNetworkAdaptersManagementApi.md#getipaddress) | **Get** /vms/{id}/ip | Returns the IP address of a VM
*VMNetworkAdaptersManagementApi* | [**GetNicInfo**](docs/VMNetworkAdaptersManagementApi.md#getnicinfo) | **Get** /vms/{id}/nicips | Returns the IP stack configuration of all NICs of a VM
*VMNetworkAdaptersManagementApi* | [**UpdateNICDevice**](docs/VMNetworkAdaptersManagementApi.md#updatenicdevice) | **Put** /vms/{id}/nic/{index} | Updates a network adapter in the VM
*VMPowerManagementApi* | [**ChangePowerState**](docs/VMPowerManagementApi.md#changepowerstate) | **Put** /vms/{id}/power | Changes the VM power state
*VMPowerManagementApi* | [**GetPowerState**](docs/VMPowerManagementApi.md#getpowerstate) | **Get** /vms/{id}/power | Returns the power state of the VM
*VMSharedFoldersManagementApi* | [**CreateSharedFolder**](docs/VMSharedFoldersManagementApi.md#createsharedfolder) | **Post** /vms/{id}/sharedfolders | Mounts a new shared folder in the VM
*VMSharedFoldersManagementApi* | [**DeleteSharedFolder**](docs/VMSharedFoldersManagementApi.md#deletesharedfolder) | **Delete** /vms/{id}/sharedfolders/{folder id} | Deletes a shared folder
*VMSharedFoldersManagementApi* | [**GetAllSharedFolders**](docs/VMSharedFoldersManagementApi.md#getallsharedfolders) | **Get** /vms/{id}/sharedfolders | Returns all shared folders mounted in the VM
*VMSharedFoldersManagementApi* | [**UpdataSharedFolder**](docs/VMSharedFoldersManagementApi.md#updatasharedfolder) | **Put** /vms/{id}/sharedfolders/{folder id} | Updates a shared folder mounted in the VM

## Documentation For Models

 - [ConfigVmParamsParameter](docs/ConfigVmParamsParameter.md)
 - [CreateVmnetParameter](docs/CreateVmnetParameter.md)
 - [DaemonState](docs/DaemonState.md)
 - [DhcpConfig](docs/DhcpConfig.md)
 - [DnsConfig](docs/DnsConfig.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [MacToIpParameter](docs/MacToIpParameter.md)
 - [MactoIp](docs/MactoIp.md)
 - [MactoIps](docs/MactoIps.md)
 - [Network](docs/Network.md)
 - [Networks](docs/Networks.md)
 - [NicDevice](docs/NicDevice.md)
 - [NicDeviceParameter](docs/NicDeviceParameter.md)
 - [NicDevices](docs/NicDevices.md)
 - [NicIpStack](docs/NicIpStack.md)
 - [NicIpStackAll](docs/NicIpStackAll.md)
 - [Portforward](docs/Portforward.md)
 - [PortforwardGuest](docs/PortforwardGuest.md)
 - [PortforwardParameter](docs/PortforwardParameter.md)
 - [Portforwards](docs/Portforwards.md)
 - [RouteEntry](docs/RouteEntry.md)
 - [SharedFolder](docs/SharedFolder.md)
 - [SharedFolderParameter](docs/SharedFolderParameter.md)
 - [VmApplianceView](docs/VmApplianceView.md)
 - [VmCloneParameter](docs/VmCloneParameter.md)
 - [VmConnectedDevice](docs/VmConnectedDevice.md)
 - [VmConnectedDeviceList](docs/VmConnectedDeviceList.md)
 - [VmGuestIsolation](docs/VmGuestIsolation.md)
 - [VmInformation](docs/VmInformation.md)
 - [VmParameter](docs/VmParameter.md)
 - [VmPowerOperation](docs/VmPowerOperation.md)
 - [VmPowerState](docs/VmPowerState.md)
 - [VmRegisterParameter](docs/VmRegisterParameter.md)
 - [VmRemoteVnc](docs/VmRemoteVnc.md)
 - [VmRestrictionsInformation](docs/VmRestrictionsInformation.md)
 - [VmRrgistrationInformation](docs/VmRrgistrationInformation.md)
 - [VmUsbDevice](docs/VmUsbDevice.md)
 - [VmUsbList](docs/VmUsbList.md)
 - [Vmcpu](docs/Vmcpu.md)
 - [Vmid](docs/Vmid.md)
 - [WinsConfig](docs/WinsConfig.md)

## Documentation For Authorization

## BasicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author


