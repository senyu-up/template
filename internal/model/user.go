package model

func init() {
	Register(&User{})
}

// User 用户信息
type User struct {
	Id             int32  `json:"id" gorm:"primary_key"`
	UserId         string `json:"user_id" gorm:"type:varchar(255);index:idx_user_id;not null;default:'';comment:用户id"`        //用户id
	Username       string `json:"username" gorm:"type:varchar(255);index:idx_user_name;not null;default:'';comment:用户名称"`     //用户名称
	Phone          string `json:"phone" gorm:"type:varchar(255);index:idx_phone;not null;default:'';comment:用户电话"`            //用户电话
	Password       string `json:"password" gorm:"type:varchar(255);not null;default:'';comment:用户密码"`                         //用户密码
	VipEndTime     int64  `json:"vip_end_time" gorm:"type:int(11);not null;default:0;comment:vip到期时间"`                        //vip过期时间
	VipFreeEndTime int64  `json:"vip_free_end_time"  gorm:"type:int(11);not null;default:0;comment:vip体验到期时间"`                //vip体验到期时间
	AttendanceNum  int32  `json:"attendance_num" gorm:"type:int(11);not null;default:0;comment"`                              //用户签到次数
	ShareNum       int32  `json:"share_num" gorm:"type:int(11);not null;default:0;comment:用户分享次数"`                            //用户分享次数
	Identifier     string `json:"identifier" gorm:"type:varchar(255);index:idx_identifier;not null;default:'';comment:唯一标识符"` //唯一标识符
	InviterId      string `json:"inviter_id" gorm:"type:varchar(255);index:idx_inviter_id;not null;default:'';comment:邀请人id"` //邀请人

	UpdatedAt int64 `json:"updated_at" gorm:"type:bigint;not null;default:0;comment:更新时间"`
	CreatedAt int64 `json:"created_at" gorm:"type:bigint;index:idx_created_at;not null;default:0;comment:创建时间"`
}

func (u User) TableName() string {
	return "user"
}

type UserRegisterParams struct {
	Phone    string `json:"phone" validate:"required"`    //手机号
	PassWord string `json:"password" validate:"required"` //密码
}

type UserLoginParams struct {
	Identifier     string `json:"identifier" validate:"required"` //唯一标识符（存在则登录，不存在则会自动先注册在登录）
	InvitationCode string `json:"invitation_code"`                //邀请码（分享成功后，用户注注册时候）

}

type UserLoginResp struct {
	UserId    string `json:"user_id"`    //用户id
	Username  string `json:"username"`   //用户名称
	ExpiresAt int64  `json:"expires_at"` //过期时间
	Token     string `json:"token"`
}

type UserTokenInfo struct {
	UserId     string `json:"user_id"`
	Username   string `json:"user_name"`
	Phone      string `json:"phone"`
	Identifier string `json:"identifier"`
}

type UserSearchParams struct {
	UserId     string `json:"user_id"`    //通过用户id搜索
	Phone      string `json:"phone"`      //通过手机号搜索
	Identifier string `json:"identifier"` //通过唯一标识符搜索
}

type UserSearchResp struct {
	UserId         string `json:"user_id"`           //用户id
	Username       string `json:"username"`          //用户名称
	Phone          string `json:"phone"`             //手机号
	Identifier     string `json:"identifier"`        //唯一标识符
	VipEndTime     int64  `json:"vip_end_time"`      //vip到期时间
	VipFreeEndTime int64  `json:"vip_free_end_time"` //vip体验到期时间
	AttendanceNum  int32  `json:"attendance_num"`    //签到次数
	ShareNum       int32  `json:"share_num"`         //分享次数
}
