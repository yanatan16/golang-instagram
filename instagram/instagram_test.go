package instagram

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
	"time"
)

var DoAuthorizedRequests bool
var api *Api
var ccistulli_id string = "401243155"
var ladygaga_id string = "184692323"

func init() {
	DoAuthorizedRequests = (TestConfig["access_token"] != "")
	if !DoAuthorizedRequests {
		fmt.Println("*** Authorized requests will not performed because no access_token was specified in config_test.go")
	}
	api = createApi()
}

func TestVerifyCredentials(t *testing.T) {
	authorizedRequest(t)

	if ok, err := api.VerifyCredentials(); !ok {
		t.Error(err)
	}
}

func TestUser(t *testing.T) {
	resp, err := api.GetUser(ccistulli_id, nil)
	checkRes(t, resp.Meta, err)

	user := resp.User
	if user.Username != "ccistulli" {
		t.Error("username isn't right", user.Username)
	}
	if user.Id != ccistulli_id {
		t.Error("id isn't right", user.Id)
	}
	if user.Counts == nil {
		t.Error("user doesn't have counts!")
	} else if user.Counts.Media < 28 {
		t.Error("Media count is way off", user.Counts)
	}
}

func TestSelf(t *testing.T) {
	authorizedRequest(t)

	self, err := api.GetSelf()
	checkRes(t, self.Meta, err)

	user, err := api.GetUser(TestConfig["my_id"], nil)
	checkRes(t, user.Meta, err)

	if self.User.Id != TestConfig["my_id"] {
		t.Error("self user Id isn't my_id")
	}
	if !reflect.DeepEqual(self, user) {
		t.Error("Self != user!?", self, user)
	}
}

func TestGetUserFeed(t *testing.T) {
	authorizedRequest(t)

	_, err := api.GetUserFeed(nil)
	if err != nil {
		t.Error(err)
	}

	// Can't really do much here. Don't know who you are.
}

func TestGetUserRecentMedia(t *testing.T) {
	params := url.Values{}
	params.Set("count", "3") // 4 images in this set
	params.Set("max_timestamp", "1385119239")
	params.Set("min_timestamp", "1380215422")
	res, err := api.GetUserRecentMedia(ccistulli_id, params)
	checkRes(t, res.Meta, err)

	if len(res.Medias) != 3 {
		t.Error("Count didn't apply")
	}

	nextRes, err := api.NextMedias(res.Pagination)
	checkRes(t, nextRes.Meta, err)

	if len(nextRes.Medias) != 1 {
		t.Error("Timestamps didn't apply")
	}
	if nextRes.Pagination.Pagination != nil {
		t.Error("Pagination should be not valid!", nextRes.Pagination.Pagination)
	}

	nextNextRes, err := api.NextMedias(nextRes.Pagination)
	if len(nextNextRes.Medias) > 0 {
		t.Error("Pagination returned non-nil next request after nil pagination!")
	}
}

func TestGetUserLikedMedia(t *testing.T) {
	authorizedRequest(t)

	res, err := api.GetUserLikedMedia(nil)
	checkRes(t, res.Meta, err)

	// Can't really do much here. Don't know who you are.
	// We can however go through each image and make sure UserHasLiked is true
	for _, media := range res.Medias {
		if !media.UserHasLiked {
			t.Error("Media from GetUserLikedMedia has UserHasLiked=false ?")
		}
	}
}

func TestGetUserSearch(t *testing.T) {
	term := "traf"
	res, err := api.GetUserSearch(values("q", term, "count", "10")) // If anyone signs up with the name traf, this could fail
	checkRes(t, res.Meta, err)

	if len(res.Users) != 10 {
		t.Error("Users search length not 10? This could mean the search term has an exact match and needs to be changed.")
	}

	for _, user := range res.Users {
		if user.Id == "" {
			t.Error("No ID on a user?")
		} else if user.Username == "" {
			t.Error("no Username on a user?")
		}
	}
}

func TestGetMedia(t *testing.T) {
	res, err := api.GetMedia("594914758412103315_2134762", nil)
	checkRes(t, res.Meta, err)

	if res.Media.Attribution != nil {
		t.Error("Attribution")
	}
	if res.Media.Videos.LowResolution.Url != "http://distilleryimage3.s3.amazonaws.com/0703c326539711e3a8080e78bfefbac0_102.mp4" {
		t.Error("Videos.LowResolution.Url")
	}
	if res.Media.Videos.StandardResolution.Width != int64(640) {
		t.Error("Videos.StandardResolution.Width")
	}
	if len(res.Media.Tags) != 0 {
		t.Error("Tags")
	}
	if res.Media.Type != "video" {
		t.Error("Type")
	}
	if res.Media.Location != nil {
		t.Error("Location")
	}
	if res.Media.Comments.Count < 128 {
		t.Error("Comments.Count")
	}
	if res.Media.Filter != "Normal" {
		t.Error("Filter")
	}
	if tm, err := res.Media.CreatedTime.Time(); err != nil || !tm.Equal(time.Unix(1385139387, 0)) {
		t.Error("CreatedTime", tm, err)
	}
	if res.Media.Link != "http://instagram.com/p/hBj9Ieym6T/" {
		t.Error("Link")
	}
	if res.Media.Likes.Count < 2000 {
		t.Error("Likes.Count")
	}
	if len(res.Media.Likes.Data) == 0 {
		t.Error("Likes.Data")
	}
	if res.Media.Images.Thumbnail.Height != 150 {
		t.Error("Images.Thumbnail.Height")
	}
	if res.Media.Images.StandardResolution.Url != "http://distilleryimage3.s3.amazonaws.com/0703c326539711e3a8080e78bfefbac0_8.jpg" {
		t.Error("Images.StandardResolution.Url")
	}
	if len(res.Media.UsersInPhoto) > 0 {
		t.Error("UsersInPhoto")
	}
	if res.Media.Caption.Text != "Welcome to the anti-stink zone." {
		t.Error("Caption.Text")
	}
	if tm, err := res.Media.Caption.CreatedTime.Time(); err != nil || tm.Unix() != 1385139432 {
		t.Error("Caption.CreatedTime")
	}
	if res.Media.Id != "594914758412103315_2134762" {
		t.Error("Id")
	}
	if res.Media.User.Username != "lululemon" {
		t.Error("User.Username")
	}
}

func TestGetMediaSearch(t *testing.T) {
	res, err := api.GetMediaSearch(values(
		"lat", "48.858844",
		"lng", "2.294351",
		"distance", "1000", // 1km
	))
	checkRes(t, res.Meta, err)

	if len(res.Medias) == 0 {
		t.Error("Paris has to have more than 0 images taken in the last 5 days. Check for a nuclear device.")
	}
}

func TestGetMediaPopular(t *testing.T) {
	res, err := api.GetMediaPopular(values("count", "2"))
	checkRes(t, res.Meta, err)

	if len(res.Medias) != 2 {
		t.Error("Not 2 popular medias?", len(res.Medias))
	}
}

func TestGetMediaSearchError(t *testing.T) {
	res, err := api.GetMediaSearch(nil)
	if err == nil {
		t.Error("Error should have been thrown!", res)
	} else if err.Error() != "Error making api call: Code 400 APIInvalidParametersError missing lat and lng" {
		t.Error("Error isn't right!")
	}
}

func TestGetTag(t *testing.T) {
	res, err := api.GetTag("tbt", nil) // Throw Back Thursday #tbt
	checkRes(t, res.Meta, err)
	if res.Tag.Name != "tbt" {
		t.Error("Tag Name")
	} else if res.Tag.MediaCount < 120000000 {
		t.Error("Tag MediaCount", res.Tag.MediaCount)
	}
}

func TestGetTagRecentMedia(t *testing.T) {
	res, err := api.GetTagRecentMedia("tbt", values("count", "10"))
	checkRes(t, res.Meta, err)

MediaLoop:
	for _, media := range res.Medias {
		for _, tag := range media.Tags {
			if tag == "tbt" {
				continue MediaLoop
			}
		}
		t.Error("No tbt tag on media", media.Id, media.Tags)
	}
}

func TestGetTagSearch(t *testing.T) {
	res, err := api.GetTagSearch(values("q", "toob"))
	checkRes(t, res.Meta, err)
	if len(res.Tags) != 1 {
		t.Error("Should be exact match", len(res.Tags))
	} else if res.Tags[0].Name != "toob" {
		t.Error("Tag name should be exact match to query")
	}
}

func TestGetMediaLikes(t *testing.T) {
	res, err := api.GetMediaLikes("594914758412103315_2134762", nil)
	checkRes(t, res.Meta, err)

	if len(res.Users) < 10 {
		t.Error("too few likers!", len(res.Users))
	}
}

func TestGetMediaComments(t *testing.T) {
	res, err := api.GetMediaComments("594914758412103315_2134762", nil)
	checkRes(t, res.Meta, err)

	if len(res.Comments) < 10 {
		t.Error("too few comments!", len(res.Comments))
	}
}

func TestGetLocation(t *testing.T) {
	res, err := api.GetLocation("3182106", nil)
	checkRes(t, res.Meta, err)

	loc := res.Location
	if ParseLocationId(loc.Id) != "3182106" {
		t.Error("location ID is wrong", loc.Id)
	} else if loc.Name != "La Parisienne" {
		t.Error("location id and name don't match")
	} else if loc.Latitude != 48.8588443 || loc.Longitude != 2.2943506 {
		t.Error("Latitude and longitude are off", loc.Latitude, loc.Longitude)
	}
}

func TestGetLocationRecentMedia(t *testing.T) {
	res, err := api.GetLocationRecentMedia("5450860", nil)
	checkRes(t, res.Meta, err)

	if len(res.Medias) == 0 {
		t.Error("Should be at least one medias in count. We are talking about the Eiffel Tower", len(res.Medias))
	}

	for _, media := range res.Medias {
		if media.Location.Name != "Eiffel Tower" {
			t.Error("Location in media isn't Eiffel Tower")
		}
	}
}

func TestGetLocationSearch(t *testing.T) {
	res, err := api.GetLocationSearch(values(
		"lat", "48.858843",
		"lng", "2.2943506",
		"distance", "1", // 1m
	))
	checkRes(t, res.Meta, err)

	if len(res.Locations) == 0 {
		t.Error("Should be at least 1 location")
	}

	for _, loc := range res.Locations {
		if ParseLocationId(loc.Id) == "3182106" {
			if loc.Name != "La Parisienne" {
				t.Error("location id and name don't match")
			} else if loc.Latitude != 48.8588443 || loc.Longitude != 2.2943506 {
				t.Error("Latitude and longitude are off", loc.Latitude, loc.Longitude)
			}
			return
		}
	}
	t.Error("La Parisienne isn't found!")
}

func TestGetUserFollows(t *testing.T) {
	res, err := api.GetUserFollows(ladygaga_id, nil)
	checkRes(t, res.Meta, err)

	if len(res.Users) > 0 {
		t.Error("ladygaga doesn't follow anyone!", len(res.Users))
	}
}

func TestGetUserFollowsNonTrivial(t *testing.T) {
	res, err := api.GetUserFollows(ccistulli_id, nil)
	checkRes(t, res.Meta, err)

	if len(res.Users) == 0 {
		t.Error("ccistulli should follow people!", len(res.Users))
	}
}

func TestGetUserFollowedBy(t *testing.T) {
	res, err := api.GetUserFollowedBy(ladygaga_id, nil)
	checkRes(t, res.Meta, err)

	if len(res.Users) == 0 {
		t.Error("ladygaga is followed by millions!", len(res.Users))
	}
}

func TestGetUserRequestedBy(t *testing.T) {
	authorizedRequest(t)

	res, err := api.GetUserRequestedBy(nil)
	checkRes(t, res.Meta, err)
	// not much to do here
}

func TestGetUserRelationship(t *testing.T) {
	authorizedRequest(t)

	res, err := api.GetUserRelationship(ladygaga_id, nil)
	checkRes(t, res.Meta, err)

	if res.Relationship.OutgoingStatus == "" {
		t.Error("OutgoingStatus should at least be none", res.Relationship.OutgoingStatus)
	}
	if res.Relationship.IncomingStatus == "" {
		t.Error("IncomingStatus should at least be none", res.Relationship.OutgoingStatus)
	}
}

// -- helpers --

func authorizedRequest(t *testing.T) {
	if !DoAuthorizedRequests {
		t.Skip("Access Token not provided.")
	}
}

func checkRes(t *testing.T, m *Meta, err error) {
	if err != nil {
		t.Error(err)
	}
	if m == nil || m.Code != 200 {
		t.Error("Meta not right", m)
	}
}

func values(keyValues ...string) url.Values {
	v := url.Values{}
	for i := 0; i < len(keyValues)-1; i += 2 {
		v.Set(keyValues[i], keyValues[i+1])
	}
	return v
}

func createApi() *Api {
	return New(TestConfig["client_id"], TestConfig["access_token"])
}
