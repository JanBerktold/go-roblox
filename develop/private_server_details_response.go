/* 
 * Roblox.Api.Develop V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// Model for private server details responses from the UniverseSettings controller.
type PrivateServerDetailsResponse struct {

	// Whether or not VIP servers are enabled on this universe.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// The price of the VIP server.
	Price int64 `json:"price,omitempty"`

	// The number of active VIP servers for this universe.
	ActiveServersCount int64 `json:"activeServersCount,omitempty"`

	// The number of active VIP server subscriptions.
	ActiveSubscriptionsCount int64 `json:"activeSubscriptionsCount,omitempty"`
}
