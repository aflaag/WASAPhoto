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
	var newArray []User

	for idx, element := range array {
		newArray[idx] = UserFromDatabaseUser(element)
	}

	return newArray
}

func UserArrayIntoDatabaseUserArray(array []User) []database.DatabaseUser {
	var newArray []database.DatabaseUser

	for idx, element := range array {
		newArray[idx] = element.UserIntoDatabaseUser()
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
}

func PhotoDefault() Photo {
	return Photo{
		Id:           0,
		User:         UserDefault(),
		Url:          "",
		Date:         "",
		LikeCount:    0,
		CommentCount: 0,
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
	}
}

func PhotoArrayFromDatabasePhotoArray(array []database.DatabasePhoto) []Photo {
	var newArray []Photo

	for idx, element := range array {
		newArray[idx] = PhotoFromDatabasePhoto(element)
	}

	return newArray
}

func PhotoArrayIntoDatabasePhotoArray(array []Photo) []database.DatabasePhoto {
	var newArray []database.DatabasePhoto

	for idx, element := range array {
		newArray[idx] = element.PhotoIntoDatabasePhoto()
	}

	return newArray
}

type Comment struct {
	Id          uint32 `json:"id"`
	CommentBody string `json:"comment_body"`
}

func CommentDefault() Comment {
	return Comment{
		Id:          0,
		CommentBody: "",
	}
}

func CommentFromDatabaseComment(dbComment database.DatabaseComment) Comment {
	return Comment{
		Id:          dbComment.Id,
		CommentBody: dbComment.CommentBody,
	}
}

func (comment *Comment) CommentIntoDatabaseComment() database.DatabaseComment {
	return database.DatabaseComment{
		Id:          comment.Id,
		CommentBody: comment.CommentBody,
	}
}

func CommentArrayFromDatabaseCommentArray(array []database.DatabaseComment) []Comment {
	var newArray []Comment

	for idx, element := range array {
		newArray[idx] = CommentFromDatabaseComment(element)
	}

	return newArray
}

func CommentArrayIntoDatabaseCommentArray(array []Comment) []database.DatabaseComment {
	var newArray []database.DatabaseComment

	for idx, element := range array {
		newArray[idx] = element.CommentIntoDatabaseComment()
	}

	return newArray
}

type Profile struct {
	User           User `json:"user"`
	PhotosCount    int  `json:"photos_count"`
	FollowersCount int  `json:"followers_count"`
	FollowingCount int  `json:"following_count"`
}

func ProfileDefault() Profile {
	return Profile{
		User:           UserDefault(),
		PhotosCount:    0,
		FollowersCount: 0,
		FollowingCount: 0,
	}
}

func ProfileFromDatabaseProfile(dbProfile database.DatabaseProfile) Profile {
	return Profile{
		User:           UserFromDatabaseUser(dbProfile.User),
		PhotosCount:    dbProfile.PhotosCount,
		FollowersCount: dbProfile.PhotosCount,
		FollowingCount: dbProfile.FollowingCount,
	}
}

func (profile *Profile) CommentIntoDatabaseComment() database.DatabaseProfile {
	return database.DatabaseProfile{
		User:           profile.User.UserIntoDatabaseUser(),
		PhotosCount:    profile.PhotosCount,
		FollowersCount: profile.PhotosCount,
		FollowingCount: profile.FollowingCount,
	}
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

func StreamDefault() Stream {
	emptyArray := [0]Photo{}

	return Stream{
		Photos: emptyArray[:],
	}
}

func StreamFromDatabaseStream(dbStream database.DatabaseStream) Stream {
	return Stream{
		Photos: PhotoArrayFromDatabasePhotoArray(dbStream.Photos),
	}
}

func (stream *Stream) CommentIntoDatabaseComment() database.DatabaseStream {
	return database.DatabaseStream{
		Photos: PhotoArrayIntoDatabasePhotoArray(stream.Photos),
	}
}

type UserList struct {
	Users []User `json:"users"`
}

func UserListDefault() UserList {
	emptyArray := [0]User{}

	return UserList{
		Users: emptyArray[:],
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
	emptyArray := [0]Comment{}

	return CommentList{
		Comments: emptyArray[:],
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
