// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	version "../cli/version"
)

func VersionCmd(log *logrus.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of this command",
		Long:  "Print the version number of this command",
		Run: func(cmd *cobra.Command, args []string) {
			ver, err := version.Get()
			if err != nil {
				panic(err)
			}
			log.WithFields(logrus.Fields{"framework_version": ver.FrameworkVersion, "application_version": ver.ApplicationVersion}).Infof("version")
		},
	}
}