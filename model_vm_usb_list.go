/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vmrest

type VmUsbList struct {
	Num int32 `json:"num,omitempty"`
	UsbDevices []VmUsbDevice `json:"usbDevices,omitempty"`
}
