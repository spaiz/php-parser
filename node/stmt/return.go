package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Return struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewReturn(expr node.Node) node.Node {
	return Return{
		"Return",
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Return) Name() string {
	return "Return"
}

func (n Return) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Return) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Return) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Return) Position() *node.Position {
	return n.position
}

func (n Return) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
