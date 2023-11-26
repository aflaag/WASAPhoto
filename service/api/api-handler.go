package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Login
	rt.router.POST("/session", rt.wrap(rt.session)) // works

	// Ban
	rt.router.GET("/user/:uname/ban", rt.wrap(rt.getBanList))
	rt.router.PUT("/user/:uname/ban/:bannded_uname", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:uname/ban/:banned_uname", rt.wrap(rt.unbanUser))

	// Follow
	rt.router.PUT("/user/:uname/follow/:follow_uname", rt.wrap(rt.followUser)) // TODO: controlla il token
	rt.router.DELETE("/user/:uname/follow/:follow_uname", rt.wrap(rt.unfollowUser)) // TODO: controlla il token
	rt.router.GET("/user/:uname/followers", rt.wrap(rt.getFollowers))
	rt.router.GET("/user/:uname/following", rt.wrap(rt.getFollowing))

	// Photo
	rt.router.POST("/user/:uname/upload", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/user/:uname/photos/:photo_id", rt.wrap(rt.deletePhoto))

	// Like
	rt.router.GET("/user/:uname/photos/:photo_id/likes", rt.wrap(rt.getPhotoLikes))
	rt.router.PUT("/user/:uname/photos/:photo_id/likes/:like_uname", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/user/:uname/photos/:photo_id/likes/:like_uname", rt.wrap(rt.unlikePhoto))

	// Comment
	rt.router.GET("/user/:uname/photos/:photo_id/comments", rt.wrap(rt.getPhotoComments))
	rt.router.POST("/user/:uname/photos/:photo_id/comment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/user/:uname/photos/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

	// User
	rt.router.GET("/user/:uname", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/user/:uname/username", rt.wrap(rt.setMyUserName))
	
	// Stream
	rt.router.GET("/user/:uname/stream", rt.wrap(rt.getMyStream))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}