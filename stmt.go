package kiviksql

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

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
	switch t := s.ast.(type) {
	case *ast.SelectStmt:
		_, err := mangoQuery(t)
		return nil, err
	default:
		return nil, errors.New("Unsupported query")
	}
}

func mangoQuery(sel *ast.SelectStmt) (map[string]interface{}, error) {
	where := sel.Where
	if where == nil {
		return nil, nil
	}
	switch t := where.(type) {
	case *ast.BinaryOperationExpr:
		op, err := mangoOp(t.Op)
		if err != nil {
			return nil, err
		}
		var column, value string
		L := t.L
		R := t.R
		// Ensure column name is on left
		if _, ok := R.(*ast.ColumnNameExpr); ok {
			L, R = R, L
		}
		if cn, ok := L.(*ast.ColumnNameExpr); ok {
			column = cn.Name.String()
		} else {
			buf := &bytes.Buffer{}
			where.Format(buf)
			return nil, fmt.Errorf("no column name found in WHERE expression '%s'", buf.String())
		}
		if vl, ok := R.(ast.ValueExpr); ok {
			v, err := json.Marshal(vl.GetValue())
			if err != nil {
				return nil, err
			}
			value = string(v)
		} else {
			buf := &bytes.Buffer{}
			where.Format(buf)
			return nil, fmt.Errorf("no value found in WHERE expression '%s'", buf.String())
		}
		return map[string]interface{}{
			column: map[string]interface{}{
				op: value,
			},
		}, nil
		// fmt.Printf("XXXX: %s %s %v (%T)\n", column, op, value, R)
	}
	fmt.Println(where.Text())
	buf := &bytes.Buffer{}
	where.Format(buf)
	fmt.Println(buf.String())
	tp := where.GetType()
	fmt.Println(tp)
	fmt.Printf("where type %T", where)
	return nil, nil
}
