package kiviksql

import (
	"database/sql/driver"
	"errors"

	"github.com/go-kivik/kivik"
)

type conn struct {
	client *kivik.Client
}

var _ driver.Conn = &conn{}

func (conn) Begin() (driver.Tx, error) {
	return nil, errors.New("kiviksql: transactions not supported by driver")
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, errors.New("kiviksql: not yet implemented")
}
