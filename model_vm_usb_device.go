/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VmUsbDevice struct {
	Index int32 `json:"index,omitempty"`
	Connected string `json:"connected,omitempty"`
	BackingInfo string `json:"backingInfo,omitempty"`
	BackingType int32 `json:"BackingType,omitempty"`
}
