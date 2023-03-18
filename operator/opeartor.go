// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"context"
)

type DBType string

const (
	MySQL      DBType = "MYSQL"
	Redis      DBType = "REDIS"
	MongoDB    DBType = "MONGODB"
	ClickHouse DBType = "CLICKHOUSE"
)

type DBConnInfo struct {
	// General DB filed
	Host   string
	Port   string
	User   string
	Passwd string

	// For specific DBType
	// For MySQL/ClickHouse
	Database string

	// For MongoDB
	SRV    bool
	AuthDB string
}

// Driver is the interface for supported database driver.
type Driver interface {
	Ping(ctx context.Context) error
	Open(ctx context.Context, dbType DBType, config DBConnInfo) (Driver, error)
	Close(ctx context.Context) error
	GetType() DBType
	// GetDBConn(ctx context.Context) (*sql.DB, error)
	// Execute(ctx context.Context, statement string, createDatabase bool) (int64, error)
	// QueryConn(ctx context.Context, conn *sql.Conn, statement string) ([]interface{}, error)
}
