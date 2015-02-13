/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import (
	"strings"
	"testing"
	"fmt"
	"os"
)

const (
	FAKE_MARATHON_URL = "http://127.0.0.1:3000,127.0.0.1:3000"
	FAKE_APP_NAME = "/fake_app"
)

var test_client Marathon

func NewFakeMarathonEndpoint() {
	if test_client == nil {
		var err error
		config := DefaultConfig
		config.URL = FAKE_MARATHON_URL
		test_client, err = NewClient(config)
		if err != nil {
			fmt.Printf("Failed to create the fake client, %s, error: %s", FAKE_MARATHON_URL, err)
			os.Exit(1)
		}
	}
}

func AssertOnNull(data interface {}, t *testing.T) {
	if data == nil {
		t.Fail()
	}
}

func AssertOnError(err error, t *testing.T) {
	if err != nil {
		t.Fail()
	}
}

func AssertOnBool(value, expect bool, t *testing.T) {
	if value != expect {
		t.Fail()
	}
}

func AssertOnString(value, expected string, t *testing.T) {
	if !strings.Contains(value, expected) {
		t.Errorf("failed, value %s, expected: %s", value, expected)
	}
}

func AssertOnInteger(value, expected int, t *testing.T) {
	if value != expected {
		t.Errorf("failed, value %d, expected: %d", value, expected)
	}
}

