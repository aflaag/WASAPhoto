package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// type Default[T any] interface {
// 	Default() 
// }

type User struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
}

func UserDefault() User {
	return User {
		Id: 0,
		Username: "",
	}
}

func UserFromDatabaseUser(dbUser database.DatabaseUser) User {
	return User {
		Id: dbUser.Id,
		Username: dbUser.Username,
	}
}

func (user *User) UserIntoDatabaseUser() database.DatabaseUser {
	return database.DatabaseUser {
		Id: user.Id,
		Username: user.Username,
	}
}

type Photo struct {
	Id uint64 `json:"id"`
	Url string `json:"url"`
	Date string `json:"date"`
	LikeCount int `json:"like_count"`
	CommentCount int `json:"comment_count"`
}

func PhotoDefault() Photo {
	return Photo {
		Id: 0,
		Url: "",
		Date: "",
		LikeCount: 0,
		CommentCount: 0,
	}
}

func PhotoFromDatabasePhoto(dbPhoto database.DatabasePhoto) Photo {
	return Photo {
		Id: dbPhoto.Id,
		Url: dbPhoto.Url,
		Date: dbPhoto.Date,
		LikeCount: dbPhoto.LikeCount,
		CommentCount: dbPhoto.CommentCount,
	}
}

func (photo *Photo) PhotoIntoDatabasePhoto() database.DatabasePhoto {
	return database.DatabasePhoto {
		Id: photo.Id,
		Url: photo.Url,
		Date: photo.Date,
		LikeCount: photo.LikeCount,
		CommentCount: photo.CommentCount,
	}
}

type Comment struct {
	Id uint64 `json:"id"`
	CommentBody string `json:"comment_body"`
}

func CommentDefault() Comment {
	return Comment {
		Id: 0,
		CommentBody: "",
	}
}

func CommentFromDatabaseComment(dbComment database.DatabaseComment) Comment {
	return Comment {
		Id: dbComment.Id,
		CommentBody: dbComment.CommentBody,
	}
}

func (comment *Comment) CommentIntoDatabaseComment() database.DatabaseComment {
	return database.DatabaseComment {
		Id: comment.Id,
		CommentBody: comment.CommentBody,
	}
}

type Profile struct {
	User User `json:"user"`
	PhotosCount int `json:"photos_count"`
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
}

func ProfileDefault() Profile {
	return Profile {
		User: UserDefault(),
		PhotosCount: 0,
		FollowersCount: 0,
		FollowingCount: 0,
	}
}

func ProfileFromDatabaseProfile(dbProfile database.DatabaseProfile) Profile {
	return Profile {
		User: UserFromDatabaseUser(dbProfile.User),
		PhotosCount: dbProfile.PhotosCount,
		FollowersCount: dbProfile.PhotosCount,
		FollowingCount: dbProfile.FollowingCount,
	}
}

func (profile *Profile) CommentIntoDatabaseComment() database.DatabaseProfile {
	return database.DatabaseProfile {
		User: profile.User.UserIntoDatabaseUser(),
		PhotosCount: profile.PhotosCount,
		FollowersCount: profile.PhotosCount,
		FollowingCount: profile.FollowingCount,
	}
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

func StreamDefault() Stream {
	emptyArray := [0]Photo{}

	return Stream {
		Photos: emptyArray[:],
	}
}

// TODO: INVECE DI FARLA 3 VOLTE FAI LE INTERFACCE E FALLA GENERICS
func PhotoArrayFromDatabasePhotoArray(array []database.DatabasePhoto) []Photo {
	var databaseArray []Photo

	for idx, element := range array {
		databaseArray[idx] = PhotoFromDatabasePhoto(element)
	}

	return databaseArray
}

func StreamFromDatabaseStream(dbStream database.DatabaseStream) Stream {
	return Stream {
		Photos: PhotoArrayFromDatabasePhotoArray(dbStream.Photos),
	}
}

func (stream *Stream) CommentIntoDatabaseComment() database.DatabaseStream {
	return database.DatabaseStream {
		Photos: stream.Photos,
	}
}

type UserList struct {
	Users []User `json:"users"`
}

func UserListFromDatabaseUserList(dbUserList database.DatabaseUserList) UserList {
	return UserList {
		Users: dbUserList.Users,
	}
}

func (userList *UserList) UserListIntoDatabaseUserList() database.DatabaseUserList {
	return database.DatabaseUserList {
		Users: userList.Users,
	}
}

type CommentList struct {
	Comments []Comment `json:"comments"`
}

func CommentListFromDatabaseCommentList(dbCommentList database.DatabaseCommentList) CommentList {
	return CommentList {
		Comments: dbCommentList.Comments,
	}
}

func (commentList *CommentList) CommentListIntoDatabaseCommentList() database.DatabaseCommentList {
	return database.DatabaseCommentList {
		Comments: commentList.Comments,
	}
}