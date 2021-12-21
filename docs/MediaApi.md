# \MediaApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMedia**](MediaApi.md#GetMedia) | **Get** /media/get | 获取临时素材
[**GetMediaJssdk**](MediaApi.md#GetMediaJssdk) | **Get** /media/get/jssdk | 获取高清语音素材
[**UploadImgMedia**](MediaApi.md#UploadImgMedia) | **Post** /media/uploadimg | 上传图片
[**UploadMedia**](MediaApi.md#UploadMedia) | **Post** /media/upload | 上传临时素材


# **GetMedia**
> *os.File GetMedia(ctx, accessToken, mediaId)
获取临时素材

获取临时素材

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **mediaId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaJssdk**
> *os.File GetMediaJssdk(ctx, accessToken, mediaId)
获取高清语音素材

可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **mediaId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadImgMedia**
> MediaUploadImgRsp UploadImgMedia(ctx, accessToken, filename)
上传图片

上传图片得到图片URL，该URL永久有效;返回的图片URL，仅能用于图文消息正文中的图片展示，或者给客户发送欢迎语等；若用于非企业微信环境下的页面，图片将被屏蔽。每个企业每天最多可上传100张图片

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **filename** | ***os.File**|  | 

### Return type

[**MediaUploadImgRsp**](MediaUploadImgRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadMedia**
> MediaUploadRsp UploadMedia(ctx, accessToken, type_, optional)
上传临时素材

素材上传得到media_id，该media_id仅三天内有效;media_id在同一企业内应用之间可以共享

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **type_** | **string**| 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file） | 
 **optional** | ***MediaApiUploadMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaApiUploadMediaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filename** | **optional.Interface of *os.File**|  | 

### Return type

[**MediaUploadRsp**](MediaUploadRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

