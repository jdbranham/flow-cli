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

package accounts

import (
	"bytes"
	"fmt"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/output"
	service "github.com/onflow/flow-cli/pkg/flowkit/services"
	"github.com/onflow/flow-cli/tests"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
)

func setup() (*flowkit.State, *service.Services, *tests.TestGateway) {
	readerWriter := tests.ReaderWriter()
	state, err := flowkit.Init(readerWriter, crypto.ECDSA_P256, crypto.SHA3_256)
	if err != nil {
		panic(err)
	}

	gw := tests.DefaultMockGateway()
	s := service.NewServices(gw.Mock, state, output.NewStdoutLogger(output.NoneLog))

	return state, s, gw
}
func TestAccountsGetCmd(t *testing.T) {
	state, _, _ := setup()
	serviceAcc, _ := state.EmulatorServiceAccount()
	serviceAddress := serviceAcc.Address()
	//cmdOutput := new(bytes.Buffer)
	//Cmd.SetOut(cmdOutput)
	//Cmd.SetErr(cmdOutput)
	//Cmd.SetArgs([]string{"get", "0xf8d6e0586b0a20c7"})
	//Cmd.Execute()
	//expected := "This-is-command-a1"
	//assert.Equal(t, cmdOutput.String(), expected, "actual is not expected")
	_, _, gw := setup()

	//assert.NoError(t, err)
	//assert.Equal(t, serviceAddress, account.Address)
	gw.GetAccount.Run(func(args mock.Arguments) {
		address := args.Get(0).(flow.Address)
		gw.GetAccount.Return(
			tests.NewAccountWithAddress(address.String()), nil,
		)
	})
	//gw := tests.DefaultMockGateway()
	//gw.SendSignedTransaction.Run(func(args mock.Arguments) {
	//	tx := args.Get(0).(*flowkit.Transaction)
	//	assert.Equal(t, serviceAddress, tx.FlowTransaction().Authorizers[0])
	//	assert.Equal(t, serviceAddress, tx.Signer().Address())
	//
	//	gw.SendSignedTransaction.Return(tests.NewTransaction(), nil)
	//})

	b := bytes.NewBufferString("")
	Cmd.SetOut(b)
	//Cmd.SetErr(cmdOutput)
	Cmd.SetArgs([]string{"get", "0xf8d6e0586b0a20c7"})
	Cmd.Execute()
	gw.Mock.AssertCalled(t, "GetAccount", serviceAddress)
	out, err := ioutil.ReadAll(b)
	cmdOutput := string(out)
	fmt.Printf("output is \n")
	fmt.Printf(cmdOutput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "This-is-command-a1", "actual is not expected")
}
