/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The list of virtual networks
type Networks struct {
	Num int32 `json:"num"`
	// The list of virtual networks
	Vmnets []Network `json:"vmnets"`
}
