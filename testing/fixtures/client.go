// Kiebitz - Privacy-Friendly Appointment Scheduling
// Copyright (C) 2021-2021 The Kiebitz Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package fixtures

import (
	"fmt"
	"github.com/kiebitz-oss/services"
	"github.com/kiebitz-oss/services/crypto"
	"github.com/kiebitz-oss/services/testing"
)

type Client struct {
	Key string
}

func (c Client) Setup(fixtures map[string]interface{}) (interface{}, error) {

	settings, ok := fixtures["settings"].(*services.Settings)

	if !ok {
		return nil, fmt.Errorf("settings missing")
	}

	var key *crypto.Key

	if c.Key != "" {

		var ok bool

		key, ok = fixtures[c.Key].(*crypto.Key)

		if !ok {
			return nil, fmt.Errorf("key missing")
		}

	}

	return testing.MakeClient(settings, key), nil
}

func (c Client) Teardown(fixture interface{}) error {
	return nil
}