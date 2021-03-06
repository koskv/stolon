// Copyright 2015 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/satori/go.uuid"
)

const (
	StoreBasePath = "stolon/cluster"

	SentinelLeaderKey = "sentinel-leader"
)

type Role string

const (
	RoleUndefined Role = "undefined"
	RoleMaster    Role = "master"
	RoleStandby   Role = "standby"
)

func UID() string {
	u := uuid.NewV4()
	return fmt.Sprintf("%x", u[:4])
}

func UUID() string {
	return uuid.NewV4().String()
}

const (
	stolonPrefix = "stolon_"
)

func StolonName(name string) string {
	return stolonPrefix + name
}

func NameFromStolonName(stolonName string) string {
	return strings.TrimPrefix(stolonName, stolonPrefix)
}

func IsStolonName(name string) bool {
	return strings.HasPrefix(name, stolonPrefix)
}

type Parameters map[string]string

func (s Parameters) Equals(is Parameters) bool {
	return reflect.DeepEqual(s, is)
}
