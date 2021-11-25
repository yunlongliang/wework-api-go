/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TransferCustomerReq struct {
	HandoverUserid string `json:"handover_userid,omitempty"`
	TakeoverUserid string `json:"takeover_userid,omitempty"`
	ExternalUserid []string `json:"external_userid,omitempty"`
	TransferSuccessMsg string `json:"transfer_success_msg,omitempty"`
}
