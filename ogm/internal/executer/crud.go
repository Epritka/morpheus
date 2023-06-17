package executer

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/Epritka/morpheus/ogm/types"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

var (
	baseEntityType = reflect.TypeOf(&types.BaseEntity{})
)

func (e *Executer) Save(model types.Entity) error {
	e.Builder.Create(&entity.Node{
		Base: &entity.Base{
			Alias:      "n",
			Labels:     append(model.Labels(), model.Type()),
			Properties: model.Properies(),
		},
	}).Return(entity.NewAlias("n"))

	err := e.doParse(context.Background(), func(rwc neo4j.ResultWithContext) error {
		record, err := rwc.Single(context.Background())
		if err != nil {
			return err
		}

		data, _ := record.Get("n")
		node := data.(dbtype.Node)

		id, err := strconv.Atoi(strings.Split(node.ElementId, ":")[2])
		if err != nil {
			return err
		}

		model.SetId(int64(id))
		model.SetElementId(node.ElementId)

		bytes, _ := json.Marshal(node.Props)
		json.Unmarshal(bytes, model)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *Executer) Update(model types.Entity) error {
	bytes, _ := json.Marshal(model.Properies())
	e.Builder.
		Match(&entity.Node{Base: &entity.Base{Alias: "n"}}).
		Where(fmt.Sprintf("ID(n) = %s", model.GetElementId())).
		Set(fmt.Sprintf("n = %s", string(bytes))).
		Return(entity.NewAlias("n"))

	err := e.doParse(context.Background(), func(rwc neo4j.ResultWithContext) error {
		record, err := rwc.Single(context.Background())
		if err != nil {
			return err
		}

		data, _ := record.Get("n")
		node := data.(dbtype.Node)

		id, err := strconv.Atoi(strings.Split(node.ElementId, ":")[2])
		if err != nil {
			return err
		}

		model.SetId(int64(id))
		model.SetElementId(node.ElementId)

		bytes, _ := json.Marshal(node.Props)
		json.Unmarshal(bytes, model)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
