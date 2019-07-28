/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"time"
)

type IdentityGroup struct {
	AvatarUrl         string                 `json:"avatar_url,omitempty"`
	Company           string                 `json:"company,omitempty"`
	CreatedAt         time.Time              `json:"created_at,omitempty"`
	CustomFieldValues map[string]interface{} `json:"custom_field_values,omitempty"`
	Emails            []string               `json:"emails,omitempty"`
	Firstname         string                 `json:"firstname,omitempty"`
	Gender            string                 `json:"gender,omitempty"`
	HomePhones        []string               `json:"home_phones,omitempty"`
	Id                string                 `json:"id"`
	IdentityIds       []string               `json:"identity_ids,omitempty"`
	Lastname          string                 `json:"lastname,omitempty"`
	MobilePhones      []string               `json:"mobile_phones,omitempty"`
	Notes             string                 `json:"notes,omitempty"`
	TagIds            []string               `json:"tag_ids,omitempty"`
	UpdatedAt         time.Time              `json:"updated_at,omitempty"`
}
