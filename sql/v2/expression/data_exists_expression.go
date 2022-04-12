/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package expression

import (
	cesql "github.com/cloudevents/sdk-go/sql/v2"
	"github.com/cloudevents/sdk-go/sql/v2/utils"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type dataExistsExpression struct {
	identifier string
}

func (l dataExistsExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	// TODO
	return utils.ContainsAttribute(event, l.identifier), nil
}

func NewDataExistsExpression(identifier string) cesql.Expression {
	return dataExistsExpression{identifier: identifier}
}
