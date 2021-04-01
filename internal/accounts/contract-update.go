/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
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

package accounts

import (
	"fmt"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowcli/services"
	"github.com/spf13/cobra"
)

type flagsUpdateContract struct {
	Signer  string `default:"emulator-account" flag:"signer"`
	Results bool   `default:"false" flag:"results" info:"⚠️  DEPRECATED: results are provided by default"`
}

var updateFlags = flagsUpdateContract{}

var UpdateCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "update-contract <name> <filename>",
		Short:   "Update a contract deployed to an account",
		Example: `flow accounts update-contract FungibleToken ./FungibleToken.cdc`,
		Args:    cobra.ExactArgs(2),
	},
	Flags: &updateFlags,
	Run: func(
		cmd *cobra.Command,
		args []string,
		globalFlags command.GlobalFlags,
		services *services.Services,
	) (command.Result, error) {
		if createFlags.Results {
			return nil, fmt.Errorf("⚠️  DEPRECATED: results flags is deprecated, results are by default included in all executions.")
		}

		account, err := services.Accounts.AddContract(
			updateFlags.Signer,
			args[0], // name
			args[1], // filename
			true,
		)
		if err != nil {
			return nil, err
		}

		return &AccountResult{
			Account:  account,
			showCode: false,
		}, nil
	},
}
