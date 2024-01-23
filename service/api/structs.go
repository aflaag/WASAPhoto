package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type Login struct {
	Username string `json:"username"`
}

func LoginDefault() Login {
	return Login{
		Username: "",
	}
}

func LoginFromDatabaseLogin(dbLogin database.DatabaseLogin) Login {
	return Login{
		Username: dbLogin.Username,
	}
}

func (login *Login) LoginIntoDatabaseLogin() database.DatabaseLogin {
	return database.DatabaseLogin{
		Username: login.Username,
	}
}

func LoginFromUsername(username string) Login {
	return Login{
		Username: username,
	}
}

type User struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
}

func UserDefault() User {
	return User{
		Id:       0,
		Username: "",
	}
}

func UserFromDatabaseUser(dbUser database.DatabaseUser) User {
	return User{
		Id:       dbUser.Id,
		Username: dbUser.Username,
	}
}

func (user *User) UserIntoDatabaseUser() database.DatabaseUser {
	return database.DatabaseUser{
		Id:       user.Id,
		Username: user.Username,
	}
}

func UserArrayFromDatabaseUserArray(array []database.DatabaseUser) []User {
	newArray := make([]User, 0)

	for _, element := range array {
		newArray = append(newArray, UserFromDatabaseUser(element))
	}

	return newArray
}

func UserArrayIntoDatabaseUserArray(array []User) []database.DatabaseUser {
	newArray := make([]database.DatabaseUser, 0)

	for _, element := range array {
		newArray = append(newArray, element.UserIntoDatabaseUser())
	}

	return newArray
}

type Photo struct {
	Id           uint32 `json:"id"`
	User         User   `json:"user"`
	Url          string `json:"url"`
	Date         string `json:"date"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	LikeStatus   bool   `json:"like_status"`
}

func PhotoDefault() Photo {
	return Photo{
		Id:           0,
		User:         UserDefault(),
		Url:          "",
		Date:         "",
		LikeCount:    0,
		CommentCount: 0,
		LikeStatus:   false,
	}
}

func PhotoFromDatabasePhoto(dbPhoto database.DatabasePhoto) Photo {
	return Photo{
		Id:           dbPhoto.Id,
		User:         UserFromDatabaseUser(dbPhoto.User),
		Url:          dbPhoto.Url,
		Date:         dbPhoto.Date,
		LikeCount:    dbPhoto.LikeCount,
		CommentCount: dbPhoto.CommentCount,
		LikeStatus:   dbPhoto.LikeStatus,
	}
}

func (photo *Photo) PhotoIntoDatabasePhoto() database.DatabasePhoto {
	return database.DatabasePhoto{
		Id:           photo.Id,
		User:         photo.User.UserIntoDatabaseUser(),
		Url:          photo.Url,
		Date:         photo.Date,
		LikeCount:    photo.LikeCount,
		CommentCount: photo.CommentCount,
		LikeStatus:   photo.LikeStatus,
	}
}

func PhotoArrayFromDatabasePhotoArray(array []database.DatabasePhoto) []Photo {
	newArray := make([]Photo, 0)

	for _, element := range array {
		newArray = append(newArray, PhotoFromDatabasePhoto(element))
	}

	return newArray
}

func PhotoArrayIntoDatabasePhotoArray(array []Photo) []database.DatabasePhoto {
	newArray := make([]database.DatabasePhoto, 0)

	for _, element := range array {
		newArray = append(newArray, element.PhotoIntoDatabasePhoto())
	}

	return newArray
}

type Comment struct {
	Id          uint32 `json:"id"`
	User        User   `json:"user"`
	Photo       Photo  `json:"photo"`
	Date        string `json:"date"`
	CommentBody string `json:"comment_body"`
}

func CommentDefault() Comment {
	return Comment{
		Id:          0,
		User:        UserDefault(),
		Photo:       PhotoDefault(),
		Date:        "",
		CommentBody: "",
	}
}

func CommentFromDatabaseComment(dbComment database.DatabaseComment) Comment {
	return Comment{
		Id:          dbComment.Id,
		User:        UserFromDatabaseUser(dbComment.User),
		Photo:       PhotoFromDatabasePhoto(dbComment.Photo),
		Date:        dbComment.Date,
		CommentBody: dbComment.CommentBody,
	}
}

func (comment *Comment) CommentIntoDatabaseComment() database.DatabaseComment {
	return database.DatabaseComment{
		Id:          comment.Id,
		User:        comment.User.UserIntoDatabaseUser(),
		Photo:       comment.Photo.PhotoIntoDatabasePhoto(),
		Date:        comment.Date,
		CommentBody: comment.CommentBody,
	}
}

func CommentArrayFromDatabaseCommentArray(array []database.DatabaseComment) []Comment {
	newArray := make([]Comment, 0)

	for _, element := range array {
		newArray = append(newArray, CommentFromDatabaseComment(element))
	}

	return newArray
}

func CommentArrayIntoDatabaseCommentArray(array []Comment) []database.DatabaseComment {
	newArray := make([]database.DatabaseComment, 0)

	for _, element := range array {
		newArray = append(newArray, element.CommentIntoDatabaseComment())
	}

	return newArray
}

type Profile struct {
	User           User    `json:"user"`
	Photos         []Photo `json:"photos"`
	PhotoCount     int     `json:"photo_count"`
	FollowersCount int     `json:"followers_count"`
	FollowingCount int     `json:"following_count"`
	FollowStatus   bool    `json:"follow_status"`
	BanStatus      bool    `json:"ban_status"`
}

func ProfileDefault() Profile {
	emptyArray := make([]Photo, 0)

	return Profile{
		User:           UserDefault(),
		Photos:         emptyArray,
		PhotoCount:     0,
		FollowersCount: 0,
		FollowingCount: 0,
		FollowStatus:   false,
		BanStatus:      false,
	}
}

func ProfileFromDatabaseProfile(dbProfile database.DatabaseProfile) Profile {
	return Profile{
		User:           UserFromDatabaseUser(dbProfile.User),
		Photos:         PhotoArrayFromDatabasePhotoArray(dbProfile.Photos),
		PhotoCount:     dbProfile.PhotoCount,
		FollowersCount: dbProfile.PhotoCount,
		FollowingCount: dbProfile.FollowingCount,
		FollowStatus:   dbProfile.FollowStatus,
		BanStatus:      dbProfile.BanStatus,
	}
}

func (profile *Profile) ProfileIntoDatabaseProfile() database.DatabaseProfile {
	return database.DatabaseProfile{
		User:           profile.User.UserIntoDatabaseUser(),
		Photos:         PhotoArrayIntoDatabasePhotoArray(profile.Photos),
		PhotoCount:     profile.PhotoCount,
		FollowersCount: profile.PhotoCount,
		FollowingCount: profile.FollowingCount,
		FollowStatus:   profile.FollowStatus,
		BanStatus:      profile.BanStatus,
	}
}

type Stream struct {
	User   User    `json:"user"`
	Photos []Photo `json:"photos"`
}

func StreamDefault() Stream {
	emptyArray := make([]Photo, 0)

	return Stream{
		User:   UserDefault(),
		Photos: emptyArray,
	}
}

func StreamFromDatabaseStream(dbStream database.DatabaseStream) Stream {
	return Stream{
		User:   UserFromDatabaseUser(dbStream.User),
		Photos: PhotoArrayFromDatabasePhotoArray(dbStream.Photos),
	}
}

func (stream *Stream) CommentIntoDatabaseComment() database.DatabaseStream {
	return database.DatabaseStream{
		User:   stream.User.UserIntoDatabaseUser(),
		Photos: PhotoArrayIntoDatabasePhotoArray(stream.Photos),
	}
}

type UserList struct {
	Users []User `json:"users"`
}

func UserListDefault() UserList {
	emptyArray := make([]User, 0)

	return UserList{
		Users: emptyArray,
	}
}

func UserListFromDatabaseUserList(dbUserList database.DatabaseUserList) UserList {
	return UserList{
		Users: UserArrayFromDatabaseUserArray(dbUserList.Users),
	}
}

func (userList *UserList) UserListIntoDatabaseUserList() database.DatabaseUserList {
	return database.DatabaseUserList{
		Users: UserArrayIntoDatabaseUserArray(userList.Users),
	}
}

type CommentList struct {
	Comments []Comment `json:"comments"`
}

func CommentListDefault() CommentList {
	emptyArray := make([]Comment, 0)

	return CommentList{
		Comments: emptyArray,
	}
}

func CommentListFromDatabaseCommentList(dbCommentList database.DatabaseCommentList) CommentList {
	return CommentList{
		Comments: CommentArrayFromDatabaseCommentArray(dbCommentList.Comments),
	}
}

func (commentList *CommentList) CommentListIntoDatabaseCommentList() database.DatabaseCommentList {
	return database.DatabaseCommentList{
		Comments: CommentArrayIntoDatabaseCommentArray(commentList.Comments),
	}
}
