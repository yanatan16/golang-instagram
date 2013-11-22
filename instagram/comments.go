package instagram

import (
	"fmt"
	"net/url"
)

// Get a full list of comments on a media.
// Required Scope: comments
// Gets /media/{media-id}/comments
func (api *Api) GetMediaComments(mediaId string, params url.Values) (res *CommentsResponse, err error) {
	res = new(CommentsResponse)
	err = api.get(fmt.Sprintf("/media/%s/comments", mediaId), params, res)
	return
}
