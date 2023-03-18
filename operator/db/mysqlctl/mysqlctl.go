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

package mysqlctl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/frabits/frabit/operator"
)

type Driver struct {
	Host   string
	Port   uint32
	Passwd string
	DBName operator.DBType
	db     *sql.DB
}

func (driver *Driver) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.Driver, error) {
	protocol := "tcp"
	if strings.HasPrefix(config.Host, "/") {
		protocol = "unix"
	}

	params := []string{"multiStatements=true"}

	port := config.Port
	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf("%s@%s(%s:%s)/%s?%s", config.User, protocol, config.Host, port, config.Database, strings.Join(params, "&"))
	if config.Passwd != "" {
		dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", config.User, config.Passwd, protocol, config.Host, port, config.Database, strings.Join(params, "&"))
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	driver.DBName = dbName
	driver.db = db

	return driver, nil
}

func (driver *Driver) GetType() operator.DBType {
	return operator.MySQL
}

func (driver *Driver) Ping(ctx context.Context) error {
	return driver.db.PingContext(ctx)
}

func (driver *Driver) Close(ctx context.Context) error {
	return driver.db.Close()
}
