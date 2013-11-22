package instagram

import (
	"fmt"
	"net/url"
)

// Get information about a tag object.
// Gets /tags/{tag-name}
func (api *Api) GetTag(tagName string, params url.Values) (res *TagResponse, err error) {
	res = new(TagResponse)
	err = api.get(fmt.Sprintf("/tags/%s", tagName), params, res)
	return
}

// Get a list of recently tagged media. Note that this media is ordered by when the media was tagged with this tag, rather than the order it was posted. Use the max_tag_id and min_tag_id parameters in the pagination response to paginate through these objects. Can return a mix of image and video types.
// Gets /tags/{tag-name}/media/recent
func (api *Api) GetTagRecentMedia(tagName string, params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get(fmt.Sprintf("/tags/%s/media/recent", tagName), params, res)
	return
}

// Search for tags by name. Results are ordered first as an exact match, then by popularity. Short tags will be treated as exact matches.
// Gets /tags/search
func (api *Api) GetTagSearch(params url.Values) (res *TagsResponse, err error) {
	res = new(TagsResponse)
	err = api.get("/tags/search", params, res)
	return
}
