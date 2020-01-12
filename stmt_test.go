package kiviksql

import (
	"testing"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"gitlab.com/flimzy/testy"
)

func TestMangoQuery(t *testing.T) {
	type tt struct {
		query string
		err   string
	}
	tests := testy.NewTable()
	tests.Add("simple select", tt{
		query: "SELECT * FROM foo",
	})
	tests.Add("simple where", tt{
		query: "SELECT * FROM foo WHERE id=10",
	})
	tests.Add("reversed where", tt{
		query: "SELECT * FROM foo WHERE 10=id",
	})

	parser := parser.New()
	tests.Run(t, func(t *testing.T, tt tt) {
		qq, err := parser.ParseOneStmt(tt.query, "", "")
		if err != nil {
			t.Fatal(err)
		}
		mango, err := mangoQuery(qq.(*ast.SelectStmt))
		testy.Error(t, tt.err, err)
		if d := testy.DiffInterface(testy.Snapshot(t), mango); d != nil {
			t.Error(d)
		}
	})
}
