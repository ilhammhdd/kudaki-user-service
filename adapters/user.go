package adapters

// import (
// 	"database/sql"

// 	"github.com/ilhammhdd/kudaki-entities/user"
// )

// type UserSql struct {
// 	Id          sql.NullInt64
// 	Uuid        sql.NullString
// 	Email       sql.NullString
// 	Password    sql.NullString
// 	Token       sql.NullString
// 	Role        sql.NullString
// 	PhoneNumber sql.NullString
// 	AccountType sql.NullString
// }

// func (us *UserSql) ScanRow(*sql.Row) error {

// 	usrProtoMessage := user.User{
// 		Uuid:        us.Uuid.String,
// 		Email:       us.Email.String,
// 		Password:    us.Password.String,
// 		Token:       us.Token.String,
// 		PhoneNumber: us.PhoneNumber.String}

// 	switch us.Role.String {
// 	case user.Role_name[int32(user.Role_USER)]:
// 		usrProtoMessage.Role = user.Role_USER
// 	case user.Role_name[int32(user.Role_ORGANIZER)]:
// 		usrProtoMessage.Role = user.Role_ORGANIZER
// 	case user.Role_name[int32(user.Role_KUDAKI_TEAM)]:
// 		usrProtoMessage.Role = user.Role_KUDAKI_TEAM
// 	}

// 	switch us.AccountType.String {
// 	case user.AccountType_name[int32(user.AccountType_FACEBOOK)]:
// 		usrProtoMessage.AccountType = user.AccountType_FACEBOOK
// 	case user.AccountType_name[int32(user.AccountType_GOOGLE)]:
// 		usrProtoMessage.AccountType = user.AccountType_GOOGLE
// 	case user.AccountType_name[int32(user.AccountType_NATIVE)]:
// 		usrProtoMessage.AccountType = user.AccountType_NATIVE
// 	}

// 	return nil
// }

// func (us *UserSql) ScanRows(*sql.Rows) error {

// 	return nil
// }
