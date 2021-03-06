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

import (
	"time"
)

type UniverseModel struct {

	// The universe Id
	Id int64 `json:"id,omitempty"`

	// The name of the universe
	Name string `json:"name,omitempty"`

	// The description of the universe
	Description string `json:"description,omitempty"`

	// The universe's root place id
	RootPlaceId int64 `json:"rootPlaceId,omitempty"`

	// Is this universe active
	IsActive bool `json:"isActive,omitempty"`

	// The creator type, either \"User\" or \"Group\"
	CreatorType string `json:"creatorType,omitempty"`

	// The id of the user or group that created this universe.
	CreatorTargetId int64 `json:"creatorTargetId,omitempty"`

	// The name of the creator of the universe.
	CreatorName string `json:"creatorName,omitempty"`

	// The created {System.DateTime}
	Created time.Time `json:"created,omitempty"`

	// The updated {System.DateTime}
	Updated time.Time `json:"updated,omitempty"`
}
