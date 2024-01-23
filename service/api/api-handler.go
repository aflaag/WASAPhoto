package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Login
	rt.router.POST("/session", rt.wrap(rt.session)) // DONE DONE

	// Ban
	rt.router.PUT("/user/:uname/ban/:banned_uname", rt.wrap(rt.banUser))      // DONE
	rt.router.DELETE("/user/:uname/ban/:banned_uname", rt.wrap(rt.unbanUser)) // DONE

	// Follow
	rt.router.PUT("/user/:uname/follow/:followed_uname", rt.wrap(rt.followUser))      // DONE
	rt.router.DELETE("/user/:uname/follow/:followed_uname", rt.wrap(rt.unfollowUser)) // DONE
	rt.router.GET("/user/:uname/followers", rt.wrap(rt.getFollowers))                 // DONE
	rt.router.GET("/user/:uname/following", rt.wrap(rt.getFollowing))                 // DONE

	// Photo
	rt.router.POST("/user/:uname/upload", rt.wrap(rt.uploadPhoto))             // DONE
	rt.router.DELETE("/user/:uname/photos/:photo_id", rt.wrap(rt.deletePhoto)) // DONE

	// Like
	rt.router.GET("/user/:uname/photos/:photo_id/likes", rt.wrap(rt.getPhotoLikes))              // DONE
	rt.router.PUT("/user/:uname/photos/:photo_id/likes/:like_uname", rt.wrap(rt.likePhoto))      // DONE
	rt.router.DELETE("/user/:uname/photos/:photo_id/likes/:like_uname", rt.wrap(rt.unlikePhoto)) // DONE

	// Comment
	rt.router.GET("/user/:uname/photos/:photo_id/comments", rt.wrap(rt.getPhotoComments))              // DONE
	rt.router.POST("/user/:uname/photos/:photo_id/comment", rt.wrap(rt.commentPhoto))                  // DONE
	rt.router.DELETE("/user/:uname/photos/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto)) // DONE

	// User
	rt.router.GET("/user/:uname", rt.wrap(rt.getUserProfile))            // DONE
	rt.router.PUT("/user/:uname/setusername", rt.wrap(rt.setMyUserName)) // DONE
	rt.router.GET("/user/:uname/users", rt.wrap(rt.getUsers))            // DONE

	// Stream
	rt.router.GET("/user/:uname/stream", rt.wrap(rt.getMyStream)) // DONE

	// Liveness
	rt.router.GET("/liveness", rt.liveness) // TODO:

	return rt.router
}
