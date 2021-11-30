# \MenuApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMenu**](MenuApi.md#CreateMenu) | **Post** /menu/create | 创建菜单
[**DeleteMenu**](MenuApi.md#DeleteMenu) | **Get** /menu/delete | 删除菜单
[**GetMenu**](MenuApi.md#GetMenu) | **Get** /menu/get | 获取菜单


# **CreateMenu**
> CreateMenuRsp CreateMenu(ctx, accessToken, optional)
创建菜单

自定义菜单接口可实现多种类型按钮

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***MenuApiCreateMenuOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiCreateMenuOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateMenuReq**](CreateMenuReq.md)|  | 

### Return type

[**CreateMenuRsp**](CreateMenuRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMenu**
> BaseResponse DeleteMenu(ctx, accessToken, agentid)
删除菜单

删除菜单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **agentid** | **string**|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMenu**
> GetMenuRsp GetMenu(ctx, accessToken, agentid)
获取菜单

获取菜单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **agentid** | **string**|  | 

### Return type

[**GetMenuRsp**](GetMenuRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

