package eventsourcing

import (
	"log"

	"github.com/ilhammhdd/go_tool/go_error"
	"github.com/ilhammhdd/kudaki-user-service/entities"
	"github.com/ilhammhdd/kudaki-user-service/externals/kafka"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type Signup struct{}

func NewSignup() Signup { return Signup{} }

func (su Signup) Work() interface{} {

	cons := kafka.NewConsumption()
	cons.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_COMMAND), sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()

	for {
		select {
		case msg := <-partCons.Messages():
			// adapters.Signup(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
			log.Println("consumed : ", string(msg.Value))
		case errs := <-partCons.Errors():
			go_error.ErrorHandled(errs.Err)
		case <-sig:
			close(closeChan)
		}
	}

	return nil
}

func (su Signup) Handle(interface{}) {}
