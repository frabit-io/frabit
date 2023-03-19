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

package mongoctl

import (
	"context"
	"fmt"
	"github.com/frabits/frabit/pkg/operator"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoType string

const (
	Standalone    MongoType = "STANDALONE"
	ReplicateSet  MongoType = "REPLICATESET"
	SharedCluster MongoType = "SHAREDCLUSTER"
)

// Driver is the MongoDB driver.
type Driver struct {
	DBType  operator.DBType
	connCfg operator.DBConnInfo
	client  *mongo.Client
}

func (driver *Driver) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.DBOperator, error) {
	connectionURI := genMongoDBConnectionURI(config)
	opts := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create MongoDB client")
	}
	driver.client = client
	driver.connCfg = config
	driver.DBType = dbName
	return driver, nil
}

func (driver *Driver) GetType() operator.DBType {
	return operator.MongoDB
}

func (driver *Driver) Ping(ctx context.Context) error {
	return nil
}

func (driver *Driver) Close(ctx context.Context) error {
	return driver.client.Disconnect(ctx)
}

// genMongoDBConnectionURI generate a connection string based provide MongoDB config
func genMongoDBConnectionURI(connCfg operator.DBConnInfo) string {
	connectionURI := "mongodb://"
	if connCfg.SRV {
		connectionURI = "mongodb+srv://"
	}
	if connCfg.User != "" {
		percentEncodingUser := replaceCharacterWithPercentEncoding(connCfg.User)
		percentEncodingPasswd := replaceCharacterWithPercentEncoding(connCfg.Passwd)
		connectionURI = fmt.Sprintf("%s%s:%s@", connectionURI, percentEncodingUser, percentEncodingPasswd)
	}
	connectionURI = fmt.Sprintf("%s%s", connectionURI, connCfg.Host)
	if connCfg.Port != "" {
		connectionURI = fmt.Sprintf("%s:%s", connectionURI, connCfg.Port)
	}
	if connCfg.AuthDB != "" {
		connectionURI = fmt.Sprintf("%s/%s", connectionURI, connCfg.Database)
	}
	// We use admin as the default authentication database.
	// https://www.mongodb.com/docs/manual/reference/connection-string/#mongodb-urioption-urioption.authSource
	authenticationDatabase := connCfg.AuthDB
	if authenticationDatabase == "" {
		authenticationDatabase = "admin"
	}

	if connCfg.Database == "" {
		connectionURI = fmt.Sprintf("%s/", connectionURI)
	}
	connectionURI = fmt.Sprintf("%s?authSource=%s", connectionURI, authenticationDatabase)

	return connectionURI
}

func replaceCharacterWithPercentEncoding(s string) string {
	m := map[string]string{
		":": `%3A`,
		"/": `%2F`,
		"?": `%3F`,
		"#": `%23`,
		"[": `%5B`,
		"]": `%5D`,
		"@": `%40`,
	}
	for k, v := range m {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}