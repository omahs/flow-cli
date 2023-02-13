/*
 * Flow CLI
 *
 * Copyright 2022 Dapper Labs, Inc.
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

package test

import (
	"fmt"
	cdcTests "github.com/onflow/cadence-tools/test"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/flow-cli/pkg/flowkit"
	"github.com/onflow/flow-cli/pkg/flowkit/config"
	"github.com/onflow/flow-cli/pkg/flowkit/util"
)

func importResolver(
	scriptPath string,
	readerWriter flowkit.ReaderWriter,
	contracts *config.Contracts,
) cdcTests.ImportResolver {
	return func(location common.Location) (string, error) {
		stringLocation, isFileImport := location.(common.StringLocation)
		if !isFileImport {
			return "", fmt.Errorf("cannot import from %s", location)
		}

		importedContract, err := resolveContract(contracts, stringLocation)
		if err != nil {
			return "", err
		}

		importedContractFilePath := util.AbsolutePath(scriptPath, importedContract.Location)

		contractCode, err := readerWriter.ReadFile(importedContractFilePath)
		if err != nil {
			return "", err
		}

		return string(contractCode), nil
	}
}

func resolveContract(
	contracts *config.Contracts,
	stringLocation common.StringLocation,
) (config.Contract, error) {
	relativePath := stringLocation.String()
	for _, contract := range *contracts {
		if contract.Location == relativePath {
			return contract, nil
		}
	}

	return config.Contract{},
		fmt.Errorf("cannot find contract with location '%s' in configuration", relativePath)
}

func fileResolver(
	scriptPath string,
	readerWriter flowkit.ReaderWriter,
) cdcTests.FileResolver {
	return func(path string) (string, error) {
		importFilePath := util.AbsolutePath(scriptPath, path)

		content, err := readerWriter.ReadFile(importFilePath)
		if err != nil {
			return "", err
		}

		return string(content), nil
	}
}