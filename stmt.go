package kiviksql

import (
	"database/sql/driver"

	"github.com/pingcap/parser/ast"
)

type stmt struct {
	ast ast.StmtNode
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) Exec(_ []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (s *stmt) NumInput() int {
	return 0
}

func (s *stmt) Query(_ []driver.Value) (driver.Rows, error) {
	return nil, nil
}
