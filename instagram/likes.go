package instagram

import (
	"fmt"
	"net/url"
)

// Get a list of users who have liked this media.
// Required Scope: likes
// Gets /media/{media-id}/likes
func (api *Api) GetMediaLikes(mediaId string, params url.Values) (res *UsersResponse, err error) {
	res = new(UsersResponse)
	err = api.get(fmt.Sprintf("/media/%s/likes", mediaId), params, res)
	return
}

// PostMediaLike ...
func (api *Api) PostMediaLike(mediaId string, params url.Values) (res *UsersResponse, err error) {
	res = new(UsersResponse)
	err = api.post(fmt.Sprintf("/media/%s/likes", mediaId), params, res)
	return
}

// DeleteMediaLike ...
func (api *Api) DeleteMediaLike(mediaId string, params url.Values) (res *UsersResponse, err error) {
	res = new(UsersResponse)
	err = api.delete(fmt.Sprintf("/media/%s/likes", mediaId), nil, res)
	return
}
