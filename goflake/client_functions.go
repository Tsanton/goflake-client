package goflake

import (
	"context"
	"encoding/json"

	a "github.com/tsanton/goflake-client/goflake/models/assets"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	m "github.com/tsanton/goflake-client/goflake/models/mergeables"
	u "github.com/tsanton/goflake-client/goflake/utilities"
	"golang.org/x/exp/constraints"

	"github.com/snowflakedb/gosnowflake"
)

type executeScalarConstraint interface {
	constraints.Float | constraints.Integer | string | bool
}

func ExecuteScalar[T executeScalarConstraint](g *GoflakeClient, query string) (T, error) {
	var ret T
	res := g.db.QueryRow(query)
	err := res.Scan(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func RegisterAsset(g *GoflakeClient, asset a.ISnowflakeAsset, stack *u.Stack[a.ISnowflakeAsset]) error {
	stack.Put(asset)
	return CreateAsset(g, asset)
}

func CreateAsset(g *GoflakeClient, asset a.ISnowflakeAsset) error {
	query, numStatements := asset.GetCreateStatement()
	multiStatementContext, _ := gosnowflake.WithMultiStatement(context.Background(), numStatements)
	_, err := g.db.ExecContext(multiStatementContext, query)
	return err
}

func DeleteAssets(g *GoflakeClient, stack *u.Stack[a.ISnowflakeAsset]) {
	for !stack.IsEmpty() {
		err := DeleteAsset(g, stack.Get())
		if err != nil {
			panic("unable to delete all assets in stack")
		}
	}
}

func DeleteAsset(g *GoflakeClient, asset a.ISnowflakeAsset) error {
	query, numStatements := asset.GetDeleteStatement()
	multiStatementContext, _ := gosnowflake.WithMultiStatement(context.Background(), numStatements)
	_, err := g.db.ExecContext(multiStatementContext, query)
	return err
}

func Describe[T e.ISnowflakeEntity](g *GoflakeClient, obj d.ISnowflakeDescribable) (T, error) {
	var ret T
	if obj.IsProcedure() {
		var procedureResponse string
		err := g.db.Get(&procedureResponse, obj.GetDescribeStatement())
		if err != nil {
			return ret, err
		}
		err = json.Unmarshal([]byte(procedureResponse), &ret)
		if err != nil {
			return ret, err
		}
	} else {
		err := g.db.Get(&ret, obj.GetDescribeStatement())
		if err != nil {
			return ret, err
		}
	}

	return ret, nil
}

func MergeInto(g *GoflakeClient, obj m.ISnowflakeMergeable) error {
	return nil
}

func GetMergable(g *GoflakeClient, obj m.ISnowflakeMergeable) (m.ISnowflakeMergeable, error) {
	return nil, nil
}

// TODO: change any for entity.Procedure
// TODO: also need assets.Procedure
func ExecuteProcedure(g *GoflakeClient, proc any) error {
	return nil
}
