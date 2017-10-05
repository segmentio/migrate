// +build clickhouse

package main

import (
	_ "github.com/kshvakov/clickhouse"
	_ "github.com/segmentio/migrate/database/clickhouse"
)
