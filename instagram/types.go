package instagram

// Instagram User Object. Note that user objects are not always fully returned.
// Be sure to see the descriptions on the instagram documentation for any given endpoint.
type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string
	Website        string
	Counts         *UserCounts
}

// Instagram User Counts object. Returned on User objects
type UserCounts struct {
	Media      int64
	Follows    int64
	FollowedBy int64 `json:"followed_by"`
}

// Instagram Media object
type Media struct {
	Type         string
	Id           string
	UsersInPhoto []UserPosition `json:"users_in_photo"`
	Filter       string
	Tags         []string
	Comments     *Comments
	Caption      *Caption
	Likes        *Likes
	Link         string
	User         *User
	CreatedTime  StringUnixTime `json:"created_time"`
	Images       *Images
	Videos       *Videos
	Location     *Location
	UserHasLiked bool `json:"user_has_liked"`
	Attribution  *Attribution
}

// A pair of user object and position
type UserPosition struct {
	User     *User
	Position *Position
}

// A position in a media
type Position struct {
	X float64
	Y float64
}

// Instagram tag
type Tag struct {
	MediaCount int64 `json:"media_count"`
	Name       string
}

type Comments struct {
	Count int64
	Data  []Comment
}

type Comment struct {
	CreatedTime StringUnixTime `json:"created_time"`
	Text        string
	From        *User
	Id          string
}

type Caption Comment

type Likes struct {
	Count int64
	Data  []User
}

type Images struct {
	LowResolution      *Image `json:"low_resolution"`
	Thumbnail          *Image
	StandardResolution *Image `json:"standard_resolution"`
}

type Image struct {
	Url    string
	Width  int64
	Height int64
}

type Videos struct {
	LowResolution      *Video `json:"low_resolution"`
	StandardResolution *Video `json:"standard_resolution"`
}

type Video Image

type Location struct {
	Id        LocationId
	Name      string
	Latitude  float64
	Longitude float64
}

type Relationship struct {
	IncomingStatus string `json:"incoming_status"`
	OutgoingStatus string `json:"outgoing_status"`
}

// If another app uploaded the media, then this is the place it is given. As of 11/2013, Hipstamic is the only allowed app
type Attribution struct {
	Website   string
	ItunesUrl string
	Name      string
}
