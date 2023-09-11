package terragrunt

import (
	"bytes"
	"fmt"
	phcl "github.com/alecthomas/hcl"
	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log/slog"
	"os"
	"reflect"
	"runtime/debug"
)

func ParseHCLFile(fileName string, contents []byte) (file *IndexedAST, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = errors.Errorf("panic while parsing %s: %+v\n%s", fileName, recovered, string(debug.Stack()))
		}
	}()

	hclFile, diags := hclsyntax.ParseConfig(contents, fileName, hcl.Pos{Byte: 0, Line: 1, Column: 1})
	if diags != nil && diags.HasErrors() {
		return nil, diags
	}

	return indexAST(hclFile), nil
}

type IndexedNode struct {
	Parent *IndexedNode
	hclsyntax.Node
}

func (n *IndexedNode) GoString() string {
	r := n.Range()
	return fmt.Sprintf("[%d:%d-%d:%d] %s", r.Start.Line, r.Start.Column, r.End.Line, r.End.Column, reflect.TypeOf(n.Node))
}

func (n *IndexedNode) String() string {
	return n.GoString()
}

type IndexedAST struct {
	HCLFile    *hcl.File
	Index      NodeIndex
	RootScopes RootScopes
}

func (d *IndexedAST) FindNodeAt(pos hcl.Pos) *IndexedNode {
	// Iterate backwards to find a node that starts before the position
	nodes, ok := d.Index[pos.Line]
	if !ok {
		return nil
	}
	var closest *IndexedNode
	// First try finding a matching node on the same line
	for _, node := range nodes {
		if node.Range().Start.Column <= pos.Column {
			closest = node
		}
	}
	if closest == nil {
		// Iterate backwards by line
		for i := pos.Line - 1; i >= 1; i-- {
			slog.Info("Check", "line", i)
			nodes, ok = d.Index[i]
			if !ok || len(nodes) == 0 {
				continue
			}
			closest = nodes[len(nodes)-1]
			break
		}
	}
	if closest == nil {
		return nil
	}
	// Navigate up the AST to find the first node that contains the position.
	node := closest
	for node != nil {
		end := node.Range().End
		if end.Line > pos.Line || end.Line == pos.Line && end.Column > pos.Column {
			return node
		}
		node = node.Parent
	}
	return nil
}

type RootScopes map[string]Scope

func (s RootScopes) SetNode(scope string, node *IndexedNode) {
	rootScope, ok := s[scope]
	if !ok {
		rootScope = make(Scope)
		s[scope] = rootScope
	}
	switch n := node.Node.(type) {
	case *hclsyntax.Attribute:
		rootScope[n.Name] = node
	case *hclsyntax.Block:
		rootScope[n.Labels[0]] = node
	default:
		panic("invalid node type " + reflect.TypeOf(node.Node).String())
	}
}

type Scope map[string]*IndexedNode

type NodeIndex map[int][]*IndexedNode

type nodeIndexBuilder struct {
	stack      []*IndexedNode
	index      NodeIndex
	rootScopes RootScopes
}

func newNodeIndexBuilider() *nodeIndexBuilder {
	return &nodeIndexBuilder{
		index:      make(map[int][]*IndexedNode),
		rootScopes: make(RootScopes),
	}
}

func (w *nodeIndexBuilder) Enter(node hclsyntax.Node) hcl.Diagnostics {
	var parent *IndexedNode
	if len(w.stack) > 0 {
		parent = w.stack[len(w.stack)-1]
	}
	line := node.Range().Start.Line
	inode := &IndexedNode{
		Parent: parent,
		Node:   node,
	}
	w.stack = append(w.stack, inode)
	w.index[line] = append(w.index[line], inode)
	if IsLocalAttribute(inode) {
		w.rootScopes.SetNode("local", inode)
	}
	if block, ok := node.(*hclsyntax.Block); ok && block.Type == "include" && len(block.Labels) > 0 {
		w.rootScopes.SetNode("include", inode)
	}
	return nil
}

func (w *nodeIndexBuilder) Exit(node hclsyntax.Node) hcl.Diagnostics {
	w.stack = w.stack[0 : len(w.stack)-1]
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

var _ hclsyntax.Walker = &nodeIndexBuilder{}

func indexAST(ast *hcl.File) *IndexedAST {
	body := ast.Body.(*hclsyntax.Body)
	builder := newNodeIndexBuilider()
	_ = hclsyntax.Walk(body, builder)
	return &IndexedAST{
		Index:      builder.index,
		RootScopes: builder.rootScopes,
		HCLFile:    ast,
	}
}

func ParseHCLParticiple(fileName string, contents []byte) (*phcl.AST, error) {
	return phcl.Parse(bytes.NewReader(contents))
}

func Evaluate(filePath string, contents []byte) (*config.TerragruntConfig, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return config.ParseConfigString(string(contents), &options.TerragruntOptions{
		TerragruntConfigPath:         filePath,
		OriginalTerragruntConfigPath: filePath,
		MaxFoldersToCheck:            options.DefaultMaxFoldersToCheck,
		Logger:                       logrus.NewEntry(logrus.New()),
	}, nil, filePath, nil)
}
