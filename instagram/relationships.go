package instagram

import (
	"fmt"
	"net/url"
)

// Get the list of users this user follows.
// Required Scope: relationships
// Gets /users/{user-id}/follows
func (api *Api) GetUserFollows(userId string, params url.Values) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.get(fmt.Sprintf("/users/%s/follows", userId), params, res)
	return
}

// Get the list of users this user follows.
// Required Scope: relationships
// Gets /users/{user-id}/followed-by
func (api *Api) GetUserFollowedBy(userId string, params url.Values) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.get(fmt.Sprintf("/users/%s/followed-by", userId), params, res)
	return
}

// List the users who have requested this user's permission to follow.
// Required Scope: relationships
// Gets /users/self/requested-by
func (api *Api) GetUserRequestedBy(params url.Values) (res *UsersResponse, err error) {
	res = new(UsersResponse)
	err = api.get("/users/self/requested-by", params, res)
	return
}

// Get information about a relationship to another user.
// Required Scope: relationships
// Gets /users/{user-id}/relationship
func (api *Api) GetUserRelationship(userId string, params url.Values) (res *RelationshipResponse, err error) {
	res = new(RelationshipResponse)
	err = api.get(fmt.Sprintf("/users/%s/relationship", userId), params, res)
	return
}
