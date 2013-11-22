package instagram

import (
	"fmt"
	"net/url"
)

// Get information about a location.
// Gets /locations/{location-id}
func (api *Api) GetLocation(locationId string, params url.Values) (res *LocationResponse, err error) {
	res = new(LocationResponse)
	err = api.get(fmt.Sprintf("/locations/%s", locationId), params, res)
	return
}

// Get a list of recent media objects from a given location. May return a mix of both image and video types.
// Gets /locations/{location-id}/media/recent
func (api *Api) GetLocationRecentMedia(locationId string, params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get(fmt.Sprintf("/locations/%s/media/recent", locationId), params, res)
	return
}

// Search for a location by geographic coordinate.
// Gets /locations/search
func (api *Api) GetLocationSearch(params url.Values) (res *LocationsResponse, err error) {
	res = new(LocationsResponse)
	err = api.get("/locations/search", params, res)
	return
}
