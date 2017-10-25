# \TwoStepVerificationApi

All URIs are relative to *https://auth.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTV1TwostepverificationResend**](TwoStepVerificationApi.md#POSTV1TwostepverificationResend) | **Post** /v1/twostepverification/resend | Resends a two step verification code.
[**POSTV1TwostepverificationVerify**](TwoStepVerificationApi.md#POSTV1TwostepverificationVerify) | **Post** /v1/twostepverification/verify | Verifies a two step verification code.


# **POSTV1TwostepverificationResend**
> TwoStepVerificationSentResponse POSTV1TwostepverificationResend($request)

Resends a two step verification code.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**TwoStepVerificationTicketRequestModel**](TwoStepVerificationTicketRequestModel.md)| The request. | 

### Return type

[**TwoStepVerificationSentResponse**](TwoStepVerificationSentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1TwostepverificationVerify**
> ApiEmptyResponseModel POSTV1TwostepverificationVerify($request)

Verifies a two step verification code.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**TwoStepVerificationVerifyRequestModel**](TwoStepVerificationVerifyRequestModel.md)| The request model containing information needed to verify with two step verification. | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

