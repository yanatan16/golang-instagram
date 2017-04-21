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

// PostMediaComment ...
func (api *Api) PostMediaComment(mediaId string, params url.Values) (res *CommentsResponse, err error) {
	res = new(CommentsResponse)
	err = api.post(fmt.Sprintf("/media/%s/comments", mediaId), params, res)
	return
}

// DeleteMediaComment ...
func (api *Api) DeleteMediaComment(mediaId string, commendId string) (res *CommentsResponse, err error) {
	res = new(CommentsResponse)
	err = api.delete(fmt.Sprintf("/media/%s/comments/%s", mediaId, commendId), nil, res)
	return
}
