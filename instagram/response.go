package instagram

type UserResponse struct {
	MetaResponse
	User *User `json:"data"`
}

type UsersResponse struct {
	MetaResponse
	Users []User `json:"data"`
}

type PaginatedUsersResponse struct {
	UsersResponse
	Pagination *UserPagination
}

type MediaResponse struct {
	MetaResponse
	Media *Media `json:"data"`
}

type MediasResponse struct {
	MetaResponse
	Medias []Media `json:"data"`
}

type PaginatedMediasResponse struct {
	MediasResponse
	Pagination *MediaPagination
}

type CommentsResponse struct {
	MetaResponse
	Comments []Comment `json:"data"`
}

type TagResponse struct {
	MetaResponse
	Tag *Tag `json:"data"`
}

type TagsResponse struct {
	MetaResponse
	Tags []Tag `json:"data"`
}

type LocationResponse struct {
	MetaResponse
	Location *Location `json:"data"`
}

type LocationsResponse struct {
	MetaResponse
	Locations []Location `json:"data"`
}

type RelationshipResponse struct {
	MetaResponse
	Relationship *Relationship `json:"data"`
}

type MetaResponse struct {
	Meta *Meta
}

type Pagination struct {
	NextUrl   string `json:"next_url"`
	NextMaxId string `json:"next_max_id"`

	// Used only on GetTagRecentMedia()
	NextMaxTagId string `json:"next_max_tag_id"`
	// Used only on GetTagRecentMedia()
	MinTagId string `json:"min_tag_id"`
}

type Meta struct {
	Code         int
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}

// MediaPagination will give you an easy way to request the next page of media.
type MediaPagination struct {
	*Pagination
}

// UserPagination will give you an easy way to request the next page of media.
type UserPagination struct {
	*Pagination
}
