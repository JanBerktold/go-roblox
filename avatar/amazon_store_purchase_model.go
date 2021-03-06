/* 
 * Roblox.Api.Billing V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// A model containing information about making a purchase through Amazon store
type AmazonStorePurchaseModel struct {

	// ReceiptId
	ReceiptId string `json:"receiptId,omitempty"`

	// UserId
	AmazonUserId string `json:"amazonUserId,omitempty"`

	// Is the purchase a Retry
	IsRetry bool `json:"isRetry,omitempty"`
}
