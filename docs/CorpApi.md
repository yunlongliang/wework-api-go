# \CorpApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJoinQrcode**](CorpApi.md#GetJoinQrcode) | **Get** /corp/get_join_qrcode | 获取加入企业二维码


# **GetJoinQrcode**
> GetJoinQrcodeRsp GetJoinQrcode(ctx, accessToken, optional)
获取加入企业二维码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***CorpApiGetJoinQrcodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CorpApiGetJoinQrcodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sizeType** | **optional.String**|  | 

### Return type

[**GetJoinQrcodeRsp**](GetJoinQrcodeRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

