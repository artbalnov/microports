package persisntence

import "github.com/microports/app/service/port/persisntence/model"

// Port data source contract
type PortRepository interface {
	Save(port *model.Port) error
	SaveAll(ports []*model.Port) error
	GetAll() ([]*model.Port, error)
}
