package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// login
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))

	// search
	rt.router.GET("/users", rt.wrap(rt.getUsersQuery))

	// user
	rt.router.PUT("/users/:id", rt.wrap(rt.putNickname))
	rt.router.GET("/users/:id", rt.wrap(rt.getUserProfile))

	// ban
	rt.router.PUT("/users/:id/banned_users/:banned_id", rt.wrap(rt.putBan))
	rt.router.DELETE("/users/:id/banned_users/:banned_id", rt.wrap(rt.deleteBan))

	// followers
	rt.router.PUT("/users/:id/followers/:follower_id", rt.wrap(rt.putFollow))
	rt.router.DELETE("/users/:id/followers/:follower_id", rt.wrap(rt.deleteFollow))

	// stream
	rt.router.GET("/users/:id/home", rt.wrap(rt.getHome))

	// photo
	rt.router.POST("/users/:id/photos", rt.wrap(rt.postPhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:id/photos/:photo_id", rt.wrap(rt.getPhoto))

	// comments
	rt.router.POST("/users/:id/photos/:photo_id/comments", rt.wrap(rt.postComment))
	rt.router.DELETE("/users/:id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.deleteComment))

	// likes
	rt.router.PUT("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.putLike))
	rt.router.DELETE("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.deleteLike))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
