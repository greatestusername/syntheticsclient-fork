// Copyright 2021 Splunk, Inc.
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

package syntheticsclient

import (
	"net/http"
	"testing"
)

var (
	deleteBrowserRespBody = `{"result":"success","message":"testtaste successfully deleted","errors":[]}`
)

func TestDeleteBrowseCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/real_browsers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		_, err := w.Write([]byte(deleteBrowserRespBody))
		if err != nil {
			t.Errorf("returned error: %#v", err)
		}
	})

	resp, err := testClient.DeleteBrowserCheck(10)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "testtaste successfully deleted" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Message, "testtaste successfully deleted")
	}
	if resp.Result != "success" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Result, "success")
	}
}
