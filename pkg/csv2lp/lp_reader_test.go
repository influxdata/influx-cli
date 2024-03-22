package csv2lp

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData = []string{
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10149335861206055 1680843957087735646",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10599446296691895 1680844070301214389",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11674189567565918 1680844094394325888",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.12114214897155762 1680844139344799828",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11209964752197266 1680844301203178190",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1039726734161377 1680844347876425153",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1268465518951416 1680844370106414825",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.12303543090820312 1680844646311648005",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11537313461303711 1680844695335169551",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10717368125915527 1680845112709821861",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1027822494506836 1680845196950133820",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.12188148498535156 1680845320198611842",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10436582565307617 1680845335147764695",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11761021614074707 1680845473279430934",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10162472724914551 1680845611228521315",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.12307000160217285 1680845680031549517",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11197781562805176 1680845736813021569",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11208724975585938 1680846223638448981",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.0956425666809082 1680846370123147298",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1171119213104248 1680846439166725830",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1193227767944336 1680846639652055730",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11297607421875 1680846714839454444",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.11377334594726562 1680846991135713316",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.10156655311584473 1680847055803239552",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.09974861145019531 1680847194100005178",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration=0.1182551383972168 1680847198099239798",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680843957087735646",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844070301214389",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844094394325888",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844139344799828",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844301203178190",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844347876425153",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844370106414825",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844646311648005",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680844695335169551",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845112709821861",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845196950133820",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845320198611842",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845335147764695",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845473279430934",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845611228521315",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845680031549517",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680845736813021569",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846223638448981",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846370123147298",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846439166725830",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846639652055730",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846714839454444",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680846991135713316",
	"celery_task_duration,env=MOOO,host=foobar-batch,queue=high,service=BATCH,task=monitoring.tasks.celery_throughput_high duration_limit_exceeded=0i 1680847055803239552",
}

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
			strings.Join(testData, "\n"),
			strings.Join(testData, "\n"),
		},
	}
	for _, tt := range tests {
		reader := LineProtocolFilter(strings.NewReader(tt.input))
		// test data should fit in b
		b := make([]byte, 163984)
		bytesRead, err := io.ReadFull(reader, b)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			t.Errorf("failed reading: %v", err)
			continue
		}
		require.Equal(t, strings.TrimSpace(tt.expected), strings.TrimSpace(string(b[0:bytesRead])))
	}
}
