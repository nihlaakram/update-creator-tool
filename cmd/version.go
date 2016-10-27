// Copyright (c) 2016, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display wum-uc version information",
	Long:  `Display wum-uc version information.`,
	Run:   versionCommand,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "wum-uc version: %v\n", Version)
	fmt.Fprintf(os.Stdout, "Release date: %v\n", BuildDate)
	fmt.Fprintf(os.Stdout, "OS\\Arch: %v\\%v\n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(os.Stdout, "Go version: %v\n\n", runtime.Version())
}
