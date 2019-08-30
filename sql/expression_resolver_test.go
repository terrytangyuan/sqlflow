// Copyright 2019 The SQLFlow Authors. All rights reserved.
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

package sql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	trainStatement = `select c1, c2, c3 from kaggle_credit_fraud_training_data 
		TRAIN DNNClassifier 
		WITH 
			%v
		COLUMN
			%v
		LABEL c3 INTO model_table;`
)

func statementWithColumn(column string) string {
	return fmt.Sprintf(trainStatement, "estimator.hidden_units = [10, 20]", column)
}

func statementWithAttrs(attrs string) string {
	return fmt.Sprintf(trainStatement, attrs, "DENSE(c2, 5, comma)")
}

func TestExecResource(t *testing.T) {
	a := assert.New(t)
	parser := newParser()
	s := statementWithAttrs("exec.worker_num = 2")
	r, e := parser.Parse(s)
	a.NoError(e)
	attrs, err := resolveAttribute(&r.trainAttrs)
	a.NoError(err)
	attr := attrs["exec.worker_num"]
	a.Equal(attr.Value, "2")
}
