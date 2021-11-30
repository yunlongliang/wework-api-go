# GetUserRsp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | **int32** |  | [optional] [default to null]
**Errmsg** | **string** |  | [optional] [default to null]
**Userid** | **string** | 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。 | [optional] [default to null]
**Name** | **string** | 成员名称。长度为1~64个utf8字符 | [optional] [default to null]
**Department** | **[]int32** | 成员所属部门id列表,不超过20个 | [optional] [default to null]
**Order** | **[]int32** | 部门内的排序值 | [optional] [default to null]
**Position** | **string** | 职务信息 | [optional] [default to null]
**Mobile** | **string** | 手机号码。企业内必须唯一，mobile/email二者不能同时为空 | [optional] [default to null]
**Gender** | **int32** | 性别 | [optional] [default to null]
**Email** | **string** | 邮箱 | [optional] [default to null]
**IsLeaderInDept** | **int32** | 个数必须和参数department的个数一致 | [optional] [default to null]
**Avatar** | **string** | 成员头像的mediaid | [optional] [default to null]
**ThumbAvatar** | **string** | 头像缩略图url | [optional] [default to null]
**Telephone** | **string** | 座机 | [optional] [default to null]
**Alias** | **string** | 成员别名。长度1~32个utf8字符 | [optional] [default to null]
**Address** | **string** | 地址 | [optional] [default to null]
**OpenUserid** | **string** | 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取 | [optional] [default to null]
**MainDepartment** | **string** | 主部门 | [optional] [default to null]
**Status** | **int32** | 激活状态 1&#x3D;已激活，2&#x3D;已禁用，4&#x3D;未激活，5&#x3D;退出企业。 | [optional] [default to null]
**QrCode** | **string** | 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；第三方仅通讯录应用可获取。 | [optional] [default to null]
**ExternalPosition** | **string** | 对外职务 | [optional] [default to null]
**ExternalProfile** | [***ExternalProfile**](ExternalProfile.md) | 成员对外属性 | [optional] [default to null]
**Extattr** | [***ExtAttrs**](ExtAttrs.md) | 自定义字段 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


