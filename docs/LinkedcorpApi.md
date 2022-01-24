# \LinkedcorpApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLinkedcorpDepartmentList**](LinkedcorpApi.md#GetLinkedcorpDepartmentList) | **Post** /linkedcorp/department/list | 获取互联企业部门列表
[**GetLinkedcorpPermList**](LinkedcorpApi.md#GetLinkedcorpPermList) | **Post** /linkedcorp/agent/get_perm_list | 获取应用的可见范围
[**GetLinkedcorpUser**](LinkedcorpApi.md#GetLinkedcorpUser) | **Post** /linkedcorp/user/get | 获取互联企业成员详细信息
[**GetLinkedcorpUserList**](LinkedcorpApi.md#GetLinkedcorpUserList) | **Post** /linkedcorp/user/list | 获取互联企业部门成员详情
[**GetLinkedcorpUserSimplelist**](LinkedcorpApi.md#GetLinkedcorpUserSimplelist) | **Post** /linkedcorp/user/simplelist | 获取互联企业部门成员


# **GetLinkedcorpDepartmentList**
> GetLinkedcorpDepartmentListRsp GetLinkedcorpDepartmentList(ctx, accessToken, optional)
获取互联企业部门列表

获取互联企业部门列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***LinkedcorpApiGetLinkedcorpDepartmentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinkedcorpApiGetLinkedcorpDepartmentListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetLinkedcorpDepartmentListReq**](GetLinkedcorpDepartmentListReq.md)|  | 

### Return type

[**GetLinkedcorpDepartmentListRsp**](GetLinkedcorpDepartmentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedcorpPermList**
> GetAgentPermListRsp GetLinkedcorpPermList(ctx, accessToken)
获取应用的可见范围

本接口只返回互联企业中非本企业内的成员和部门的信息，如果要获取本企业的可见范围，请调用“获取应用”接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 

### Return type

[**GetAgentPermListRsp**](GetAgentPermListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedcorpUser**
> GetLinkedcorpUserRsp GetLinkedcorpUser(ctx, accessToken, optional)
获取互联企业成员详细信息

获取互联企业成员详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***LinkedcorpApiGetLinkedcorpUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinkedcorpApiGetLinkedcorpUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetLinkedcorpUserReq**](GetLinkedcorpUserReq.md)|  | 

### Return type

[**GetLinkedcorpUserRsp**](GetLinkedcorpUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedcorpUserList**
> GetLinkedcorpUserListRsp GetLinkedcorpUserList(ctx, accessToken, optional)
获取互联企业部门成员详情

获取互联企业部门成员详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***LinkedcorpApiGetLinkedcorpUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinkedcorpApiGetLinkedcorpUserListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetLinkedcorpUserListReq**](GetLinkedcorpUserListReq.md)|  | 

### Return type

[**GetLinkedcorpUserListRsp**](GetLinkedcorpUserListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedcorpUserSimplelist**
> GetLinkedcorpUserSimplelistRsp GetLinkedcorpUserSimplelist(ctx, accessToken, optional)
获取互联企业部门成员

获取互联企业部门成员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***LinkedcorpApiGetLinkedcorpUserSimplelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinkedcorpApiGetLinkedcorpUserSimplelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetLinkedcorpUserSimplelistReq**](GetLinkedcorpUserSimplelistReq.md)|  | 

### Return type

[**GetLinkedcorpUserSimplelistRsp**](GetLinkedcorpUserSimplelistRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

