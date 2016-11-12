package instagram

import (
	"net/url"
	"testing"
)

func TestIterate_GetUserFollowedBy(t *testing.T) {
	res, err := api.GetUserFollowedBy(values("count", "5"))
	checkRes(t, res.Meta, err)

	doneChan := make(chan bool) // This is only needed if you want to close early
	defer close(doneChan)

	userChan, errChan := api.IterateUsers(res, doneChan)

	i := 0
	for user := range userChan {
		if user.Id == "" {
			t.Error("user has empty id", user)
		}
		i++
		if i > 19 {
			// breaking early
			doneChan <- true
			// userChan should close immediately afterward, exiting the loop, and not closing the channel again
		}
	}

	// should be closed
	if u, ok := <-userChan; ok {
		t.Error("User Channel shouldn't have any more data on it. It should close!", u)
	}

	if err := <-errChan; err != nil {
		t.Error(err)
	}
}

func TestIterate_GetUserRecentMedia(t *testing.T) {
	params := url.Values{}
	params.Set("count", "2") // 5 images in this set. Get them 2 at time
	params.Set("max_timestamp", "1384161094")
	params.Set("min_timestamp", "1382656250")
	res, err := api.GetUserRecentMedia(ladygaga_id, params)
	checkRes(t, res.Meta, err)

	mediaChan, errChan := api.IterateMedia(res, nil)
	for media := range mediaChan {
		if media.User.Username != "ladygaga" {
			t.Error("Got a media with wrong username?", media.User)
		}
	}
	if err := <-errChan; err != nil {
		t.Error(err)
	}
}
