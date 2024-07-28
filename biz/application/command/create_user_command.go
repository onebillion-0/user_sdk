package command

type CreateUserCommand struct {
	Uid         int64
	NickName    string
	Avatar      string
	Sex         string
	PassWord    string
	PhoneNumber string
	IdCard      string
	Age         int
}

type MemberCommand struct {
	Uid      int64
	NickName string
	Password string
}
