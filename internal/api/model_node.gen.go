/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"errors"
)

// Node - struct for Node
type Node struct {
	Block      *Block
	Expression *Expression
}

// BlockAsNode is a convenience function that returns Block wrapped in Node
func BlockAsNode(v *Block) Node {
	return Node{Block: v}
}

// ExpressionAsNode is a convenience function that returns Expression wrapped in Node
func ExpressionAsNode(v *Expression) Node {
	return Node{Expression: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Node) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Block
	err = json.Unmarshal(data, &dst.Block)
	if err == nil {
		jsonBlock, _ := json.Marshal(dst.Block)
		if string(jsonBlock) == "{}" { // empty struct
			dst.Block = nil
		} else {
			match++
		}
	} else {
		dst.Block = nil
	}

	// try to unmarshal data into Expression
	err = json.Unmarshal(data, &dst.Expression)
	if err == nil {
		jsonExpression, _ := json.Marshal(dst.Expression)
		if string(jsonExpression) == "{}" { // empty struct
			dst.Expression = nil
		} else {
			match++
		}
	} else {
		dst.Expression = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Block = nil
		dst.Expression = nil

		return errors.New("data matches more than one schema in oneOf(Node)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return errors.New("data failed to match schemas in oneOf(Node)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Node) MarshalJSON() ([]byte, error) {
	if src.Block != nil {
		return json.Marshal(&src.Block)
	}

	if src.Expression != nil {
		return json.Marshal(&src.Expression)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Node) GetActualInstance() interface{} {
	if obj.Block != nil {
		return obj.Block
	}

	if obj.Expression != nil {
		return obj.Expression
	}

	// all schemas are nil
	return nil
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
