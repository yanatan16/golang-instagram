package instagram

import (
	"fmt"
	"net/url"
)

// Get basic information about a user.
// Gets /users/{user-id}
func (api *Api) GetUser(userId string, params url.Values) (res *UserResponse, err error) {
	res = new(UserResponse)
	err = api.get(fmt.Sprintf("/users/%s", userId), params, res)
	return
}

// Get basic information about authenticated user.
// Gets /users/self
func (api *Api) GetSelf() (res *UserResponse, err error) {
	return api.GetUser("self", nil)
}

// Endpoint is not available in the new IG Api
// reference: https://www.instagram.com/developer/endpoints
// See the authenticated user's feed. May return a mix of both image and video types.
// Gets /users/self/feed
// func (api *Api) GetUserFeed(params url.Values) (res *PaginatedMediasResponse, err error) {
// 	res = new(PaginatedMediasResponse)
// 	err = api.get("/users/self/feed", params, res)
// 	return
// }

// Get the most recent media published by a user. May return a mix of both image and video types.
// Gets /users/{user-id}/media/recent
func (api *Api) GetUserRecentMedia(userId string, params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get(fmt.Sprintf("/users/%s/media/recent", userId), params, res)
	return
}

// See the authenticated user's list of media they've liked. May return a mix of both image and video types.
// Note: This list is ordered by the order in which the user liked the media. Private media is returned as long as the authenticated user has permission to view that media. Liked media lists are only available for the currently authenticated user.
// Gets /users/self/media/liked
func (api *Api) GetUserLikedMedia(params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get("/users/self/media/liked", params, res)
	return
}

// Search for a user by name.
// Gets /users/search
func (api *Api) GetUserSearch(params url.Values) (res *UsersResponse, err error) {
	res = new(UsersResponse)
	err = api.get("/users/search", params, res)
	return
}

// Verify a valid client keys and user tokens by making a small request
func (api *Api) VerifyCredentials() (ok bool, err error) {
	_, err = api.GetSelf()
	return err == nil, err
}
