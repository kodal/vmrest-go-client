/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The list of port forwarding
type Portforwards struct {
	Num int32 `json:"num"`
	// The list of port forwardings
	PortForwardings []Portforward `json:"port_forwardings"`
}
