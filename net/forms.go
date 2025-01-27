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

package net

import (
	"github.com/kiprotect/go-helpers/forms"
)

var TCPRateLimitsField = forms.Field{
	Name: "tcp_rate_limits",
	Validators: []forms.Validator{
		forms.IsOptional{},
		forms.IsList{
			Validators: []forms.Validator{
				forms.IsStringMap{
					Form: &forms.Form{
						Fields: []forms.Field{
							{
								Name: "type",
								Validators: []forms.Validator{
									forms.IsIn{Choices: []interface{}{"second", "minute", "hour"}},
								},
							},
							{
								Name: "limit",
								Validators: []forms.Validator{
									forms.IsInteger{HasMin: true, Min: 1},
								},
							},
						},
					},
				},
			},
		},
	},
}
