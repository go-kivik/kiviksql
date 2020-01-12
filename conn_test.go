package kiviksql

import (
	"bytes"
	"testing"

	"github.com/pingcap/parser/format"

	"gitlab.com/flimzy/testy"
)

func TestPrepare(t *testing.T) {
	type tt struct {
		query string
		err   string
	}
	tests := testy.NewTable()
	tests.Add("select *", tt{
		query: "SELECT * FROM `foo`",
	})

	d := &drv{}
	c, _ := d.Open("http://example.com/")
	tests.Run(t, func(t *testing.T, tt tt) {
		st, err := c.Prepare(tt.query)
		testy.Error(t, tt.err, err)
		buf := &bytes.Buffer{}
		rc := format.NewRestoreCtx(0, buf)
		if err := st.(*stmt).ast.Restore(rc); err != nil {
			t.Fatal(err)
		}
		if d := testy.DiffInterface(testy.Snapshot(t), buf.String()); d != nil {
			t.Fatal(d)
		}
	})
}
