package usecases

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-user-service/entities/domains"

	"github.com/ilhammhdd/kudaki-user-service/entities/domains/user"

	"github.com/ilhammhdd/kudaki-user-service/entities"
	"gopkg.in/Shopify/sarama.v1"

	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/kudaki-user-service/entities/events"

	"github.com/ilhammhdd/go_tool/go_error"
	"github.com/ilhammhdd/kudaki-user-service/entities/commands"
	"golang.org/x/crypto/bcrypt"
)

func Signup(su *commands.SignUp, dbOperator DBOperator, esp EventSourceProducer) {
	esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)

	sdu := &events.SignedUp{
		Profile:     su.Profile,
		User:        su.User,
		Uuid:        su.Uuid,
		EventStatus: &domains.EventStatus{},
	}

	rows, err := dbOperator.Query("SELECT id FROM users WHERE email=?", su.User.Email)

	defer rows.Close()

	if rows.Next() {
		sdu.EventStatus.Code = http.StatusConflict
		sdu.EventStatus.Messages = []string{"user with the given email already exists"}
		sdu.EventStatus.Timestamp = ptypes.TimestampNow()

		_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], sdu)
		go_error.ErrorHandled(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(su.User.Password), bcrypt.MinCost)
	go_error.ErrorHandled(err)

	dbOperator.Command(
		"INSERT INTO users(uuid,email,password,token,role,phone_number) VALUES(?,?,?,?,?,?)",
		su.User.Uuid,
		su.User.Email,
		password,
		"",
		user.Role_name[int32(su.User.Role)],
		su.User.PhoneNumber,
	)

	sdu.EventStatus.Code = http.StatusOK
	sdu.EventStatus.Messages = []string{"successfully registered user"}
	sdu.EventStatus.Timestamp = ptypes.TimestampNow()

	_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], sdu)
	go_error.ErrorHandled(err)
}
