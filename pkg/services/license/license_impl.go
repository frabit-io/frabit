// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
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

package license

import (
	"context"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"
	"time"
)

type LicenseImpl struct {
	store Store
	cfg   *config.Config
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, db *db.MetaStore) Service {
	metaStore := newStoreImpl(db.Gdb)
	return &LicenseImpl{
		store: metaStore,
		cfg:   cfg,
		log:   log.New("license"),
	}
}

func (s *LicenseImpl) Run(ctx context.Context) error {
	s.log.Info("License service start")
	tick := time.NewTicker(60 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			s.log.Info("shutdown service successfully")
			return ctx.Err()
		case <-tick.C:
			s.log.Info("Im working now")
		}
	}
}

// GetLicense return a validate license from license provider
func (s *LicenseImpl) GetLicense(ctx context.Context) error {
	return nil
}

func (s *LicenseImpl) UpdateLicense(ctx context.Context) error {
	return nil
}
