# \UserApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthsuccUser**](UserApi.md#AuthsuccUser) | **Get** /user/authsucc | 二次验证
[**ConvertUseridToOpenid**](UserApi.md#ConvertUseridToOpenid) | **Post** /user/convert_to_openid | userid与openid互换
[**CreateUser**](UserApi.md#CreateUser) | **Post** /user/create | 创建成员
[**DeleteBatchUser**](UserApi.md#DeleteBatchUser) | **Post** /user/batchdelete | 批量删除成员
[**DeleteUser**](UserApi.md#DeleteUser) | **Get** /user/delete | 删除成员
[**GetActiveStat**](UserApi.md#GetActiveStat) | **Post** /user/get_active_stat | 获取企业活跃成员数
[**GetUser**](UserApi.md#GetUser) | **Get** /user/get | 读取成员
[**GetUserList**](UserApi.md#GetUserList) | **Get** /user/list | 获取部门成员详情
[**GetUserSimplelist**](UserApi.md#GetUserSimplelist) | **Get** /user/simplelist | 获取部门成员
[**InviteUser**](UserApi.md#InviteUser) | **Post** /batch/invite | 邀请成员
[**ReplaceBatchUser**](UserApi.md#ReplaceBatchUser) | **Post** /batch/replaceuser | 全量覆盖成员
[**SyncBatchUser**](UserApi.md#SyncBatchUser) | **Post** /batch/syncuser | 增量更新成员
[**UpdateUser**](UserApi.md#UpdateUser) | **Post** /user/update | 更新成员


# **AuthsuccUser**
> BaseResponse AuthsuccUser(ctx, accessToken, userid)
二次验证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertUseridToOpenid**
> ConvertToOpenidRsp ConvertUseridToOpenid(ctx, accessToken, body)
userid与openid互换

userid转openid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**ConvertToOpenidReq**](ConvertToOpenidReq.md)|  | 

### Return type

[**ConvertToOpenidRsp**](ConvertToOpenidRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> BaseResponse CreateUser(ctx, accessToken, body)
创建成员

创建成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**CreateUserReq**](CreateUserReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBatchUser**
> BaseResponse DeleteBatchUser(ctx, body)
批量删除成员

批量删除成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchDeleteUserReq**](BatchDeleteUserReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> BaseResponse DeleteUser(ctx, accessToken, userid)
删除成员

删除成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveStat**
> GetActiveStatRsp GetActiveStat(ctx, accessToken, date)
获取企业活跃成员数

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **date** | **string**|  | 

### Return type

[**GetActiveStatRsp**](GetActiveStatRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> GetUserRsp GetUser(ctx, accessToken, userid)
读取成员

读取成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**GetUserRsp**](GetUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserList**
> ListUserRsp GetUserList(ctx, accessToken, departmentId, optional)
获取部门成员详情

获取部门成员详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**| 调用接口凭证 | 
  **departmentId** | **int32**| 获取的部门id | 
 **optional** | ***UserApiGetUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fetchChild** | **optional.Int32**| 1/0：是否递归获取子部门下面的成员 | 

### Return type

[**ListUserRsp**](ListUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSimplelist**
> SimplelistRsp GetUserSimplelist(ctx, accessToken, departmentId, optional)
获取部门成员

获取部门成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **departmentId** | **int32**|  | 
 **optional** | ***UserApiGetUserSimplelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserSimplelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fetchChild** | **optional.Int32**|  | 

### Return type

[**SimplelistRsp**](SimplelistRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteUser**
> InviteRsp InviteUser(ctx, accessToken, body)
邀请成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**InviteReq**](InviteReq.md)|  | 

### Return type

[**InviteRsp**](InviteRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBatchUser**
> BatchReplaceuserRsp ReplaceBatchUser(ctx, accessToken, optional)
全量覆盖成员

本接口以userid为主键，全量覆盖企业的通讯录成员，任务完成后企业的通讯录成员与提交的文件完全保持一致。请先下载CSV文件(下载全量覆盖成员模版)，根据需求填写文件内容。注意事项：1.模板中的部门需填写部门ID，多个部门用分号分隔，部门ID必须为数字，根部门的部门id默认为1;2.文件中存在、通讯录中也存在的成员，完全以文件为准;3.文件中存在、通讯录中不存在的成员，执行添加操作;4.通讯录中存在、文件中不存在的成员，执行删除操作。出于安全考虑，下面两种情形系统将中止导入并返回相应的错误码。4.1.需要删除的成员多于50人，且多于现有人数的20%以上;4.2.需要删除的成员少于50人，且多于现有人数的80%以上5.成员字段更新规则：可自行添加扩展字段。文件中有指定的字段，以指定的字段值为准；文件中没指定的字段，不更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***UserApiReplaceBatchUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiReplaceBatchUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BatchReplaceuserReq**](BatchReplaceuserReq.md)|  | 

### Return type

[**BatchReplaceuserRsp**](BatchReplaceuserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncBatchUser**
> BatchSyncuserRsp SyncBatchUser(ctx, accessToken, optional)
增量更新成员

本接口以userid（帐号）为主键，增量更新企业微信通讯录成员。请先下载CSV模板(下载增量更新成员模版)，根据需求填写文件内容。注意事项：1.模板中的部门需填写部门ID，多个部门用分号分隔，部门ID必须为数字，根部门的部门id默认为1;2.文件中存在、通讯录中也存在的成员，更新成员在文件中指定的字段值;3.文件中存在、通讯录中不存在的成员，执行添加操作;4.通讯录中存在、文件中不存在的成员，保持不变;5.成员字段更新规则：可自行添加扩展字段。文件中有指定的字段，以指定的字段值为准；文件中没指定的字段，不更新;

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***UserApiSyncBatchUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSyncBatchUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BatchSyncuserReq**](BatchSyncuserReq.md)|  | 

### Return type

[**BatchSyncuserRsp**](BatchSyncuserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> BaseResponse UpdateUser(ctx, body)
更新成员

更新成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserReq**](UpdateUserReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

