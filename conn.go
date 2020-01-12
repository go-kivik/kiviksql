package kiviksql

import (
	"context"
	"database/sql/driver"
	"errors"

	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver" // AST parser driver

	"github.com/go-kivik/kivik"
)

type conn struct {
	client *kivik.Client
	parser *parser.Parser
}

var _ driver.Conn = &conn{}

func (conn) Begin() (driver.Tx, error) {
	return nil, errors.New("kiviksql: transactions not supported by driver")
}

func (c *conn) Close() error {
	return c.client.Close(context.TODO())
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	ast, err := c.parser.ParseOneStmt(query, "", "")
	if err != nil {
		return nil, err
	}
	return &stmt{
		ast: ast,
	}, nil
}
