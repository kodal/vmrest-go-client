/*
 * VMware Fusion REST API
 *
 * vmrest 1.3.1 build-24585314
 *
 * API version: 1.3.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SharedFolder struct {
	// ID of folder which be mounted to the host
	FolderId string `json:"folder_id"`
	// Path of the host shared folder
	HostPath string `json:"host_path"`
	// The flags property specifies how the folder will be accessed by the VM. There is only one flag supported which is \"4\" and means read/write access. 
	Flags int32 `json:"flags"`
}
