package instagram

import (
	"fmt"
	"net/url"
)

// for relationships, has been modified the scope in IG new Endpoint API
// reference: https://www.instagram.com/developer/endpoints/relationships/

// Get the list of users this user follows.
// Required Scope: relationships
// Gets /users/self/follows
func (api *Api) GetUserFollows(params url.Values) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.get(fmt.Sprintf("/users/self/follows"), params, res)
	return
}

// Get the list of users this user follows.
// Required Scope: relationships
// Gets /users/self/followed-by
func (api *Api) GetUserFollowedBy(params url.Values) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.get(fmt.Sprintf("/users/self/followed-by"), params, res)
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
