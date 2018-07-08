// Copyright © 2018 https://github.com/mxssl
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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mxssl/dns/queries"
)

// nsec3Cmd represents the nsec3 command
var nsec3Cmd = &cobra.Command{
	Use:   "nsec3",
	Short: "Get NSEC record version 3 (NSEC3 records)",
	Long: "Get NSEC record version 3 (NSEC3 records)",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queries.GetQ(resolver, "nsec3", args[0])
	},
}

func init() {
	rootCmd.AddCommand(nsec3Cmd)
}
