package terragrunt

import (
	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParseHCLParticiple(t *testing.T) {
	t.Skip()
	contents, err := os.ReadFile("../testdata/terragrunt-example/terragrunt.hcl")
	require.NoError(t, err)
	ast, err := ParseHCLParticiple("account.hcl", contents)
	require.NoError(t, err)
	repr.Println(ast)
}

func TestParseHCL(t *testing.T) {
	contents, err := os.ReadFile("../testdata/terragrunt-example/terragrunt.hcl")
	require.NoError(t, err)
	ast, err := ParseHCL("account.hcl", contents)
	require.NoError(t, err)

	idx := IndexAST(ast)
	repr.Println(idx)
}
