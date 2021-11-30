# CreateUserReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Userid** | **string** | 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。 | [optional] [default to null]
**Name** | **string** | 成员名称。长度为1~64个utf8字符 | [optional] [default to null]
**Alias** | **string** | 成员别名。长度1~32个utf8字符 | [optional] [default to null]
**Mobile** | **string** | 手机号码。企业内必须唯一，mobile/email二者不能同时为空 | [optional] [default to null]
**Department** | **[]int32** | 成员所属部门id列表,不超过20个 | [optional] [default to null]
**Order** | **[]int32** | 部门内的排序值 | [optional] [default to null]
**Position** | **string** | 职务信息 | [optional] [default to null]
**Gender** | **int32** | 性别 | [optional] [default to null]
**Email** | **string** | 邮箱 | [optional] [default to null]
**Telephone** | **string** | 座机 | [optional] [default to null]
**IsLeaderInDept** | **int32** | 个数必须和参数department的个数一致 | [optional] [default to null]
**AvatarMediaid** | **string** | 成员头像的mediaid | [optional] [default to null]
**Enable** | **int32** | 启用/禁用成员。1表示启用成员，0表示禁用成员 | [optional] [default to null]
**Extattr** | [***ExtAttrs**](ExtAttrs.md) | 自定义字段 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


