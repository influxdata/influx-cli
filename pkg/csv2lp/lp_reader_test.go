package csv2lp

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLineProtocolFilter(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"awefw.,weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-blah temperature=32=33 1465839830100400204",
				"weather,,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"awefw.,weather,location=us-east temperature=36 1465839830100400203",
				"    weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-east temperature=36 1465839830100400203 13413413",
				"weather,location=us-central temperature=31 1465839830100400205",
				"# this is a comment",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"    weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
		{
			`# INFLUXDB EXPORT: 2024-03-19T10:00:00Z - 2024-03-19T15:59:59Z
				# DDL
				#CREATE DATABASE "rde-test" WITH NAME autogen
				# DML
				# CONTEXT-DATABASE:rde-test
				# CONTEXT-RETENTION-POLICY:autogen
				# writing tsm data
				Data-12,brand=A,tag_one=GT value=1 1710842400000000000
				Data-12,brand=A,tag_one=GT value=1 1710842580000000000
				Data-12,brand=A,tag_one=GT value=1 1710842760000000000
				Data-12,brand=A,tag_one=GT value=1 1710842940000000000
				Data-12,brand=A,tag_one=GT value=1 1710843120000000000
				Data-12,brand=A,tag_one=GT value=1 1710843300000000000
				Data-12,brand=A,tag_one=GT value=1 1710843480000000000
				Data-12,brand=A,tag_one=GT value=1 1710843660000000000
				Data-12,brand=A,tag_one=GT value=1 1710843840000000000
				Data-12,brand=A,tag_one=GT value=1 1710844020000000000
				Data-12,brand=A,tag_one=GT value=1 1710844200000000000
				Data-12,brand=A,tag_one=GT value=1 1710844380000000000
				Data-12,brand=A,tag_one=GT value=1 1710844560000000000
				Data-12,brand=A,tag_one=GT value=1 1710844740000000000
				Data-12,brand=A,tag_one=GT value=1 1710844920000000000
				Data-12,brand=A,tag_one=GT value=1 1710845100000000000
				Data-12,brand=A,tag_one=GT value=1 1710845280000000000
				Data-12,brand=A,tag_one=GT value=1 1710845460000000000
				Data-12,brand=A,tag_one=GT value=1 1710845640000000000
				Data-12,brand=A,tag_one=GT value=1 1710845820000000000
				Data-12,brand=A,tag_one=GT value=1 1710846000000000000
				Data-12,brand=A,tag_one=GT value=1 1710846180000000000
				Data-12,brand=A,tag_one=GT value=1 1710846360000000000
				Data-12,brand=A,tag_one=GT value=1 1710846540000000000
				Data-12,brand=A,tag_one=GT value=1 1710846720000000000
				Data-12,brand=A,tag_one=GT value=1 1710846900000000000
				Data-12,brand=A,tag_one=GT value=1 1710847080000000000
				Data-12,brand=A,tag_one=GT value=1 1710847260000000000
				Data-12,brand=A,tag_one=GT value=1 1710847440000000000
				Data-12,brand=A,tag_one=GT value=1 1710847620000000000
				Data-12,brand=A,tag_one=GT value=1 1710847800000000000
				Data-12,brand=A,tag_one=GT value=1 1710847980000000000
				Data-12,brand=A,tag_one=GT value=1 1710848160000000000
				Data-12,brand=A,tag_one=GT value=1 1710848340000000000
				Data-12,brand=A,tag_one=GT value=1 1710848520000000000
				Data-12,brand=A,tag_one=GT value=1 1710848700000000000
				Data-12,brand=A,tag_one=GT value=1 1710848880000000000
				Data-12,brand=A,tag_one=GT value=1 1710849060000000000
				Data-12,brand=A,tag_one=GT value=1 1710849240000000000
				Data-12,brand=A,tag_one=GT value=1 1710849420000000000
				Data-12,brand=A,tag_one=GT value=1 1710849600000000000
				Data-12,brand=A,tag_one=GT value=1 1710849780000000000
				Data-12,brand=A,tag_one=GT value=1 1710849960000000000
				Data-12,brand=A,tag_one=GT value=1 1710850140000000000
				Data-12,brand=A,tag_one=GT value=1 1710850320000000000
				Data-12,brand=A,tag_one=GT value=1 1710850500000000000
				Data-12,brand=A,tag_one=GT value=1 1710850680000000000
				Data-12,brand=A,tag_one=GT value=1 1710850860000000000
				Data-12,brand=A,tag_one=GT value=1 1710851040000000000
				Data-12,brand=A,tag_one=GT value=1 1710851220000000000
				Data-12,brand=A,tag_one=GT value=1 1710851400000000000
				Data-12,brand=A,tag_one=GT value=1 1710851580000000000
				Data-12,brand=A,tag_one=GT value=1 1710851760000000000
				Data-12,brand=A,tag_one=GT value=1 1710851940000000000
				Data-12,brand=A,tag_one=GT value=1 1710852120000000000
				Data-12,brand=A,tag_one=GT value=1 1710852300000000000
				Data-12,brand=A,tag_one=GT value=1 1710852480000000000
				Data-12,brand=A,tag_one=GT value=1 1710852660000000000
				Data-12,brand=A,tag_one=GT value=1 1710852840000000000
				Data-12,brand=A,tag_one=GT value=1 1710853020000000000
				Data-12,brand=A,tag_one=GT value=1 1710853200000000000
				Data-12,brand=A,tag_one=GT value=1 1710853380000000000
				Data-12,brand=A,tag_one=GT value=1 1710853560000000000
				Data-12,brand=A,tag_one=GT value=1 1710853740000000000
				Data-12,brand=A,tag_one=GT value=1 1710853920000000000
				Data-12,brand=A,tag_one=GT value=1 1710854100000000000
				Data-12,brand=A,tag_one=GT value=1 1710854280000000000
				Data-12,brand=A,tag_one=GT value=1 1710854460000000000
				Data-12,brand=A,tag_one=GT value=1 1710854640000000000
				Data-12,brand=A,tag_one=GT value=1 1710854820000000000
				Data-12,brand=A,tag_one=GT value=1 1710855000000000000
				Data-12,brand=A,tag_one=GT value=1 1710855180000000000
				Data-12,brand=A,tag_one=GT value=1 1710855360000000000
				Data-12,brand=A,tag_one=GT value=1 1710855540000000000
				Data-12,brand=A,tag_one=GT value=1 1710855720000000000
				Data-12,brand=A,tag_one=GT value=1 1710855900000000000
				Data-12,brand=A,tag_one=GT value=1 1710856080000000000
				Data-12,brand=A,tag_one=GT value=1 1710856260000000000
				Data-12,brand=A,tag_one=GT value=1 1710856440000000000
				Data-12,brand=A,tag_one=GT value=1 1710856620000000000
				Data-12,brand=A,tag_one=GT value=1 1710856800000000000
				Data-12,brand=A,tag_one=GT value=1 1710856980000000000
				Data-12,brand=A,tag_one=GT value=1 1710857160000000000
				Data-12,brand=A,tag_one=GT value=1 1710857340000000000
				Data-12,brand=A,tag_one=GT value=1 1710857520000000000
				Data-12,brand=A,tag_one=GT value=1 1710857700000000000
				Data-12,brand=A,tag_one=GT value=1 1710857880000000000
				Data-12,brand=A,ship=BL value=1 1710842400000000000
				Data-12,brand=A,ship=BL value=1 1710842580000000000
				Data-12,brand=A,ship=BL value=1 1710842760000000000
				Data-12,brand=A,ship=BL value=1 1710842940000000000
				Data-12,brand=A,ship=BL value=1 1710843120000000000
				Data-12,brand=A,ship=BL value=1 1710843300000000000
				`, `Data-12,brand=A,tag_one=GT value=1 1710842400000000000
				Data-12,brand=A,tag_one=GT value=1 1710842580000000000
				Data-12,brand=A,tag_one=GT value=1 1710842760000000000
				Data-12,brand=A,tag_one=GT value=1 1710842940000000000
				Data-12,brand=A,tag_one=GT value=1 1710843120000000000
				Data-12,brand=A,tag_one=GT value=1 1710843300000000000
				Data-12,brand=A,tag_one=GT value=1 1710843480000000000
				Data-12,brand=A,tag_one=GT value=1 1710843660000000000
				Data-12,brand=A,tag_one=GT value=1 1710843840000000000
				Data-12,brand=A,tag_one=GT value=1 1710844020000000000
				Data-12,brand=A,tag_one=GT value=1 1710844200000000000
				Data-12,brand=A,tag_one=GT value=1 1710844380000000000
				Data-12,brand=A,tag_one=GT value=1 1710844560000000000
				Data-12,brand=A,tag_one=GT value=1 1710844740000000000
				Data-12,brand=A,tag_one=GT value=1 1710844920000000000
				Data-12,brand=A,tag_one=GT value=1 1710845100000000000
				Data-12,brand=A,tag_one=GT value=1 1710845280000000000
				Data-12,brand=A,tag_one=GT value=1 1710845460000000000
				Data-12,brand=A,tag_one=GT value=1 1710845640000000000
				Data-12,brand=A,tag_one=GT value=1 1710845820000000000
				Data-12,brand=A,tag_one=GT value=1 1710846000000000000
				Data-12,brand=A,tag_one=GT value=1 1710846180000000000
				Data-12,brand=A,tag_one=GT value=1 1710846360000000000
				Data-12,brand=A,tag_one=GT value=1 1710846540000000000
				Data-12,brand=A,tag_one=GT value=1 1710846720000000000
				Data-12,brand=A,tag_one=GT value=1 1710846900000000000
				Data-12,brand=A,tag_one=GT value=1 1710847080000000000
				Data-12,brand=A,tag_one=GT value=1 1710847260000000000
				Data-12,brand=A,tag_one=GT value=1 1710847440000000000
				Data-12,brand=A,tag_one=GT value=1 1710847620000000000
				Data-12,brand=A,tag_one=GT value=1 1710847800000000000
				Data-12,brand=A,tag_one=GT value=1 1710847980000000000
				Data-12,brand=A,tag_one=GT value=1 1710848160000000000
				Data-12,brand=A,tag_one=GT value=1 1710848340000000000
				Data-12,brand=A,tag_one=GT value=1 1710848520000000000
				Data-12,brand=A,tag_one=GT value=1 1710848700000000000
				Data-12,brand=A,tag_one=GT value=1 1710848880000000000
				Data-12,brand=A,tag_one=GT value=1 1710849060000000000
				Data-12,brand=A,tag_one=GT value=1 1710849240000000000
				Data-12,brand=A,tag_one=GT value=1 1710849420000000000
				Data-12,brand=A,tag_one=GT value=1 1710849600000000000
				Data-12,brand=A,tag_one=GT value=1 1710849780000000000
				Data-12,brand=A,tag_one=GT value=1 1710849960000000000
				Data-12,brand=A,tag_one=GT value=1 1710850140000000000
				Data-12,brand=A,tag_one=GT value=1 1710850320000000000
				Data-12,brand=A,tag_one=GT value=1 1710850500000000000
				Data-12,brand=A,tag_one=GT value=1 1710850680000000000
				Data-12,brand=A,tag_one=GT value=1 1710850860000000000
				Data-12,brand=A,tag_one=GT value=1 1710851040000000000
				Data-12,brand=A,tag_one=GT value=1 1710851220000000000
				Data-12,brand=A,tag_one=GT value=1 1710851400000000000
				Data-12,brand=A,tag_one=GT value=1 1710851580000000000
				Data-12,brand=A,tag_one=GT value=1 1710851760000000000
				Data-12,brand=A,tag_one=GT value=1 1710851940000000000
				Data-12,brand=A,tag_one=GT value=1 1710852120000000000
				Data-12,brand=A,tag_one=GT value=1 1710852300000000000
				Data-12,brand=A,tag_one=GT value=1 1710852480000000000
				Data-12,brand=A,tag_one=GT value=1 1710852660000000000
				Data-12,brand=A,tag_one=GT value=1 1710852840000000000
				Data-12,brand=A,tag_one=GT value=1 1710853020000000000
				Data-12,brand=A,tag_one=GT value=1 1710853200000000000
				Data-12,brand=A,tag_one=GT value=1 1710853380000000000
				Data-12,brand=A,tag_one=GT value=1 1710853560000000000
				Data-12,brand=A,tag_one=GT value=1 1710853740000000000
				Data-12,brand=A,tag_one=GT value=1 1710853920000000000
				Data-12,brand=A,tag_one=GT value=1 1710854100000000000
				Data-12,brand=A,tag_one=GT value=1 1710854280000000000
				Data-12,brand=A,tag_one=GT value=1 1710854460000000000
				Data-12,brand=A,tag_one=GT value=1 1710854640000000000
				Data-12,brand=A,tag_one=GT value=1 1710854820000000000
				Data-12,brand=A,tag_one=GT value=1 1710855000000000000
				Data-12,brand=A,tag_one=GT value=1 1710855180000000000
				Data-12,brand=A,tag_one=GT value=1 1710855360000000000
				Data-12,brand=A,tag_one=GT value=1 1710855540000000000
				Data-12,brand=A,tag_one=GT value=1 1710855720000000000
				Data-12,brand=A,tag_one=GT value=1 1710855900000000000
				Data-12,brand=A,tag_one=GT value=1 1710856080000000000
				Data-12,brand=A,tag_one=GT value=1 1710856260000000000
				Data-12,brand=A,tag_one=GT value=1 1710856440000000000
				Data-12,brand=A,tag_one=GT value=1 1710856620000000000
				Data-12,brand=A,tag_one=GT value=1 1710856800000000000
				Data-12,brand=A,tag_one=GT value=1 1710856980000000000
				Data-12,brand=A,tag_one=GT value=1 1710857160000000000
				Data-12,brand=A,tag_one=GT value=1 1710857340000000000
				Data-12,brand=A,tag_one=GT value=1 1710857520000000000
				Data-12,brand=A,tag_one=GT value=1 1710857700000000000
				Data-12,brand=A,tag_one=GT value=1 1710857880000000000
				Data-12,brand=A,ship=BL value=1 1710842400000000000
				Data-12,brand=A,ship=BL value=1 1710842580000000000
				Data-12,brand=A,ship=BL value=1 1710842760000000000
				Data-12,brand=A,ship=BL value=1 1710842940000000000
				Data-12,brand=A,ship=BL value=1 1710843120000000000
				Data-12,brand=A,ship=BL value=1 1710843300000000000
				`,
		},
	}
	for _, tt := range tests {
		reader := LineProtocolFilter(strings.NewReader(tt.input))
		b, err := io.ReadAll(reader)
		if err != nil {
			t.Errorf("failed reading: %v", err)
			continue
		}
		require.Equal(t, strings.TrimSpace(tt.expected), strings.TrimSpace(string(b)))
	}
}
