/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vmrest

// Information about all NICs
type NicIpStackAll struct {
	Nics *NicIpStack `json:"nics,omitempty"`
	Routes []RouteEntry `json:"routes,omitempty"`
	Dns *DnsConfig `json:"dns,omitempty"`
	Wins *WinsConfig `json:"wins,omitempty"`
	Dhcpv4 *DhcpConfig `json:"dhcpv4,omitempty"`
	Dhcpv6 *DhcpConfig `json:"dhcpv6,omitempty"`
}
