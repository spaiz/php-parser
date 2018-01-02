package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanAnd struct {
	BinaryOp
}

func NewBooleanAnd(variable node.Node, expression node.Node) node.Node {
	return BooleanAnd{
		BinaryOp{
			"BinaryBooleanAnd",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n BooleanAnd) Name() string {
	return "BooleanAnd"
}

func (n BooleanAnd) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BooleanAnd) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BooleanAnd) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n BooleanAnd) Position() *node.Position {
	return n.position
}

func (n BooleanAnd) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BooleanAnd) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
