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

package auth

import (
	"time"

	"github.com/frabits/frabit/common/utils"
)

const apiKeyLength int = 32

type ApiKey struct {
	Prefix      string
	PublicAuth  string
	PrivateAuth string
	CreateAt    time.Time
	LastSeen    time.Time
}

func NewAPIKey() *ApiKey {
	return &ApiKey{
		Prefix:      "frabit_tkn",
		PublicAuth:  utils.NewToken(apiKeyLength).Hash,
		PrivateAuth: utils.NewToken(apiKeyLength).Hash,
		CreateAt:    time.Now(),
		LastSeen:    time.Now(),
	}
}