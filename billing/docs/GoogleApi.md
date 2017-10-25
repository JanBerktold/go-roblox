# \GoogleApi

All URIs are relative to *https://billing.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTV1GooglePurchase**](GoogleApi.md#POSTV1GooglePurchase) | **Post** /v1/google/purchase | Perform a purchase and grant desired product.
[**POSTV1GoogleValidate**](GoogleApi.md#POSTV1GoogleValidate) | **Post** /v1/google/validate | Validate a ProductId before making a purchase.


# **POSTV1GooglePurchase**
> ApiEmptyResponseModel POSTV1GooglePurchase($purchaseModel)

Perform a purchase and grant desired product.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseModel** | [**GooglePlayPurchaseModel**](GooglePlayPurchaseModel.md)|  | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1GoogleValidate**
> ApiEmptyResponseModel POSTV1GoogleValidate($validateModel)

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

