package eventsourcing

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-user-service/adapters"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-user-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"
	sarama "gopkg.in/Shopify/sarama.v1"
)

func Signup() {
	consMemberName := "SignupRequested"
	topic := events.UserTopic_name[int32(events.UserTopic_SIGN_UP_REQUESTED)]
	groupID := uuid.New().String()

	for i := 0; i < 5; i++ {
		consMember := kafka.NewConsumptionMember(groupID, []string{topic}, sarama.OffsetNewest, consMemberName, i)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		safekit.Do(func() {
			<-consMember.Ready
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					adapters.Signup(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					return
				}
			}
		})
	}
}

func VerifyUser() {
	consMemberName := "VerifyUserRequested"
	topic := events.UserTopic_name[int32(events.UserTopic_VERIFY_USER_REQUESTED)]
	groupID := uuid.New().String()

	for i := 0; i < 5; i++ {
		consMember := kafka.NewConsumptionMember(groupID, []string{topic}, sarama.OffsetNewest, consMemberName, i)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		safekit.Do(func() {
			<-consMember.Ready
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					key, value, err := adapters.VerifyUser(mysql.NewDBOperation(), msg.Value)
					errorkit.ErrorHandled(err)

					prod := kafka.NewProduction()
					prod.Set(events.UserTopic_name[int32(events.UserTopic_SIGNED_UP)])
					partition, offset, err := prod.SyncProduce(key, value)
					errorkit.ErrorHandled(err)
					log.Printf("produced Signedup: partition = %d, offset = %d, key = %s", partition, offset, key)
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					return
				}
			}
		})
	}
}

func Login() {
	consMemberName := "LoginRequested"
	topic := events.UserTopic_name[int32(events.UserTopic_LOGIN_REQUESTED)]
	groupID := uuid.New().String()

	for i := 0; i < 5; i++ {
		member := kafka.NewConsumptionMember(groupID, []string{topic}, sarama.OffsetNewest, consMemberName, i)

		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt)

		safekit.Do(func() {
			<-member.Ready
			defer close(member.Close)
			for {
				select {
				case msg := <-member.Messages:
					key, value, err := adapters.Login(mysql.NewDBOperation(), msg)
					errorkit.ErrorHandled(err)

					prod := kafka.NewProduction()
					prod.Set(events.UserTopic_name[int32(events.UserTopic_LOGGED_IN)])

					partition, offset, err := prod.SyncProduce(key, value)
					log.Printf("produced Loggedin: partition = %d, offset = %d, key = %s", partition, offset, key)
				case errs := <-member.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					return
				}
			}
		})
	}
}

func ChangePassword() {
	consMemberName := "ChangePasswordRequested"
	topic := events.UserTopic_name[int32(events.UserTopic_CHANGE_PASSWORD_REQUESTED)]
	groupID := uuid.New().String()

	for i := 0; i < 5; i++ {
		member := kafka.NewConsumptionMember(groupID, []string{topic}, sarama.OffsetNewest, consMemberName, i)

		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt)

		prod := kafka.NewProduction()
		prod.Set(events.UserTopic_name[int32(events.UserTopic_PASSWORD_CHANGED)])

		safekit.Do(func() {
			<-member.Ready
			defer close(member.Close)
			for {
				select {
				case msg := <-member.Messages:
					key, value, err := adapters.ChangePassword(mysql.NewDBOperation(), msg.Value)
					errorkit.ErrorHandled(err)

					partition, offset, err := prod.SyncProduce(key, value)
					log.Printf("produced PasswordChanged : partition = %d, offset = %d, key = %s", partition, offset, key)
				case errs := <-member.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					return
				}
			}
		})
	}
}

type SendResetPasswordEmail struct{}

func (pr SendResetPasswordEmail) Handle(interface{}) {}

func (pr SendResetPasswordEmail) Work() interface{} {

	groupID := uuid.New().String()
	topics := []string{events.UserTopic_name[int32(events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED)]}

	for i := 0; i < 5; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, "SendResetPasswordEmailRequested", i)

		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		safekit.Do(func() {
			<-consMember.Ready
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					prAdapter := adapters.SendResetPasswordEmail{
						Key:       string(msg.Key),
						Message:   &msg.Value,
						Offset:    msg.Offset,
						Partition: msg.Partition,
					}
					key, prMsg, err := prAdapter.SendEmail(mysql.NewDBOperation())
					if err == nil {
						log.Printf("consumed ResetPasswordRequested : key = %s, offset = %v, partition = %v", string(msg.Key), msg.Offset, msg.Partition)
						pr.produce(key, &prMsg)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}

	return nil
}

func (pr SendResetPasswordEmail) produce(key string, msg *[]byte) {
	prod := kafka.NewProduction()
	prod.Set(events.UserTopic_name[int32(events.UserTopic_RESET_PASSWORD_EMAIL_SENT)])
	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, *msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced ResetPasswordEmailSent : partition = %d, offset = %d, key = %s, duration = %f seconds", partition, offset, key, duration.Seconds())
}

// Boundary

type ResetPassword struct{}

func (rp ResetPassword) Handle(interface{}) {}

func (rp ResetPassword) Work() interface{} {

	groupID := uuid.New().String()
	topics := []string{events.UserTopic_name[int32(events.UserTopic_RESET_PASSWORD_REQUESTED)]}

	for i := 0; i < 5; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, "ResetPasswordRequested", i)

		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		safekit.Do(func() {
			<-consMember.Ready
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					prAdapter := adapters.ResetPassword{
						Key:       string(msg.Key),
						Message:   &msg.Value,
						Offset:    msg.Offset,
						Partition: msg.Partition,
					}
					key, prMsg, err := prAdapter.Reset(mysql.NewDBOperation())
					if err == nil {
						log.Printf("consumed ResetPasswordRequested : key = %s, offset = %v, partition = %v", string(msg.Key), msg.Offset, msg.Partition)
						rp.produce(key, &prMsg)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}

	return nil
}

func (rp ResetPassword) produce(key string, msg *[]byte) {
	prod := kafka.NewProduction()
	prod.Set(events.UserTopic_name[int32(events.UserTopic_PASSWORD_RESETED)])
	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, *msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced PasswordReseted : partition = %d, offset = %d, key = %s, duration = %f seconds", partition, offset, key, duration.Seconds())
}
