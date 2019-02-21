package adapters

import (
	"github.com/ilhammhdd/kudaki-user-service/entities"
	"gopkg.in/Shopify/sarama.v1"
)

type SignUp struct{}

func (su SignUp) Work() interface{} {
	consumer, err := sarama.NewConsumer(entities.GlobalConfig.KafkaBrokers, nil)
	go_error.HandleError(err)
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition(entities.Topics_name[int32(entities.Topics_USER)], 0, sarama.OffsetNewest)
	go_error.HandleError(err)
	defer partConsumer.Close()

	var signUp commands.SignUp

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partConsumer.Messages():
			{
				if err := proto.Unmarshal(msg.Value, &signUp); err == nil {
					log.Println("IT'S A SIGN UP COMMAND")

					signUp.User.Uuid = uuid.New().String()
					signedUp := usecases.SignUp(&signUp)
					log.Println("RESULT SIGNED UP EVENT : ", signedUp)

					protoSignedUp, err := proto.Marshal(signedUp)
					go_error.HandleError(err)

					syncProd, err := sarama.NewSyncProducer(entities.GlobalConfig.KafkaBrokers, nil)
					go_error.HandleError(err)
					defer syncProd.Close()

					prodMsg := sarama.ProducerMessage{
						Topic:  entities.Topics_name[int32(entities.Topics_USER)],
						Offset: sarama.OffsetNewest,
						Value:  sarama.ByteEncoder(protoSignedUp)}

					_, _, err = syncProd.SendMessage(&prodMsg)
					go_error.HandleError(err)
				} else {
					log.Println("IT'S NOT A SIGN UP COMMAND, THE ERROR : ", err.Error())
				}
			}
		case <-signals:
			{
				break ConsumerLoop
			}
		}
	}

	return nil
}

func (su SignUp) Work() interface{} {

	// consReq := services.ConsumerRequest{Topic: "USER"}

	return nil
}

func (su SignUp) Handle(interface{}) {}
