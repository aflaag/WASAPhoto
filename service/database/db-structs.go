package database

type DatabaseLogin struct {
	Username string `json:"username"`
}

func DatabaseLoginDefault() DatabaseLogin {
	return DatabaseLogin{
		Username: "",
	}
}

type DatabaseUser struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
}

func DatabaseUserDefault() DatabaseUser {
	return DatabaseUser{
		Id:       0,
		Username: "",
	}
}

type DatabasePhoto struct {
	Id           uint32       `json:"id"`
	User         DatabaseUser `json:"user"`
	Url          string       `json:"url"`
	Date         string       `json:"date"`
	LikeCount    int          `json:"like_count"`
	CommentCount int          `json:"comment_count"`
}

func DatabasePhotoDefault() DatabasePhoto {
	return DatabasePhoto{
		Id:           0,
		User:         DatabaseUserDefault(),
		Url:          "",
		Date:         "",
		LikeCount:    0,
		CommentCount: 0,
	}
}

type DatabaseComment struct {
	Id          uint32       `json:"id"`
	User        DatabaseUser `json:"user"`
	CommentBody string       `json:"comment_body"`
}

func DatabaseCommentDefault() DatabaseComment {
	return DatabaseComment{
		Id:          0,
		User:        DatabaseUserDefault(),
		CommentBody: "",
	}
}

type DatabaseProfile struct {
	User           DatabaseUser `json:"user"`
	PhotosCount    int          `json:"photos_count"`
	FollowersCount int          `json:"followers_count"`
	FollowingCount int          `json:"following_count"`
}

func DatabaseProfileDefault() DatabaseProfile {
	return DatabaseProfile{
		User:           DatabaseUserDefault(),
		PhotosCount:    0,
		FollowersCount: 0,
		FollowingCount: 0,
	}
}

type DatabaseStream struct {
	Photos []DatabasePhoto `json:"photos"`
}

func DatabaseStreamDefault() DatabaseStream {
	emptyArray := [0]DatabasePhoto{}

	return DatabaseStream{
		Photos: emptyArray[:],
	}
}

type DatabaseUserList struct {
	Users []DatabaseUser `json:"users"`
}

func DatabaseUserListDefault() DatabaseUserList {
	emptyArray := [0]DatabaseUser{}

	return DatabaseUserList{
		Users: emptyArray[:],
	}
}

type DatabaseCommentList struct {
	Comments []DatabaseComment `json:"comments"`
}

func DatabaseCommentListDefault() DatabaseCommentList {
	emptyArray := [0]DatabaseComment{}

	return DatabaseCommentList{
		Comments: emptyArray[:],
	}
}
