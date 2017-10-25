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

// Model for UniverseSettings patch requests
type UniverseSettingsRequest struct {

	// The name of the universe.
	Name string `json:"name,omitempty"`

	// Which avatar types are allowed in the universe.
	UniverseAvatarType string `json:"universeAvatarType,omitempty"`

	// Whether custom scales allowed in the universe.
	UniverseScaleType string `json:"universeScaleType,omitempty"`

	// Whether custom animations are allowed in the universe.
	UniverseAnimationType string `json:"universeAnimationType,omitempty"`

	// What type of collisions are used by the universe.
	UniverseCollisionType string `json:"universeCollisionType,omitempty"`
}
