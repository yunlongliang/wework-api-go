/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetLinkedcorpDepartmentListRsp struct {
	Errcode float32 `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
	DepartmentList []interface{} `json:"department_list,omitempty"`
}
