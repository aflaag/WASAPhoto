package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type User struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
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
	LikeCount uint64 `json:"like_count"`
	CommentCount uint64 `json:"comment_count"`
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

func CommentFromDatabaseComment(dbComment database.DatabaseComment) Comment {
	return Photo {
		Id: dbPhoto.Id,
		CommentBody: dbPhoto.CommentBody,
	}
}

func (comment *Comment) CommentIntoDatabaseComment() database.DatabaseComent {
	return database.DatabaseComment {
		Id: comment.Id,
		CommentBody: comment.CommentBody
	}
}

type Profile struct {
	User User `json:"user"`
	PhotosCount uint64 `json:"photos_count"`
	FollowersCount uint64 `json:"followers_count"`
	FollowingCount uint64 `json:"following_count"`
}

func ProfileFromDatabaseProfile(dbProfile database.DatabaseProfile) Profile {
	return Profile {
		User: dbProfile.User,
		PhotosCount: dbProfile.PhotosCount,
		FollowersCount: dbProfile.PhotosCount,
		FollowingCount: dbProfile.FollowingCount,
	}
}

func (profile *Profile) CommentIntoDatabaseComment() database.DatabaseProfile {
	return database.DatabaseProfile {
		User: profile.User,
		PhotosCount: profile.PhotosCount,
		FollowersCount: profile.PhotosCount,
		FollowingCount: profile.FollowingCount,
	}
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

func StreamFromDatabaseStream(dbStream database.DatabaseStream) Stream {
	return Stream {
		Photos: dbStream.Photos,
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
		Users: dbCommentList.users,
	}
}

func (commentList *CommentList) CommentListIntoDatabaseCommentList() database.DatabaseCommentList {
	return database.DatabaseCommmentList {
		Users: commentList.Users,
	}
}