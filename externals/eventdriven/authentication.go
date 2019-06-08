package eventdriven

import (
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-user-service/adapters"
	"github.com/ilhammhdd/kudaki-user-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-service/usecases"
)

func Signup() {
	usecase := &usecases.Signup{DBO: mysql.NewDBOperation(), Producer: kafka.NewProduction()}

	ede := EventDrivenExternal{
		eventDrivenAdapter: new(adapters.Signup),
		eventDrivenUsecase: usecase,
		eventName:          events.UserTopic_SIGN_UP_REQUESTED.String(),
		inTopics:           []string{events.UserTopic_SIGN_UP_REQUESTED.String()},
		outTopic:           events.UserTopic_SIGNED_UP.String()}
	ede.handle()
}
