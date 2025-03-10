package dao

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id             int64  `json:"id,omitempty"`
	UserId         int64  `json:"user_id,omitempty"`
	Author         User   `gorm:"foreignKey:UserId"`
	PlayUrl        string `json:"play_url,omitempty"`
	CoverUrl       string `json:"cover_url,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
	CommentCount   int64  `json:"comment_count,omitempty"`
	IsFavorite     bool   `json:"is_favorite,omitempty"`
	SubmissionTime string `json:"submission_time,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Password        string `json:"password,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	avatar          string `json:"avatar,omitempty"` //用户头像
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  string `json:"totalFavorited,omitempty"`
	WorkCount       int64  `json:"workCount,omitempty"`
	FavoriteCount   int64  `json:"favorite_Count,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type Follow struct {
	UserId     int64 `json:"user_id,omitempty"`
	FollowerId int64 `json:"follow_id,omitempty"`
}

type Friend struct {
	UserId     int64 `json:"userId,omitempty"`
	FriendID   int64 `json:"friendID,omitempty"`
	FriendName int64 `json:"friend_name,omitempty"`
}
