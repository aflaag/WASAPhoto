package api

type User struct {
	Username string `json:"username"`
}

type Photo struct {
	Id uint64 `json:"id"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCount uint64 `json:"like_count"`
	CommentCount uint64 `json:"comment_count"`
}

type Comment struct {
	Id uint64 `json:"id"`
	CommentBody string `json:"comment_body"`
}

type Profile struct {
	User User `json:"user"`
	PhotosCount uint64 `json:"photos_count"`
	FollowersCount uint64 `json:"followers_count"`
	FollowingCount uint64 `json:"following_count"`
}