package instagram

import (
	"fmt"
	"net/url"
)

// ExampleNew sets up the whole instagram API
func ExampleNew() {
	apiAuthenticatedUser := New("client_key", "secret", "", true)
	if ok, err := apiAuthenticatedUser.VerifyCredentials(); !ok {
		panic(err)
	}
	fmt.Println("Successfully created instagram.Api with user credentials")
}

// ExampleApi_GetUser shows how to get a user object
func ExampleApi_GetUser() {
	// *** or ***
	api := New("client_id", "client_secret", "access_token", true)

	userResponse, err := api.GetUser("user-id", nil)
	if err != nil {
		panic(err)
	}

	user := userResponse.User
	processUser(user)
}

// ExampleApi_GetUserSearch_Params shows how to use parameters
func ExampleApi_GetUserSearch_Params() {
	// *** need ***
	api := New("", "client_secret", "access_token", true)

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

// ExampleApi_GetUserRecentMedia : Get the most recent media published by the owner
func ExampleApi_GetUserRecentMedia() {
	// *** or ***
	api := New("client_id", "client_secret", "access_token", true)

	params := url.Values{}
	params.Set("count", "3") // 4 images in this set
	params.Set("max_timestamp", "1466809870")
	params.Set("min_timestamp", "1396751898")
	mediasResponse, err := api.GetUserRecentMedia(ccistulli_id, params)

	if err != nil {
		panic(err)
	}

	for _, media := range mediasResponse.Medias {
		processMedia(&media)
	}
}

// ExampleApi_IterateMedia shows how to use iteration on a channel to avoid the complex pagination calls
func ExampleApi_IterateMedia() {
	// *** or ***
	api := New("client_id", "client_secret", "access_token", true)

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
