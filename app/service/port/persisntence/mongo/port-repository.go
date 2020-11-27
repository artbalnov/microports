package mongo

import (
	"github.com/microports/app/service/port/persisntence"
	"github.com/microports/app/service/port/persisntence/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const portsCollectionName = "ports"

type portRepository struct {
	DB *mgo.Database
}

func NewPortRepository(database *mgo.Database) persisntence.PortRepository {
	return &portRepository{
		DB: database,
	}
}

func (rcv *portRepository) Save(port *model.Port) error {
	_, err := rcv.DB.C(portsCollectionName).Upsert(bson.M{"_id": port.ID}, port)
	if err != nil {
		return err
	}

	return nil
}

func (rcv *portRepository) SaveAll(ports []*model.Port) error {
	bulk := rcv.DB.C(portsCollectionName).Bulk()

	for i := range ports {
		bulk.Upsert(bson.M{"_id": ports[i].ID}, ports[i])
	}

	if _, err := bulk.Run(); err != nil {
		return err
	}

	return nil
}

func (rcv *portRepository) GetAll() ([]*model.Port, error) {
	query := rcv.DB.C(portsCollectionName).Find(bson.M{})

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var ports []*model.Port

	if count == 0 {
		return ports, nil
	}

	err = query.All(&ports)
	if err != nil {
		return nil, err
	}

	return ports, nil
}
