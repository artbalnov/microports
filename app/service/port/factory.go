package port

import (
	"log"

	"github.com/microports/app/service"
	"github.com/microports/app/service/port/internal"
	"github.com/microports/app/service/port/persisntence/mongo"
	"github.com/microports/app/util/env"
	"github.com/microports/app/util/persistence"
)

const (
	ServiceName = "service-port"
)

func Factory() (service.Service, error) {
	// Fetch database URL
	mongoURL, err := env.GetVar(env.MongoURL)
	if err != nil {
		log.Fatalf("[service-port] can't fetch mongo URL: %+v", err)
	}

	// Init mongo session
	mongoSession, err := persistence.GetMongoSession(mongoURL)
	if err != nil {
		log.Fatalf("[service-port] can't init mongo session: %+v", err)
	}

	// Init repository
	portRepository := mongo.NewPortRepository(mongoSession.DB(ServiceName))

	log.Println("[service-port] finish initializing")

	return internal.NewPortService(portRepository), nil
}
