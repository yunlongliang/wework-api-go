# \TagApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagUsers**](TagApi.md#AddTagUsers) | **Post** /tag/addtagusers | 增加标签成员
[**CreateTag**](TagApi.md#CreateTag) | **Post** /tag/create | 创建内部标签
[**DeleteTag**](TagApi.md#DeleteTag) | **Get** /tag/delete | 删除标签
[**DeleteTagUsers**](TagApi.md#DeleteTagUsers) | **Post** /tag/deltagusers | 删除标签成员
[**GetTagList**](TagApi.md#GetTagList) | **Get** /tag/list | 获取标签列表
[**GetTagUsers**](TagApi.md#GetTagUsers) | **Get** /tag/get | 获取标签成员
[**UpdateTag**](TagApi.md#UpdateTag) | **Post** /tag/update | 更新内部标签名字


# **AddTagUsers**
> BaseResponse AddTagUsers(ctx, accessToken, body)
增加标签成员

增加标签成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**AddTagUsersReq**](AddTagUsersReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTag**
> CreateTagRsp CreateTag(ctx, accessToken, body)
创建内部标签

创建企业内部标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**CreateTagReq**](CreateTagReq.md)|  | 

### Return type

[**CreateTagRsp**](CreateTagRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> BaseResponse DeleteTag(ctx, accessToken, tagid)
删除标签

删除企业内部标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **tagid** | **int32**|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagUsers**
> BaseResponse DeleteTagUsers(ctx, accessToken, body)
删除标签成员

删除标签成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**DeleteTagUsersReq**](DeleteTagUsersReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagList**
> GetTagListRsp GetTagList(ctx, accessToken)
获取标签列表

获取标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 

### Return type

[**GetTagListRsp**](GetTagListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagUsers**
> GetTagUsersRsp GetTagUsers(ctx, accessToken, tagid)
获取标签成员

获取标签成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **tagid** | **string**|  | 

### Return type

[**GetTagUsersRsp**](GetTagUsersRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTag**
> BaseResponse UpdateTag(ctx, accessToken, body)
更新内部标签名字

更新内部标签名字

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**UpdateTagReq**](UpdateTagReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

