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

package authn

import (
	"context"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/user"
	"log/slog"
)

type Frabit struct {
	user user.Service
	log  *slog.Logger
}

func ProviderFrabit(user user.Service) PasswordClient {
	return &Frabit{
		user: user,
		log:  log.New("authn.client.frabit"),
	}
}

func (c *Frabit) AuthenticatePasswd(ctx context.Context, username string, password string) (*Identity, error) {
	return &Identity{}, nil
}
