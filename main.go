// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/helloworlddan/tortuneai/tortuneai"
)

func main() {
	// Handle requests to "/" by responding with a random joke from the tortuneai lib.
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		joke, err := tortuneai.HitMe("", "devfest24ber-9324")
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}
	fmt.Fprint(w, joke)
	})

	// Listen on incoming TCP requests to $PORT or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Listen and serve HTTP
	fmt.Printf("listening on :%s... \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
