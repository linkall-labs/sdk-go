/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package expression

import (
	"fmt"

	cesql "github.com/cloudevents/sdk-go/sql/v2"
	"github.com/cloudevents/sdk-go/sql/v2/utils"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type dataIdentifierExpression struct {
	identifier string
}

func (l dataIdentifierExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	// TODO
	value := utils.GetAttribute(event, l.identifier)
	if value == nil {
		return nil, fmt.Errorf("missing attribute '%s'", l.identifier)
	}

	return value, nil
}

func NewDataIdentifierExpression(identifier string) cesql.Expression {
	return dataIdentifierExpression{identifier: identifier}
}
