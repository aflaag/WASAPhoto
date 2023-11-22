package database

type DatabaseUser struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
}

type DatabasePhoto struct {
	Id uint64 `json:"id"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCount int `json:"like_count"`
	CommentCount int `json:"comment_count"`
}

type DatabaseComment struct {
	Id uint64 `json:"id"`
	CommentBody string `json:"comment_body"`
}

type DatabaseProfile struct {
	User DatabaseUser `json:"user"`
	PhotosCount int `json:"photos_count"`
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
}

type DatabaseStream struct {
	Photos []DatabasePhoto `json:"photos"`
}

type DatabaseUserList struct {
	Users []DatabaseUser `json:"users"`
}

type DatabaseCommentList struct {
	Comments []DatabaseComment `json:"comments"`
}