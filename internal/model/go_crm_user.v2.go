package model

const TableNameGoCrmUserV2 = "go_crm_user_v2"

// GoCrmUser Account
type GoCrmUserV2 struct {
	UserID            int32  `gorm:"column:user_id;primaryKey;autoIncrement:true;comment:Account ID" json:"user_id"`                // Account ID
	UserEmail         string `gorm:"column:user_email;not null;comment:Email" json:"user_email"`                                    // Email
	UserPhone         string `gorm:"column:user_phone;not null;comment:Phone Number" json:"user_phone"`                             // Phone Number
	UserUsername      string `gorm:"column:user_username;not null;comment:Username" json:"user_username"`                           // Username
	UserPassword      string `gorm:"column:user_password;not null;comment:Password" json:"user_password"`                           // Password
	UserCreatedAt     int32  `gorm:"column:user_created_at;not null;comment:Creation Time" json:"user_created_at"`                  // Creation Time
	UserUpdatedAt     int32  `gorm:"column:user_updated_at;not null;comment:Update Time" json:"user_updated_at"`                    // Update Time
	UserCreateIPAt    string `gorm:"column:user_create_ip_at;not null;comment:Creation IP" json:"user_create_ip_at"`                // Creation IP
	UserLastLoginAt   int32  `gorm:"column:user_last_login_at;not null;comment:Last Login Time" json:"user_last_login_at"`          // Last Login Time
	UserLastLoginIPAt string `gorm:"column:user_last_login_ip_at;not null;comment:Last Login IP" json:"user_last_login_ip_at"`      // Last Login IP
	UserLoginTimes    int32  `gorm:"column:user_login_times;not null;comment:Login Times" json:"user_login_times"`                  // Login Times
	UserStatus        bool   `gorm:"column:user_status;not null;comment:Status 1:enable, 0:disable, -1:deleted" json:"user_status"` // Status 1:enable, 0:disable, -1:deleted
}

// TableName GoCrmUserV2's table name
func (*GoCrmUserV2) TableName() string {
	return TableNameGoCrmUserV2
}
