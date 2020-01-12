package kiviksql

import (
	"testing"

	"gitlab.com/flimzy/testy"

	_ "github.com/go-kivik/couchdb"
)

func TestOpen(t *testing.T) {
	type tt struct {
		dsn string
		err string
	}
	tests := testy.NewTable()
	tests.Add("Invalid DSN", tt{
		dsn: "https://%xxx",
		err: `parse https://%xxx: invalid URL escape "%xx"`,
	})
	tests.Add("unknown driver", tt{
		dsn: "foo,http://example.com",
		err: `kivik: unknown driver "foo" (forgotten import?)`,
	})
	tests.Add("default to couch", tt{
		dsn: "http://example.com",
	})
	tests.Add("explicit couch", tt{
		dsn: "couch,http://example.com",
	})

	drv := &drv{}
	tests.Run(t, func(t *testing.T, tt tt) {
		cx, err := drv.Open(tt.dsn)
		testy.Error(t, tt.err, err)
		if cx.(*conn).client == nil {
			t.Fatal("Expected non-nil conn.client")
		}
	})
}
