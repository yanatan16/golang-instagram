# golang-instagram

A [instagram](http://instagram.com) [API](http://instagram.com/developer) wrapper.

## Features

Implemented:

- All GET requests are implemented
- Both authenticated and unauthenticated requests can be made
- A nice iterator facility for paginated requests
- No `interface{}` data types! (1 exception, see Location.Id note below)

Todo:

- Authentication
- POST / DELETE requests (you need special permissions for these so no way to test.)

## Documentation

[Documentation on godoc.org](http://godoc.org/github.com/yanatan16/golang-instagram/instagram)

## Install

```
go get github.com/yanatan16/golang-instagram/instagram
```

## Creation

```go
import (
  "github.com/yanatan16/golang-instagram/instagram"
)

unauthenticatedApi := &instagram.Api{
  ClientId: "my-client-id",
}

authenticatedApi := &instagram.Api{
  AccessToken: "my-access-token",
}

- if enforceSigned false
anotherAuthenticatedApi := instagram.New("", "", "my-access-token", false)

- if enforceSigned true
anotherAuthenticatedApi := instagram.New("client_id", "client_secret", "my-access-token", false)

```

## Usage

See the [documentation](http://godoc.org/github.com/yanatan16/golang-instagram/instagram), [endpoint examples](https://github.com/yanatan16/golang-instagram/blob/master/instagram/example_test.go), and the [iteration tests](https://github.com/yanatan16/golang-instagram/blob/master/instagram/iterate_test.go) for a deeper dive than whats below.

```go
import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"
)

func DoSomeInstagramApiStuff(accessToken string) {
  api := New("", accessToken)

  if ok, err := api.VerifyCredentials(); !ok {
    panic(err)
  }

  var myId string

  // Get yourself!
  if resp, err := api.GetSelf(); err != nil {
    panic(err)
  } else {
    // A response has two fields: Meta which you shouldn't really care about
    // And whatever your getting, in this case, a User
    me := resp.User
    fmt.Printf("My userid is %s and I have %d followers\n", me.Id, me.Counts.FollowedBy)
  }

  params := url.Values{}
  params.Set("count", "1")
  if resp, err := api.GetUserRecentMedia("self" /* this works :) */, params); err != nil {
    panic(err)
  } else {
    if len(resp.Medias) == 0 { // [sic]
      panic("I should have some sort of media posted on instagram!")
    }
    media := resp.Medias[0]
    fmt.Println("My last media was a %s with %d comments and %d likes. (url: %s)\n", media.Type, media.Comments.Count, media.Like.Count, media.Link)
  }
}
```

There's many more endpoints and a fancy iteration wrapper. Check it out in the code and documentation!

## Iteration

So pagination makes iterating through a list of users or media possible, but its not easy. So, because Go has nice iteration facilities (i.e. `range`), this package includes two useful methods for iterating over paginating: `api.IterateMedias` and `api.IterateUsers`. You can see [the tests](https://github.com/yanatan16/golang-instagram/blob/master/instagram/iterate_test.go) and [the docs](http://godoc.org/github.com/yanatan16/golang-instagram/instagram/#Api.IterateMedia) for more info.

```go
// First go and make the original request, passing in any additional parameters you need
res, err := api.GetUserRecentMedia("some-user", params)
if err != nil {
  panic(err)
}

// If you plan to break early, create a done channel. Pass in nil if you plan to exhaust the pagination
done := make(chan bool)
defer close(done)

// Here we get back two channels. Don't worry about the error channel for now
medias, errs := api.IterateMedia(res, done)

for media := range medias {
  processMedia(media)

  if doneWithMedia(media) {
    // This is how we signal to the iterator to quit early
    done <- true
  }
}

// If we exited early due to an error, we can check here
if err, ok := <- errs; ok && err != nil {
  panic(err)
}
```

## Tests

To run the tests, you'll need at least a `ClientId` (which you can get from [here](http://instagram.com/developer/clients/manage/)), and preferably an authenticated users' `AccessToken`, which you can get from making a request on the [API Console](http://instagram.com/developer/api-console/)

First, fill in `config_test.go.example` and save it as `config_test.go`. Then run `go test`

## Notes

- Certain methods require an access token so check the official documentation before using an unauthenticated `Api`. Also, there is a 5000 request per hour rate limit on any one ClientId or AccessToken, so it is advisable to use AccessTokens when available. This package will use it if it is given over a ClientId.
- Location.Id is sometimes returned as an integer (in media i think) and sometimes a string. Because of this, we have to call it an `interface{}`. But there is a facility to force it to a string, as follows:

```go
var loc Location
stringIdVersion := instagram.ParseLocationId(loc.Id)
```

If anyone can prove to me that they fixed this bug, just let me know and we can change it to a string (all other IDs are strings...)

- `created_time` fields come back as strings. So theres a handy type `StringUnixTimeStringUnixTime` which has a nice method `func (sut StringUnixTime) Time() (time.Time, error)` that you can use to cast it to a golang time.
- I apologize for using Medias [sic] everywhere, I needed a plural version that isn't spelled the same.

## License

MIT-style. See LICENSE file.
