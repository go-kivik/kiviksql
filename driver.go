package kiviksql

import (
	"database/sql/driver"
	"regexp"
	"strings"

	"github.com/go-kivik/kivik"
)

type drv struct{}

var _ driver.Driver = &drv{}

var dsnRE = regexp.MustCompile("^([a-z]+),")

// Open accepts a DSN in the following format:
//
//     [driver,]DSN
//
// Where driver, if provided, is the name of the registered Kivik driver, and
// defaults to 'couch', and DSN is the driver-specific Kivik DSN. driver may
// only contain lowercase letters.
func (drv) Open(dsn string) (driver.Conn, error) {
	var kivikDriver string
	if match := dsnRE.FindStringSubmatch(dsn); len(match) > 1 {
		kivikDriver = match[1]
		dsn = strings.TrimPrefix(dsn, kivikDriver)
	} else {
		kivikDriver = "couch"
	}
	client, err := kivik.New(kivikDriver, dsn)
	return &conn{
		client: client,
	}, err
}
