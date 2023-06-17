package ogm

import (
	"context"

	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm/internal/executer"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Driver struct {
	Driver neo4j.DriverWithContext
}

func Connect(config *config.Config, configurers ...func(c *neo4j.Config)) (*Driver, error) {
	driver, err := neo4j.NewDriverWithContext(
		config.URI,
		neo4j.BasicAuth(
			config.Username,
			config.Password,
			config.Realm,
		),
		configurers...,
	)
	if err != nil {
		return nil, err
	}

	return &Driver{
		Driver: driver,
	}, nil
}

func (driver *Driver) NewExecuter() *executer.Executer {
	return executer.New(driver.Driver.NewSession(context.Background(), neo4j.SessionConfig{}))
}

func (driver *Driver) NewExecuterWithContext(ctx context.Context) *executer.Executer {
	return executer.New(driver.Driver.NewSession(ctx, neo4j.SessionConfig{}))
}
