/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vmrest

type VmApplianceView struct {
	Author string `json:"author,omitempty"`
	Version string `json:"version,omitempty"`
	Port int32 `json:"port,omitempty"`
	ShowAtPowerOn string `json:"showAtPowerOn,omitempty"`
}
