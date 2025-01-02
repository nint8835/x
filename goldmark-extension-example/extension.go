package goldmarkextensionexample

import (
	"fmt"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var ExprKind = ast.NewNodeKind("Expr")

type exprAstNode struct {
	ast.BaseInline

	exprString   string
	compiledExpr *vm.Program
}

func (e *exprAstNode) Kind() ast.NodeKind {
	return ExprKind
}

func (e *exprAstNode) Dump(source []byte, level int) {
	ast.DumpHelper(e, source, level, map[string]string{"exprString": e.exprString}, nil)
}

type exprParser struct{}

func (e *exprParser) Trigger() []byte {
	return []byte("`")
}

func (e *exprParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, _ := block.PeekLine()

	lineHead := 0
	inExpr := false
	exprString := ""

	for ; lineHead < len(line); lineHead++ {
		if line[lineHead] == '`' {
			if !inExpr && string(line[lineHead:lineHead+5]) == "`math" {
				inExpr = true
				lineHead += 5
			} else if inExpr {
				inExpr = false
				break
			}
		} else if inExpr {
			exprString += string(line[lineHead])
		}
	}

	if inExpr || exprString == "" {
		return nil
	}

	compiledExpr, err := expr.Compile(exprString)
	if err != nil {
		fmt.Printf("Failed to compile expression: %v\n", err)
		return nil
	}

	block.Advance(lineHead + 1)

	return &exprAstNode{exprString: exprString, compiledExpr: compiledExpr}
}

var _ parser.InlineParser = (*exprParser)(nil)

type exprRenderer struct{}

func (r *exprRenderer) render(
	w util.BufWriter, source []byte, n ast.Node, entering bool,
) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	node := n.(*exprAstNode)

	execResult, err := expr.Run(node.compiledExpr, nil)
	if err != nil {
		return ast.WalkStop, fmt.Errorf("failed to run expression: %w", err)
	}

	w.WriteString(fmt.Sprintf("<code>%s = %v</code>", node.exprString, execResult))

	return ast.WalkContinue, nil
}

func (r *exprRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ExprKind, r.render)
}

var _ renderer.NodeRenderer = (*exprRenderer)(nil)

type ExprExtension struct{}

func (e *ExprExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(util.Prioritized(&exprParser{}, -1000)))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(util.Prioritized(&exprRenderer{}, -1000)))
}

var _ goldmark.Extender = (*ExprExtension)(nil)

func New() goldmark.Extender {
	return &ExprExtension{}
}
