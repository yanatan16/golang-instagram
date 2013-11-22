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

anotherAuthenticatedApi := instagram.New("", "my-access-token")
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
    fmt.Printf("My userid is %s and I have %d followers\n", me.Id, me.Counts.Followers)
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

## Notes

- Certain methods require an access token so check the official documentation before using an unauthenticated `Api`. Also, there is a 5000 request per hour rate limit on any one ClientId or AccessToken, so it is advisable to use AccessTokens when available. This package will use it if it is given over a ClientId.
- Location.Id is sometimes returned as an integer (in media i think) and sometimes a string. Because of this, we have to call it an `interface{}`. But there is a facility to force it to a string, as follows:

```go
var loc Location
stringIdVersion := instagram.ParseLocationId(loc.Id)
```

If anyone can prove to me that they fixed this bug, just let me know and we can change it to a string (all other IDs are strings...)

## License

MIT-style. See LICENSE file.