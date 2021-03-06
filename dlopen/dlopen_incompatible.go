// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !cgo !linux

package dlopen

import (
	"errors"
	"unsafe"
)

var errIncompatible = errors.New("package dlopen is incompatible with this platform")

type LibHandle struct {
	Handle  unsafe.Pointer
	Libname string
}

func GetHandle(libs []string) (*LibHandle, error) {
	return nil, errIncompatible
}

func (l *LibHandle) GetSymbolPointer(symbol string) (unsafe.Pointer, error) {
	return nil, errIncompatible
}

func (l *LibHandle) Close() error {
	return errIncompatible
}
