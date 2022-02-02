package objects

import "time"

type VideoFrame struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type VideoImage struct {
	VideoFrame
	WithPadding int `json:"with_padding"`
}

type VideoLikes struct {
	Count     int64 `json:"count"`
	UserLikes int   `json:"userLikes"`
}

type VideoReposts struct {
	Count        int64 `json:"count"`
	WallCount    int64 `json:"wall_count"`
	MailCount    int64 `json:"mail_count"`
	UserReposted int   `json:"user_reposted"`
}

type Video struct {
	ID            int64         `json:"id"`
	OwnerID       int64         `json:"owner_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Duration      time.Duration `json:"duration"`
	Image         []VideoImage  `json:"image"`
	FirstFrame    []VideoFrame  `json:"first_frame"`
	Date          int64         `json:"date"`
	AddingDate    int64         `json:"adding_date"`
	Views         int64         `json:"views"`
	LocalViews    int64         `json:"local_views"`
	Comments      int64         `json:"comments"`
	Player        string        `json:"player"`
	Platform      string        `json:"platform"`
	CanAdd        int           `json:"canAdd"`
	IsPrivate     int           `json:"is_private"`
	AccessKey     string        `json:"access_key"`
	Processing    int           `json:"processing"`
	IsFavorite    bool          `json:"is_favorite"`
	CanComment    int           `json:"can_comment"`
	CanEdit       int           `json:"can_edit"`
	CanLike       int           `json:"can_like"`
	CanRepost     int           `json:"can_repost"`
	CanSubscribe  int           `json:"can_subscribe"`
	CanAddToFaves int           `json:"can_add_to_faves"`
	CanAttachLink int           `json:"can_attach_link"`
	Width         int           `json:"width"`
	Height        int           `json:"height"`
	UserID        int64         `json:"user_id,omitempty"`
	Converting    int           `json:"converting"`
	Added         int           `json:"added"`
	IsSubscribed  int           `json:"is_subscribed"`
	Repeat        int           `json:"repeat,omitempty"`
	Type          string        `json:"type"`
	Balance       int64         `json:"balance"`
	LiveStatus    string        `json:"live_status"`
	Live          int           `json:"live"`
	Upcoming      int           `json:"upcoming,omitempty"`
	Spectators    int64         `json:"spectators"`
	Likes         VideoLikes    `json:"likes"`
	Reposts       VideoReposts  `json:"reposts"`
}
