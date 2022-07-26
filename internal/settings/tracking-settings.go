/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package settings

import (
	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/services"
	"github.com/onflow/flow-cli/pkg/flowkit/util"
	"github.com/spf13/cobra"
)

type FlagsTrackingSettings struct {
	Disabled bool `flag:"disabled" info:"Allow Flow CLI to track command usage statistics"`
}

var TrackingSettingsFlags = FlagsTrackingSettings{}

var TrackingSettings = &command.Command{
	Cmd: &cobra.Command{
		Use:     "tracking",
		Short:   "Configure command usage tracking settings",
		Example: "flow tracking --disabled false",
		Args:    cobra.ExactArgs(0),
	},
	Flags: &TrackingSettingsFlags,
	Run:   disableTrackingSettings,
}

func disableTrackingSettings(
	_ []string,
	_ flowkit.ReaderWriter,
	_ command.GlobalFlags,
	_ *services.Services,
) (command.Result, error) {
	err := util.SetUserTrackingSettings()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
