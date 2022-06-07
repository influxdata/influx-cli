package v1repl_test

import (
	"testing"

	v1repl "github.com/influxdata/influx-cli/v2/clients/v1_repl"
)

var point = `weather,location=us-midwest temperature=82 1465839830100400200`
var db = "test_database"
var quotedDb = `"test_database"`
var rp = "test_retention_policy"
var quotedRp = `"test_retention_policy"`

type InsertTestCase struct {
	cmd           string
	expectedDb    string
	expectedRp    string
	expectedPoint string
}

func (it *InsertTestCase) Test(t *testing.T) {
	db, rp, point, isInsert := v1repl.ParseInsert(it.cmd)
	if !isInsert {
		t.Errorf("%q should be a valid INSERT command", it.cmd)
	} else {
		if db != it.expectedDb {
			t.Errorf("Db mismatch: expected: %s, got: %s", it.expectedDb, db)
		}
		if rp != it.expectedRp {
			t.Errorf("Rp mismatch: expected: %s, got: %s", it.expectedRp, rp)
		}
		if point != it.expectedPoint {
			t.Errorf("Point mismatch: expected: %s, got: %s", it.expectedPoint, point)
		}
	}
}

var insertIntoCmds = []InsertTestCase{
	{
		cmd:           `insert   into   ` + quotedDb + "   " + point,
		expectedDb:    db,
		expectedRp:    "",
		expectedPoint: point,
	},
	{
		cmd:           `insert   into   ` + quotedDb + "." + quotedRp + "  " + point,
		expectedDb:    db,
		expectedRp:    rp,
		expectedPoint: point,
	},
	{
		cmd:           `insert   into   ` + quotedDb + "   " + point,
		expectedDb:    db,
		expectedRp:    "",
		expectedPoint: point,
	},
	{
		cmd:           `InSeRt  INtO  ` + quotedDb + "." + quotedRp + "   " + point,
		expectedDb:    db,
		expectedRp:    rp,
		expectedPoint: point,
	},
	{
		cmd:           `INSERT  INTO  ` + quotedDb + "." + quotedRp + "   " + point,
		expectedDb:    db,
		expectedRp:    rp,
		expectedPoint: point,
	},
}

var insertCmds = []InsertTestCase{
	{
		cmd:           `INSERT ` + point,
		expectedDb:    "",
		expectedRp:    "",
		expectedPoint: point,
	},
	{
		cmd:           `insert ` + point,
		expectedDb:    "",
		expectedRp:    "",
		expectedPoint: point,
	},
	{
		cmd:           `INSERT    ` + point,
		expectedDb:    "",
		expectedRp:    "",
		expectedPoint: point,
	},
	{
		cmd:           `insert      ` + point,
		expectedDb:    "",
		expectedRp:    "",
		expectedPoint: point,
	},
}

var invalidCmds = []string{
	`insert` + point,
	`insert into ` + point,
	`INSERT` + point,
	`insert into . ` + point,
	`insert into ` + quotedDb + `. ` + point,
	`insertinto ` + quotedDb + ` ` + point,
	`insert into ` + quotedDb + `.` + quotedRp + ". " + point,
}

func TestParseInsert(t *testing.T) {
	t.Parallel()
	for _, insertCmd := range insertCmds {
		insertCmd.Test(t)
	}
}

func TestParseInsertInto(t *testing.T) {
	t.Parallel()
	for _, insertIntoCmd := range insertIntoCmds {
		insertIntoCmd.Test(t)
	}
}

func TestParseInsertInvalid(t *testing.T) {
	t.Parallel()
	for _, cmd := range invalidCmds {
		if _, _, _, isValid := v1repl.ParseInsert(cmd); isValid {
			t.Errorf("%q should be an invalid INSERT command", cmd)
		}
	}
}
