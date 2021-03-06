/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file personal-gettransactionbyhash_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */
package test

import (
	"strings"
	"testing"

	"github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func TestGetTransactionByHash(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	tx, err := connection.Eth.GetTransactionByHash("0x4cad48d861726a414e875d0f9395f9a6eeb7fb761108d521aee07017415256c9")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if strings.Compare(tx.BlockHash, "0x0c6ff72bb2eaea3ad532cc294cfacdeb4428223b0ce648bc5c3ca3b98ab64910") != 0 {
		t.Error("Invalid transaction")
		t.Fail()
	}

}
