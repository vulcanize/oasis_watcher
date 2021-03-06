// Copyright 2018 Vulcanize
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

package oasis_dex

import (
	"database/sql"
	"time"
)

type TradeStateModel struct {
	ID        int64          `db:"id"` //id
	Pair      string         //pair
	Guy       string         //maker
	Gem       string         //pay_gem
	Lot       string         //give_amt
	Gal       sql.NullString //taker
	Pie       string         //buy_gem
	Bid       string         //take_amt
	Block     int64
	Tx        string
	Timestamp time.Time `db:"time"`
	Deleted   bool
}
