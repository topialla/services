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

package forms

import (
	"github.com/kiprotect/go-helpers/forms"
)

var ECDHEncryptedDataForm = forms.Form{
	Name: "ecdhEncryptedData",
	Fields: []forms.Field{
		{
			Name:        "iv",
			Description: "IV value associated with the encrypted data.",
			Validators: []forms.Validator{
				forms.IsBytes{
					Encoding:  "base64",
					MinLength: 10,
					MaxLength: 20,
				},
			},
		},
		{
			Name:        "data",
			Description: "The encrypted data.",
			Validators: []forms.Validator{
				forms.IsBytes{
					Encoding:  "base64",
					MinLength: 1,
					MaxLength: 200000,
				},
			},
		},
		PublicKeyField,
	},
}
