
package dbcheck

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// ParseMajorVersion reads a Postgres version string (SELECT version())
// and returns the major version number (e.g. "PostgreSQL 16.2 ..." -> 16).
func ParseMajorVersion(versionStr string) (int, error) {
	// Typical version strings:
	// "PostgreSQL 16.2 (Debian 16.2-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc ..."
	// "PostgreSQL 10.3 ..."
	// We'll take the first integer after "PostgreSQL".
	if versionStr == "" {
		return 0, errors.New("empty version string")
	}
	// normalize spacing
	s := strings.TrimSpace(versionStr)

	// find the "PostgreSQL <number>" part
	re := regexp.MustCompile(`PostgreSQL\s+([0-9]+)(?:\.([0-9]+))?`)
	m := re.FindStringSubmatch(s)
	if len(m) < 2 {
		return 0, fmt.Errorf("could not parse version string: %q", versionStr)
	}
	majorStr := m[1]
	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse major version %q: %w", majorStr, err)
	}
	return major, nil
}

// CheckPostgresVersion connects to database using connStr (lib/pq style or DSN)
// runs SELECT version(), parses major version and compares to minMajor.
// Returns nil if OK, or an error if connection fails or version < minMajor.
func CheckPostgresVersion(connStr string, minMajor int) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("open connection: %w", err)
	}
	defer db.Close()

	// Make sure we can actually ping the DB (timeout depends on driver config)
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	var versionStr string
	if err := db.QueryRow("SELECT version()").Scan(&versionStr); err != nil {
		return fmt.Errorf("query version(): %w", err)
	}

	major, err := ParseMajorVersion(versionStr)
	if err != nil {
		return err
	}

	if major < minMajor {
		return fmt.Errorf("postgres version %d detected (string: %q) < required %d", major, versionStr, minMajor)
	}

	return nil
}
