package eventsourcing

import (
	"github.com/ilhammhdd/kudaki-user-service/adapters"

	"github.com/ilhammhdd/go_tool/go_error"
	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-user-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"
	sarama "gopkg.in/Shopify/sarama.v1"
)

func Signup() {
	cons := kafka.NewConsumption()
	cons.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_COMMAND), sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()

	for {
		select {
		case msg := <-partCons.Messages():
			adapters.Signup(mysql.NewDBOperation(), kafka.NewProduction(), msg.Value)
		case errs := <-partCons.Errors():
			go_error.ErrorHandled(errs.Err)
		case <-sig:
			close(closeChan)
		}
	}
}
