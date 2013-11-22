package instagram

import (
	"net/url"
)

// Get the next page of media
func (api *Api) NextMedias(mp *MediaPagination) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.next(mp.Pagination, res)
	return
}

// Get the next page of user
func (api *Api) NextUsers(up *UserPagination) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.next(up.Pagination, res)
	return
}

func (api *Api) next(p *Pagination, res interface{}) error {
	done, uri, uriParams, err := p.NextPage()
	if err != nil || done == true {
		return err
	}

	req, err := buildGetRequest(uri, uriParams)
	if err != nil {
		return err
	}

	err = api.do(req, res)
	return nil
}

// Return the next page uri and parameters
func (p *Pagination) NextPage() (done bool, uri string, params url.Values, err error) {
	if p == nil || p.NextUrl == "" {
		// We're done. Theres no more pages
		done = true
		return
	}

	urlStruct, err := url.Parse(p.NextUrl)
	if err != nil {
		return
	}

	params = urlStruct.Query()
	urlStruct.RawQuery = ""

	done = false
	uri = urlStruct.String()
	return
}
