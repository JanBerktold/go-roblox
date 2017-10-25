# \AppleApi

All URIs are relative to *https://billing.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTV1ApplePurchase**](AppleApi.md#POSTV1ApplePurchase) | **Post** /v1/apple/purchase | Perform a purchase and grant desired product.
[**POSTV1AppleValidate**](AppleApi.md#POSTV1AppleValidate) | **Post** /v1/apple/validate | Validate a ProductId before making a purchase.


# **POSTV1ApplePurchase**
> ApiEmptyResponseModel POSTV1ApplePurchase($appleStorePurchaseModel)

Perform a purchase and grant desired product.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appleStorePurchaseModel** | [**AppleStorePurchaseModel**](AppleStorePurchaseModel.md)| Apple Purchase Model | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1AppleValidate**
> ApiEmptyResponseModel POSTV1AppleValidate($validateModel)

Validate a ProductId before making a purchase.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateModel** | [**ValidateModel**](ValidateModel.md)| Product Id for example: com.roblox.client.bc1month, com.roblox.client.tbc1month, ... | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

