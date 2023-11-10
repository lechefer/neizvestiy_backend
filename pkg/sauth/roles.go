package sauth

import "fmt"

// AsRole Example: resourceType = post; relation = owner; resourceId: 1
// Return: role_post_owner_1
// What's mean owner of post 1
func AsRole(resourceType, relation, resourceId string) string {
	return fmt.Sprintf("role_%s_%s_%s", resourceType, relation, resourceId)
}

// WithResourceType Example: resourceType = post; resourceId: 1
// Return: role_post_1
func WithResourceType(resourceType, resourceId string) string {
	return fmt.Sprintf("%s_%s", resourceType, resourceId)
}
