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

package tools

import (
	"fmt"
	"strings"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/output"
	"github.com/onflow/flow-cli/pkg/flowkit/services"

	devWallet "github.com/onflow/fcl-dev-wallet"
	"github.com/spf13/cobra"
)

type FlagsWallet struct {
	Port uint   `default:"8701" flag:"port" info:"Dev wallet port to listen on"`
	Host string `default:"http://localhost:8888" flag:"emulator-host" info:"Host for access node connection"`
}

var walletFlags = FlagsWallet{}

var DevWallet = &command.Command{
	Cmd: &cobra.Command{
		Use:     "dev-wallet",
		Short:   "Starts a dev wallet",
		Example: "flow dev-wallet",
		Args:    cobra.ExactArgs(0),
	},
	Flags: &walletFlags,
	RunS:  wallet,
}

func wallet(
	_ []string,
	_ flowkit.ReaderWriter,
	_ command.GlobalFlags,
	_ *services.Services,
	state *flowkit.State,
) (command.Result, error) {
	service, err := state.EmulatorServiceAccount()
	if err != nil {
		return nil, err
	}

	key := service.Key().ToConfig()

	conf := devWallet.Config{
		Address:               fmt.Sprintf("0x%s", service.Address().String()),
		PrivateKey:            strings.TrimPrefix(key.PrivateKey.String(), "0x"),
		PublicKey:             strings.TrimPrefix(key.PrivateKey.PublicKey().String(), "0x"),
		AccessNode:            walletFlags.Host,
		AccountKeyID:          "0",
		BaseURL:               "http://localhost:8701",
		ContractFungibleToken: "0xee82856bf20e2aa6",
		ContractFlowToken:     "0x0ae53cb6e3f42a79",
		ContractFUSD:          "0xf8d6e0586b0a20c7",
		ContractFCLCrypto:     "0xf8d6e0586b0a20c7",
		AvatarURL:             "https://avatars.onflow.org/avatar/",
		FlowInitAccountsNo:    "0",
		TokenAmountFLOW:       "100.0",
		TokenAmountFUSD:       "100.0",
	}

	srv, err := devWallet.NewHTTPServer(walletFlags.Port, &conf)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s Starting dev wallet server on port %d\n", output.SuccessEmoji(), walletFlags.Port)
	fmt.Printf("%s  Make sure the emulator is running\n", output.WarningEmoji())

	srv.Start()
	return nil, nil
}
