package expression

import (
	"strings"
	"testing"

	"github.com/matthew188/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/matthew188/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// opeErrorMode will help with error cases and checking error types
type opeErrorMode string

const (
	noOperandError opeErrorMode = ""
	// unsetName error will occur if an empty string is passed into NameBuilder
	unsetName = "unset parameter: NameBuilder"
	// invalidName error will occur if a nested name has an empty intermediary
	// attribute name (i.e. foo.bar..baz)
	invalidName = "invalid parameter: NameBuilder"
	// unsetKey error will occur if an empty string is passed into KeyBuilder
	unsetKey = "unset parameter: KeyBuilder"
)

func TestBuildOperand(t *testing.T) {
	type mockStructValue struct {
		A string `dynamodbav:"ddbA" tagb:"TagB"`
		B string
	}

	cases := []struct {
		name     string
		input    OperandBuilder
		expected exprNode
		err      opeErrorMode
	}{
		{
			name:  "basic name",
			input: Name("foo"),
			expected: exprNode{
				names:   []string{"foo"},
				fmtExpr: "$n",
			},
		},
		{
			name:  "duplicate name name",
			input: Name("foo.foo"),
			expected: exprNode{
				names:   []string{"foo", "foo"},
				fmtExpr: "$n.$n",
			},
		},
		{
			name:  "struct value",
			input: ValueWithOptions(mockStructValue{A: "abc123", B: "efg456"}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
						"ddbA": &types.AttributeValueMemberS{Value: "abc123"},
						"B":    &types.AttributeValueMemberS{Value: "efg456"},
					}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name: "struct value with TagKey",
			input: ValueWithOptions(mockStructValue{A: "abc123", B: "efg456"},
				func(o *ValueBuilderOptions) {
					o.EncoderOptions = append(o.EncoderOptions, func(o *attributevalue.EncoderOptions) {
						o.TagKey = "tagb"
					})
				}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
						"TagB": &types.AttributeValueMemberS{Value: "abc123"},
						"B":    &types.AttributeValueMemberS{Value: "efg456"},
					}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "basic value",
			input: Value(5),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "5"},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValue as pointer",
			input: Value(&types.AttributeValueMemberN{Value: "5"}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "5"},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberS as value",
			input: Value(types.AttributeValueMemberS{Value: "5"}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: "5"},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberN as value",
			input: Value(types.AttributeValueMemberN{Value: "5"}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "5"},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberB as value",
			input: Value(types.AttributeValueMemberB{Value: []byte{5}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberB{Value: []byte{5}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberSS as value",
			input: Value(types.AttributeValueMemberSS{Value: []string{"5", "6"}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberSS{Value: []string{"5", "6"}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberNS as value",
			input: Value(types.AttributeValueMemberNS{Value: []string{"5", "6"}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberNS{Value: []string{"5", "6"}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberBS as value",
			input: Value(types.AttributeValueMemberBS{Value: [][]byte{{5}, {6}}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberBS{Value: [][]byte{{5}, {6}}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name: "types.AttributeValueMemberM as value",
			input: Value(types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
				"bar": &types.AttributeValueMemberS{Value: "baz"},
			}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
						"bar": &types.AttributeValueMemberS{Value: "baz"},
					}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name: "types.AttributeValueMemberL as value",
			input: Value(types.AttributeValueMemberL{Value: []types.AttributeValue{
				&types.AttributeValueMemberS{Value: "5"},
			}}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberL{Value: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: "5"},
					}},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberNULL as value",
			input: Value(types.AttributeValueMemberNULL{Value: true}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberNULL{Value: true},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "types.AttributeValueMemberBOOL as value",
			input: Value(types.AttributeValueMemberBOOL{Value: true}),
			expected: exprNode{
				values: []types.AttributeValue{
					&types.AttributeValueMemberBOOL{Value: true},
				},
				fmtExpr: "$v",
			},
		},
		{
			name:  "nested name",
			input: Name("foo.bar"),
			expected: exprNode{
				names:   []string{"foo", "bar"},
				fmtExpr: "$n.$n",
			},
		},
		{
			name:  "no split name",
			input: NameNoDotSplit("foo.bar"),
			expected: exprNode{
				names:   []string{"foo.bar"},
				fmtExpr: "$n",
			},
		},
		{
			name:  "nested name with index",
			input: Name("foo.bar[0].baz"),
			expected: exprNode{
				names:   []string{"foo", "bar", "baz"},
				fmtExpr: "$n.$n[0].$n",
			},
		},
		{
			name:  "no split name with index",
			input: NameNoDotSplit("foo.bar[0]"),
			expected: exprNode{
				names:   []string{"foo.bar"},
				fmtExpr: "$n[0]",
			},
		},
		{
			name:  "no split name append name",
			input: NameNoDotSplit("foo.bar").AppendName(Name("foo.bar")),
			expected: exprNode{
				names:   []string{"foo.bar", "foo", "bar"},
				fmtExpr: "$n.$n.$n",
			},
		},
		{
			name: "no split name append name with list index",
			input: NameNoDotSplit("foo.bar").
				AppendName(Name("foo.bar")).
				AppendName(Name("[0]")).
				AppendName(Name("abc123")),
			expected: exprNode{
				names:   []string{"foo.bar", "foo", "bar", "abc123"},
				fmtExpr: "$n.$n.$n[0].$n",
			},
		},
		{
			name:  "basic size",
			input: Name("foo").Size(),
			expected: exprNode{
				names:   []string{"foo"},
				fmtExpr: "size ($n)",
			},
		},
		{
			name:  "key",
			input: Key("foo"),
			expected: exprNode{
				names:   []string{"foo"},
				fmtExpr: "$n",
			},
		},
		{
			name:     "unset key error",
			input:    Key(""),
			expected: exprNode{},
			err:      unsetKey,
		},
		{
			name:     "empty name error",
			input:    Name(""),
			expected: exprNode{},
			err:      unsetName,
		},
		{
			name:     "invalid name",
			input:    Name("foo..bar"),
			expected: exprNode{},
			err:      invalidName,
		},
		{
			name:     "invalid index",
			input:    Name("[foo]"),
			expected: exprNode{},
			err:      invalidName,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			operand, err := c.input.BuildOperand()

			if c.err != noOperandError {
				if err == nil {
					t.Errorf("expect error %q, got no error", c.err)
				} else if e, a := string(c.err), err.Error(); !strings.Contains(a, e) {
					t.Errorf("expect %q error message to be in %q", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect no error, got unexpected Error %q", err)
			}

			cmpOptions := cmp.Options{
				cmp.AllowUnexported(exprNode{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberM{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberN{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberNS{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberBOOL{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberB{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberBS{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberL{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberS{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberSS{}),
				cmpopts.IgnoreUnexported(types.AttributeValueMemberNULL{}),
			}
			if diff := cmp.Diff(c.expected, operand.exprNode, cmpOptions...); diff != "" {
				t.Errorf("expect operand match\n%s", diff)
			}
		})
	}
}
