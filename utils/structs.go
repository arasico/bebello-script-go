package utils

import (
	"net/http"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

//
//
//type User struct {
//	Id        int    `json:"id"`
//	FirstName string `json:"firstName"`
//	LastName  string `json:"lastName"`
//	Photo     string `json:"photo"`
//	Gender    int    `json:"gender"`
//	Username  string `json:"username"`
//}
//type UserMessage struct {
//	UserId   int
//	DeviceId string
//	Message  []byte
//}
//
//type JwtClaims struct {
//	UserId float64 `json:"user_id"`
//	Agent  string  `json:"agent"`
//	App    string  `json:"app"`
//	jwt.StandardClaims
//}
//type Conversation struct {
//	Thread     string         `json:"thread"`
//	Type       string         `json:"type"`
//	OwnerId    int            `json:"owner_id"`
//	TargetId   sql.NullInt64  `json:"target_user_id"`
//	LastUpdate int64          `json:"last_update"`
//	Unread     int32          `json:"unread_count"`
//	LastMsg    sql.NullString `json:"last_msg"`
//	Group      struct {
//		Id          int64  `json:"id"`
//		ThreadG     string `json:"thread"`
//		Name        string `json:"name"`
//		IsFree      string `json:"is_free"`
//		Username    string `json:"username"`
//		UserId      int    `json:"user_id"`
//		AccessLevel string `json:"access_level"`
//		Status      string `json:"status"`
//		ExpireAt    string `json:"expire_at"`
//		Media       string `json:"media"`
//	}
//	Channel struct {
//		Id          int64  `json:"id"`
//		ThreadG     string `json:"thread"`
//		Name        string `json:"name"`
//		IsFree      string `json:"is_free"`
//		IsPublic    string `json:"is_public"`
//		Username    string `json:"username"`
//		UserId      int    `json:"user_id"`
//		AccessLevel string `json:"access_level"`
//		Status      string `json:"status"`
//		ExpireAt    string `json:"expire_at"`
//		Media       string `json:"media"`
//	}
//}
//
//type Conversations []Conversation
