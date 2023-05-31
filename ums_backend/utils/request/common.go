package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type Register struct {
	Username     string `json:"userName"`
	Password     string `json:"passWord"`
	Email        string `json:"Email"`
	CNname       string `json:"CNname" gorm:"default:'QMPlusUser'"`
	AuthorityId  uint   `json:"authorityId" gorm:"default:888"`
	Enable       int    `json:"enable"`
	AuthorityIds []uint `json:"authorityIds"`
}

type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

type UserName struct {
	Username string `json:"username"`
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
