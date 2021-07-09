package backup_restore

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/influxdata/influx-cli/v2/api"
)

var semverRegex = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+).*`)

// ServerIsLegacy checks if the InfluxDB server targeted by the backup is running v2.0.x,
// which used different APIs for backups.
func ServerIsLegacy(ctx context.Context, client api.HealthApi) (bool, error) {
	res, err := client.GetHealth(ctx).Execute()
	if err != nil {
		return false, fmt.Errorf("API compatibility check failed: %w", err)
	}
	var version string
	if res.Version != nil {
		version = *res.Version
	}

	matches := semverRegex.FindSubmatch([]byte(version))
	if matches == nil {
		// Assume non-semver versions are only reported by nightlies & dev builds, which
		// should now support the new APIs.
		log.Printf("WARN: Couldn't parse version %q reported by server, assuming latest backup/restore APIs are supported", version)
		return false, nil
	}
	// matches[0] is the entire matched string, capture groups start at 1.
	majorStr, minorStr := matches[1], matches[2]
	// Ignore the err values here because the regex-match ensures we can parse the captured
	// groups as integers.
	major, _ := strconv.Atoi(string(majorStr))
	minor, _ := strconv.Atoi(string(minorStr))

	if major < 2 {
		return false, fmt.Errorf("InfluxDB v%d does not support the APIs required for backup/restore", major)
	}
	return minor == 0, nil
}
