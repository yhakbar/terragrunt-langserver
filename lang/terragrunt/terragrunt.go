package terragrunt

import (
	"bytes"
	"fmt"
	phcl "github.com/alecthomas/hcl"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"reflect"
)

func ParseHCL(fileName string, contents []byte) (file *hcl.File, err error) {

	defer func() {
		if recovered := recover(); recovered != nil {
			err = errors.Errorf("panic while parsing %s: %+v", fileName, recovered)
		}
	}()

	file, diags := hclsyntax.ParseConfig(contents, fileName, hcl.Pos{Byte: 0, Line: 1, Column: 1})
	if diags != nil && diags.HasErrors() {
		return nil, diags
	}

	return file, nil
}

type IndexedNode struct {
	SrcRange *hcl.Range
	Parent   *IndexedNode
	Node     hclsyntax.Node
}

func (n *IndexedNode) GoString() string {
	r := n.SrcRange
	return fmt.Sprintf("[%d:%d-%d:%d] %s", r.Start.Line, r.Start.Column, r.End.Line, r.End.Column, reflect.TypeOf(n.Node))
}

func (n *IndexedNode) String() string {
	return n.GoString()
}

type Document struct {
	Index  NodeIndex
	Locals map[string]*IndexedNode
}

func (d *Document) FindNodeAt(pos hcl.Pos) *IndexedNode {
	nodes, ok := d.Index[pos.Line]
	if !ok {
		return nil
	}
	var closest *IndexedNode
	for _, node := range nodes {
		if node.SrcRange.Start.Column < pos.Column {
			closest = node
		}
	}
	return closest
}

type NodeIndex map[int][]*IndexedNode

type nodeIndexBuilder struct {
	stack  []*IndexedNode
	index  NodeIndex
	locals map[string]*IndexedNode
}

func newTokenIndexBuilider() *nodeIndexBuilder {
	return &nodeIndexBuilder{
		index:  make(map[int][]*IndexedNode),
		locals: make(map[string]*IndexedNode),
	}
}

func (w *nodeIndexBuilder) Enter(node hclsyntax.Node) hcl.Diagnostics {
	var parent *IndexedNode
	if len(w.stack) > 0 {
		parent = w.stack[len(w.stack)-1]
	}
	line := node.Range().Start.Line
	inode := &IndexedNode{
		SrcRange: node.Range().Ptr(),
		Parent:   parent,
		Node:     node,
	}
	w.stack = append(w.stack, inode)
	w.index[line] = append(w.index[line], inode)
	if IsLocalAttribute(inode) {
		if attr, ok := node.(*hclsyntax.Attribute); ok {
			w.locals[attr.Name] = inode
		}
	}
	return nil
}

func IsLocalAttribute(node *IndexedNode) bool {
	if node.Parent == nil || node.Parent.Parent == nil || node.Parent.Parent.Parent == nil {
		return false
	}
	if _, ok := node.Parent.Node.(hclsyntax.Attributes); !ok {
		return false
	}
	if _, ok := node.Parent.Parent.Node.(*hclsyntax.Body); !ok {
		return false
	}
	return IsLocalBlock(node.Parent.Parent.Parent.Node)
}

func IsLocalBlock(node hclsyntax.Node) bool {
	block, ok := node.(*hclsyntax.Block)
	return ok && block.Type == "locals"
}

func IsAttributeNode(node hclsyntax.Node) bool {
	_, ok := node.(*hclsyntax.Attribute)
	return ok
}

func (w *nodeIndexBuilder) Exit(node hclsyntax.Node) hcl.Diagnostics {
	w.stack = w.stack[0 : len(w.stack)-1]
	return nil
}

var _ hclsyntax.Walker = &nodeIndexBuilder{}

func IndexAST(ast *hcl.File) *Document {
	body := ast.Body.(*hclsyntax.Body)
	builder := newTokenIndexBuilider()
	_ = hclsyntax.Walk(body, builder)
	return &Document{
		Index:  builder.index,
		Locals: builder.locals,
	}
}

func ParseHCLParticiple(fileName string, contents []byte) (*phcl.AST, error) {
	return phcl.Parse(bytes.NewReader(contents))
}
