package expression

import (
	"reflect"
	"strings"
	"testing"

	"github.com/matthew188/aws-sdk-go-v2/service/dynamodb/types"
)

// updateErrorMode will help with error cases and checking error types
type updateErrorMode string

const (
	noUpdateError             updateErrorMode = ""
	invalidUpdateOperand                      = "BuildOperand error"
	unsetSetValue                             = "unset parameter: SetValueBuilder"
	unsetUpdate                               = "unset parameter: UpdateBuilder"
	emptyOperationBuilderList                 = "operationBuilder list is empty"
)

func TestBuildOperation(t *testing.T) {
	cases := []struct {
		name     string
		input    operationBuilder
		expected exprNode
		err      updateErrorMode
	}{
		{
			name: "set operation",
			input: operationBuilder{
				name:  Name("foo"),
				value: Value(5),
				mode:  setOperation,
			},
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						values: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
						},
						fmtExpr: "$v",
					},
				},
				fmtExpr: "$c = $c",
			},
		},
		{
			name: "add operation",
			input: operationBuilder{
				name:  Name("foo"),
				value: Value(5),
				mode:  addOperation,
			},
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						values: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
						},
						fmtExpr: "$v",
					},
				},
				fmtExpr: "$c $c",
			},
		},
		{
			name: "remove operation",
			input: operationBuilder{
				name: Name("foo"),
				mode: removeOperation,
			},
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "$c",
			},
		},
		{
			name: "invalid operand",
			input: operationBuilder{
				name: Name(""),
				mode: removeOperation,
			},
			err: invalidUpdateOperand,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.buildOperation()
			if c.err != noUpdateError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestUpdateTree(t *testing.T) {
	cases := []struct {
		name         string
		input        UpdateBuilder
		expectedNode exprNode
		err          updateErrorMode
	}{
		{
			name:  "set update",
			input: Set(Name("foo"), Value(5)),
			expectedNode: exprNode{
				children: []exprNode{
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c = $c",
							},
						},
						fmtExpr: "$c",
					},
				},
				fmtExpr: "SET $c\n",
			},
		},
		{
			name:  "remove update",
			input: Remove(Name("foo")),
			expectedNode: exprNode{
				children: []exprNode{
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
								},
								fmtExpr: "$c",
							},
						},
						fmtExpr: "$c",
					},
				},
				fmtExpr: "REMOVE $c\n",
			},
		},
		{
			name:  "add update",
			input: Add(Name("foo"), Value(5)),
			expectedNode: exprNode{
				children: []exprNode{
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c $c",
							},
						},
						fmtExpr: "$c",
					},
				},
				fmtExpr: "ADD $c\n",
			},
		},
		{
			name:  "delete update",
			input: Delete(Name("foo"), Value(5)),
			expectedNode: exprNode{
				children: []exprNode{
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c $c",
							},
						},
						fmtExpr: "$c",
					},
				},
				fmtExpr: "DELETE $c\n",
			},
		},
		{
			name:  "multiple sets",
			input: Set(Name("foo"), Value(5)).Set(Name("bar"), Value(6)).Set(Name("baz"), Name("qux")),
			expectedNode: exprNode{
				fmtExpr: "SET $c\n",
				children: []exprNode{
					{
						fmtExpr: "$c, $c, $c",
						children: []exprNode{
							{
								fmtExpr: "$c = $c",
								children: []exprNode{
									{
										fmtExpr: "$n",
										names:   []string{"foo"},
									},
									{
										fmtExpr: "$v",
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
									},
								},
							},
							{
								fmtExpr: "$c = $c",
								children: []exprNode{
									{
										fmtExpr: "$n",
										names:   []string{"bar"},
									},
									{
										fmtExpr: "$v",
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "6"},
										},
									},
								},
							},
							{
								fmtExpr: "$c = $c",
								children: []exprNode{
									{
										fmtExpr: "$n",
										names:   []string{"baz"},
									},
									{
										fmtExpr: "$n",
										names:   []string{"qux"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "compound update",
			input: Add(Name("foo"), Value(5)).Set(Name("foo"), Value(5)).Delete(Name("foo"), Value(5)).Remove(Name("foo")),
			expectedNode: exprNode{
				children: []exprNode{
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c $c",
							},
						},
						fmtExpr: "$c",
					},
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c $c",
							},
						},
						fmtExpr: "$c",
					},
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
								},
								fmtExpr: "$c",
							},
						},
						fmtExpr: "$c",
					},
					{
						children: []exprNode{
							{
								children: []exprNode{
									{
										names:   []string{"foo"},
										fmtExpr: "$n",
									},
									{
										values: []types.AttributeValue{
											&types.AttributeValueMemberN{Value: "5"},
										},
										fmtExpr: "$v",
									},
								},
								fmtExpr: "$c = $c",
							},
						},
						fmtExpr: "$c",
					},
				},
				fmtExpr: "ADD $c\nDELETE $c\nREMOVE $c\nSET $c\n",
			},
		},
		{
			name:  "empty UpdateBuilder",
			input: UpdateBuilder{},
			err:   unsetUpdate,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.buildTree()
			if c.err != noUpdateError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expectedNode, actual; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestSetValueBuilder(t *testing.T) {
	cases := []struct {
		name     string
		input    SetValueBuilder
		expected exprNode
		err      updateErrorMode
	}{
		{
			name:  "name plus name",
			input: Name("foo").Plus(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "$c + $c",
			},
		},
		{
			name:  "name minus name",
			input: Name("foo").Minus(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "$c - $c",
			},
		},
		{
			name:  "list append name and name",
			input: Name("foo").ListAppend(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "list_append($c, $c)",
			},
		},
		{
			name:  "if not exists name and name",
			input: Name("foo").IfNotExists(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						names:   []string{"foo"},
						fmtExpr: "$n",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "if_not_exists($c, $c)",
			},
		},
		{
			name:  "value plus name",
			input: Value(5).Plus(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						values: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
						},
						fmtExpr: "$v",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "$c + $c",
			},
		},
		{
			name:  "value minus name",
			input: Value(5).Minus(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						values: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
						},
						fmtExpr: "$v",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "$c - $c",
			},
		},
		{
			name:  "list append list and name",
			input: Value([]int{1, 2, 3}).ListAppend(Name("bar")),
			expected: exprNode{
				children: []exprNode{
					{
						values: []types.AttributeValue{
							&types.AttributeValueMemberL{
								Value: []types.AttributeValue{
									&types.AttributeValueMemberN{Value: "1"},
									&types.AttributeValueMemberN{Value: "2"},
									&types.AttributeValueMemberN{Value: "3"},
								},
							},
						},
						fmtExpr: "$v",
					},
					{
						names:   []string{"bar"},
						fmtExpr: "$n",
					},
				},
				fmtExpr: "list_append($c, $c)",
			},
		},
		{
			name:  "unset SetValueBuilder",
			input: SetValueBuilder{},
			err:   unsetSetValue,
		},
		{
			name:  "invalid operand error",
			input: Name("").Plus(Name("foo")),
			err:   invalidUpdateOperand,
		},
		{
			name:  "invalid operand error",
			input: Name("foo").Plus(Name("")),
			err:   invalidUpdateOperand,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := c.input.BuildOperand()
			if c.err != noUpdateError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, actual.exprNode; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}

func TestUpdateBuildChildNodes(t *testing.T) {
	cases := []struct {
		name     string
		input    []operationBuilder
		expected exprNode
		err      updateErrorMode
	}{
		{
			name: "set operand builder",
			input: []operationBuilder{
				{
					mode: setOperation,
					name: NameBuilder{
						names: []string{"foo"},
					},
					value: ValueBuilder{
						value: 5,
					},
				},
				{
					mode: setOperation,
					name: NameBuilder{
						names: []string{"bar"},
					},
					value: ValueBuilder{
						value: 6,
					},
				},
				{
					mode: setOperation,
					name: NameBuilder{
						names: []string{"baz"},
					},
					value: ValueBuilder{
						value: 7,
					},
				},
				{
					mode: setOperation,
					name: NameBuilder{
						names: []string{"qux"},
					},
					value: ValueBuilder{
						value: 8,
					},
				},
			},
			expected: exprNode{
				fmtExpr: "$c, $c, $c, $c",
				children: []exprNode{
					{
						fmtExpr: "$c = $c",
						children: []exprNode{
							{
								fmtExpr: "$n",
								names:   []string{"foo"},
							},
							{
								fmtExpr: "$v",
								values: []types.AttributeValue{
									&types.AttributeValueMemberN{Value: "5"},
								},
							},
						},
					},
					{
						fmtExpr: "$c = $c",
						children: []exprNode{
							{
								fmtExpr: "$n",
								names:   []string{"bar"},
							},
							{
								fmtExpr: "$v",
								values: []types.AttributeValue{
									&types.AttributeValueMemberN{Value: "6"},
								},
							},
						},
					},
					{
						fmtExpr: "$c = $c",
						children: []exprNode{
							{
								fmtExpr: "$n",
								names:   []string{"baz"},
							},
							{
								fmtExpr: "$v",
								values: []types.AttributeValue{
									&types.AttributeValueMemberN{Value: "7"},
								},
							},
						},
					},
					{
						fmtExpr: "$c = $c",
						children: []exprNode{
							{
								fmtExpr: "$n",
								names:   []string{"qux"},
							},
							{
								fmtExpr: "$v",
								values: []types.AttributeValue{
									&types.AttributeValueMemberN{Value: "8"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "empty operationBuilder list",
			input: []operationBuilder{},
			err:   emptyOperationBuilderList,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := buildChildNodes(c.input)
			if c.err != noUpdateError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else {
					if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
						t.Errorf("expect %q error message to be in %q", e, a)
					}
				}
			} else {
				if err != nil {
					t.Errorf("expect no error, got unexpected Error %q", err)
				}

				if e, a := c.expected, actual; !reflect.DeepEqual(a, e) {
					t.Errorf("expect %v, got %v", e, a)
				}
			}
		})
	}
}
