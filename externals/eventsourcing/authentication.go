package eventsourcing

import (
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-user-service/adapters"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-user-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"
	sarama "gopkg.in/Shopify/sarama.v1"
)

func Signup() {
	cons := kafka.NewConsumption()
	cons.Set(entities.Topics_name[int32(entities.Topics_SIGN_UP_REQUESTED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			adapters.Signup(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			close(closeChan)
			break ConsLoop
		}
	}
}

func VerifyUser() {
	cons := kafka.NewConsumption()
	cons.Set(entities.Topics_name[int32(entities.Topics_VERIFY_USER_REQUESTED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			adapters.VerifyUser(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			close(closeChan)
			break ConsLoop
		}
	}
}

func Login() {
	cons := kafka.NewConsumption()
	cons.Set(events.User_name[int32(events.User_LOGIN_REQUESTED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			adapters.Login(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
		case err := <-partCons.Errors():
			errorkit.ErrorHandled(err)
		case <-sig:
			close(closeChan)
			break ConsLoop
		}
	}
}

func ResetPassword() {
	cons := kafka.NewConsumption()
	cons.Set(events.User_name[int32(events.User_RESET_PASSWORD_REQUESTED)], 0, sarama.OffsetNewest)

	partCons, sig, closeChan := cons.Consume()

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			adapters.ResetPassword(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
		case consErr := <-partCons.Errors():
			errorkit.ErrorHandled(consErr.Err)
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)
}
