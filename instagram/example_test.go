package instagram

import (
	"fmt"
	"net/url"
)

// ExampleNew sets up the whole instagram API
func ExampleNew() {
	apiWithoutAuthenticatedUser := New("client_key", "")
	if _, err := apiWithoutAuthenticatedUser.GetMediaPopular(nil); err != nil {
		panic(err)
	}
	fmt.Println("Successfully created instagram.Api without user credentials")

	apiAuthenticatedUser := New("", "access_token")
	if ok, err := apiAuthenticatedUser.VerifyCredentials(); !ok {
		panic(err)
		return
	}
	fmt.Println("Successfully created instagram.Api with user credentials")
}

// ExampleApi_GetUser shows how to get a user object
func ExampleApi_GetUser() {
	api := New("client_id" /* or */, "access_token")

	userResponse, err := api.GetUser("user-id", nil)
	if err != nil {
		panic(err)
	}

	user := userResponse.User
	processUser(user)
}

// ExampleApi_GetUserSearch_Params shows how to use parameters
func ExampleApi_GetUserSearch_Params() {
	api := New("" /* need */, "access_token")

	params := url.Values{}
	params.Set("count", "5") // Get 5 users
	params.Set("q", "jack")  // Search for user "jack"

	usersResponse, err := api.GetUserSearch(params)
	if err != nil {
		panic(err)
	}

	for _, user := range usersResponse.Users {
		processUser(&user)
	}
}

// ExampleApi_GetMediaPopular shows how you can paginate through popular media
func ExampleApi_GetMediaPopular() {
	api := New("client_id" /* or */, "access_token")

	mediasResponse, err := api.GetMediaPopular(nil)
	if err != nil {
		panic(err)
	}

	for _, media := range mediasResponse.Medias {
		processMedia(&media)
	}
}

// ExampleApi_IterateMedia shows how to use iteration on a channel to avoid the complex pagination calls
func ExampleApi_IterateMedia() {
	api := New("client_id" /* or */, "access_token")

	mediasResponse, err := api.GetUserRecentMedia("user-id", nil)
	if err != nil {
		panic(err)
	}

	// Stop 30 days ago
	doneChan := make(chan bool)

	mediaIter, errChan := api.IterateMedia(mediasResponse, doneChan /* optional */)
	for media := range mediaIter {
		processMedia(media)

		if isDone(media) {
			close(doneChan) // Signal to iterator to quit
			break
		}
	}

	// When mediaIter is closed, errChan will either have a single error on it or it will have been closed so this is safe.
	if err := <-errChan; err != nil {
		panic(err)
	}
}

func processMedia(m *Media) {}
func isDone(m *Media) bool {
	return false
}
func processUser(u *User) {}
