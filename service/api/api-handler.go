package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// rt.router.GET("/", rt.getHelloWorld)
	// rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Login
	rt.router.POST("/session", rt.wrap(rt.session))

	// Ban
	rt.router.GET("/user/:uid/ban", rt.wrap(rt.getBanList))
	rt.router.PUT("/user/:uid/ban/:banndeduid", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:uid/ban/:banneduid", rt.wrap(rt.unbanUser))

	// Follow
	rt.router.PUT("/user/:uid/follow/:followuid", rt.wrap(rt.followUser))
	rt.router.DELETE("/user/:uid/follow/:followuid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/user/:uid/followers", rt.wrap(rt.getFollowers))
	rt.router.GET("/user/:uid/following", rt.wrap(rt.getFollowing))

	// Photo
	rt.router.POST("/user/:uid/upload", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/user/:uid/photos/:photoid", rt.wrap(rt.deletePhoto))

	// Like
	rt.router.GET("/user/:uid/photos/:photoid/likes", rt.wrap(rt.getPhotoLikes))
	rt.router.PUT("/user/:uid/photos/:photoid/likes/:likeuid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/user/:uid/photos/:photoid/likes/:likeuid", rt.wrap(rt.unlikePhoto))

	// Comment
	rt.router.GET("/user/:uid/photos/:photoid/comments", rt.wrap(rt.getPhotoComments))
	rt.router.POST("/user/:uid/photos/:photoid/comment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/user/:uid/photos/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto))

	// User
	rt.router.GET("/user/:uid", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/user/:uid/username", rt.wrap(rt.setMyUserName))
	
	// Stream
	rt.router.GET("/user/:uid/stream", rt.wrap(rt.getMyStream))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}