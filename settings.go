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

package services

type RPCSettings struct {
	BindAddress string `json:"bind_address"`
}

type StorageSettings struct {
	SettingsTTLDays int64                  `json:"settings_ttl_days"`
	RPC             *JSONRPCServerSettings `json:"rpc"`
}

type AppointmentsSettings struct {
	RPC  *JSONRPCServerSettings `json:"rpc"`
	Keys []*Key                 `json:"keys"`
}

func (a *AppointmentsSettings) Key(name string) *Key {
	for _, key := range a.Keys {
		if key.Name == name {
			return key
		}
	}
	return nil
}

type Key struct {
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`
	Format    string                 `json:"format"`
	Params    map[string]interface{} `json:"params"`
	PublicKey []byte                 `json:"public_key"`
	Purposes  []string               `json:"purposes"`
	// only defined for local signing operations
	PrivateKey []byte `json:"private_key"`
}

type SigningSettings struct {
	Keys []*Key `json:"keys"`
}

type DatabaseSettings struct {
	Type     string `json:"type"`
	Settings *interface{}
}

type Settings struct {
	Signing      *SigningSettings      `json:"signing"`
	Definitions  *Definitions          `json:"definitions"`
	Storage      *StorageSettings      `json:"storage"`
	Appointments *AppointmentsSettings `json:"appointments"`
	Database     *DatabaseSettings     `json:"database"`
}

type TLSSettings struct {
	CACertificateFile string `json:"ca_certificate_file"`
	CertificateFile   string `json:"certificate_file"`
	KeyFile           string `json:"key_file"`
}

// Settings for the JSON-RPC server
type JSONRPCClientSettings struct {
	TLS      *TLSSettings `json:"tls"`
	Endpoint string       `json:"endpoint"`
	Local    bool         `json:"local"`
}

type CorsSettings struct {
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedHosts   []string `json:"allowed_hosts"`
	AllowedMethods []string `json:"allowed_methods"`
}

// Settings for the JSON-RPC server
type JSONRPCServerSettings struct {
	Cors        *CorsSettings `json:"cors"`
	TLS         *TLSSettings  `json:"tls"`
	BindAddress string        `json:"bind_address"`
}

// Settings for the JSON-RPC server
type HTTPServerSettings struct {
	TLS         *TLSSettings `json:"tls"`
	BindAddress string       `json:"bind_address"`
}
