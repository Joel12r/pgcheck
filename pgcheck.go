package pgcheck

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// CheckVersion connects to a Postgres database using the given DSN
// and ensures it meets the minimum major version required.
func CheckVersion(dsn string, minVersion int) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer db.Close()

	var versionStr string
	if err := db.QueryRow("SELECT version()").Scan(&versionStr); err != nil {
		return fmt.Errorf("failed to query version: %w", err)
	}

	currentVersion, err := extractMajorVersion(versionStr)
	if err != nil {
		return fmt.Errorf("failed to parse version: %w", err)
	}

	if currentVersion < minVersion {
		return fmt.Errorf("PostgreSQL version %d detected, but minimum required is %d", currentVersion, minVersion)
	}

	return nil
}

// extractMajorVersion parses the output of SELECT version() to get the major version number.
func extractMajorVersion(versionStr string) (int, error) {
	re := regexp.MustCompile(`PostgreSQL\s+(\d+)(?:\.\d+)?`)
	matches := re.FindStringSubmatch(versionStr)
	if len(matches) < 2 {
		return 0, fmt.Errorf("could not parse version string: %s", versionStr)
	}

	version, err := strconv.Atoi(strings.TrimSpace(matches[1]))
	if err != nil {
		return 0, err
	}
	return version, nil
}
