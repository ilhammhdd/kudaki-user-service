package usecases

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/ilhammhdd/go_error"
	"github.com/ilhammhdd/kudaki-user-service/entities/commands"
	"github.com/ilhammhdd/kudaki-user-service/entities/events"
)

func SignUp(su *commands.SignUp) *events.SignedUp {
	log.Println("LET'S SEE THE COMMAND : ", *su)
	dbOperation := DBOperation{
		Args: []interface{}{su.User.Username, su.User.Email},
		Stmt: "SELECT id FROM users WHERE username=? OR email=?"}

	rows, err := DBQueryOperation(&dbOperation)
	go_error.HandleError(err)

	defer rows.Close()

	var status events.Status

	if rows.Next() {
		log.Println("USER WITH THE GIVEN USERNAME OR EMAIL IS ALREADY EXISTS")
		currentTime := time.Now()
		status.Success = false
		status.Errors = []string{"USER WITH THE GIVEN USERNAME OR EMAIL IS ALREADY EXISTS"}
		status.Timestamp = &timestamp.Timestamp{
			Nanos:   int32(currentTime.UnixNano()),
			Seconds: currentTime.Unix()}

		return &events.SignedUp{Status: &status}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(su.User.Password), bcrypt.MinCost)
	go_error.HandleError(err)

	dbOperation.Stmt = "INSERT INTO users(uuid,username,email,password,account_type) VALUES(?,?,?,?,?)"
	dbOperation.Args = []interface{}{
		su.User.Uuid,
		su.User.Username,
		su.User.Email,
		string(password),
		su.User.AccountType.String()}

	DBCommandOperation(&dbOperation)

	currentTime := time.Now()

	status.Success = true
	status.Timestamp = &timestamp.Timestamp{
		Nanos:   int32(currentTime.UnixNano()),
		Seconds: currentTime.Unix()}
	status.Errors = nil

	log.Println("NO PROBLEM")

	return &events.SignedUp{
		Uuid:    su.Uuid,
		User:    su.User,
		Profile: su.Profile,
		Status:  &status}
}
