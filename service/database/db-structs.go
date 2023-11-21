package database

type DatabaseUser struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
}

type DatabasePhoto struct {
	Id uint64 `json:"id"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCount uint64 `json:"like_count"`
	CommentCount uint64 `json:"comment_count"`
}

type DatabaseComment struct {
	Id uint64 `json:"id"`
	CommentBody string `json:"comment_body"`
}

type DatabaseProfile struct {
	User DatabaseUser `json:"user"`
	PhotosCount uint64 `json:"photos_count"`
	FollowersCount uint64 `json:"followers_count"`
	FollowingCount uint64 `json:"following_count"`
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