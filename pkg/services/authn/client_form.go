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
	"log/slog"

	"github.com/frabits/frabit/pkg/infra/log"
)

type Form struct {
	client PasswordClient
	log    *slog.Logger
}

func ProviderForm(client PasswordClient) Client {
	return &Form{
		client: client,
		log:    log.New(ClientForm),
	}
}

func (c *Form) Authenticate(ctx context.Context, req *AuthRequest) (*Identity, error) {
	return c.client.AuthenticatePasswd(ctx, req.HttpReq, req.Username, req.Password)
}

func (c *Form) Name() string {
	return ClientForm
}

func (c *Form) IsEnable() bool {
	return true
}
