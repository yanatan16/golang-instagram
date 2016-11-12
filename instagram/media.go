package instagram

import (
	"fmt"
	"net/url"
)

// Get information about a media object. The returned type key will allow you to differentiate between image and video media.
// Note: if you authenticate with an OAuth Token, you will receive the user_has_liked key which quickly tells you whether the current user has liked this media item.
// Gets /media/{media-id}
func (api *Api) GetMedia(mediaId string, params url.Values) (res *MediaResponse, err error) {
	res = new(MediaResponse)
	err = api.get(fmt.Sprintf("/media/%s", mediaId), params, res)
	return
}

// Search for media in a given area. The default time span is set to 5 days. The time span must not exceed 7 days. Defaults time stamps cover the last 5 days. Can return mix of image and video types.
// Gets /media/search
func (api *Api) GetMediaSearch(params url.Values) (res *MediasResponse, err error) {
	res = new(MediasResponse)
	err = api.get("/media/search", params, res)
	return
}

// No available endpoints in the new IG Api
// reference: https://www.instagram.com/developer/endpoints/media/
// Get a list of what media is most popular at the moment. Can return mix of image and video types.
// Gets /media/popular
// func (api *Api) GetMediaPopular(params url.Values) (res *MediasResponse, err error) {
// 	res = new(MediasResponse)
// 	err = api.get("/media/popular", params, res)
// 	return
// }
