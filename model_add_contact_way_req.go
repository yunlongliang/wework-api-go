/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AddContactWayReq struct {
	Type_ int32 `json:"type,omitempty"`
	Scene int32 `json:"scene,omitempty"`
	Style int32 `json:"style,omitempty"`
	Remark string `json:"remark,omitempty"`
	SkipVerify bool `json:"skip_verify,omitempty"`
	State string `json:"state,omitempty"`
	User []string `json:"user,omitempty"`
	Party []int32 `json:"party,omitempty"`
	// 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp,omitempty"`
	ExpiresIn int32 `json:"expires_in,omitempty"`
	// 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	Unionid string `json:"unionid,omitempty"`
	Conclusions interface{} `json:"conclusions,omitempty"`
}
