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

package user

import (
	"context"
	"log/slog"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
)

type service struct {
	log   *slog.Logger
	cfg   *config.Config
	store Store
}

func NewService(conf *config.Config) Service {
	us := &service{
		log: log.New("user"),
		cfg: conf,
	}

	return us
}

func (s *service) CreateUser(ctx context.Context, user *User) (uint32, error) {
	return uint32(0), nil
}

func (s *service) UpdateUser(ctx context.Context) error {
	return nil
}

func (s *service) GetUserById(ctx context.Context) error {
	return nil
}

func (s *service) GetUserByName(ctx context.Context) error {
	return nil
}
