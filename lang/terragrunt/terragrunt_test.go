package terragrunt

import (
	"github.com/alecthomas/repr"
	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
	"github.com/sirupsen/logrus"
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
	ast, err := ParseHCLFile("account.hcl", contents)
	require.NoError(t, err)

	repr.Println(ast)
	//idx := indexAST(ast)
	//repr.Println(idx)
}

func TestIndexScopes(t *testing.T) {
	contents, err := os.ReadFile("testdata/root.hcl")
	require.NoError(t, err)
	doc, err := ParseHCLFile("terragrunt.hcl", contents)
	require.NoError(t, err)

	str := repr.String(doc.RootScopes, repr.Indent("  "))
	require.Equal(t, `map[string]map[string]*terragrunt.IndexedNode{
  "include": map[string]*terragrunt.IndexedNode{
    "root": [10:1-12:2] *hclsyntax.Block,
  },
  "local": map[string]*terragrunt.IndexedNode{
    "environment_vars": [2:3-2:79] *hclsyntax.Attribute,
    "meaning": [7:3-7:15] *hclsyntax.Attribute,
    "region": [3:3-3:23] *hclsyntax.Attribute,
  },
}`, str,
	)
}

func TestTerragruntParse(t *testing.T) {
	path := "../testdata/terragrunt-example/prod/us-east-1/prod/webserver-cluster/terragrunt.hcl"
	//path := "testdata/root.hcl"
	contents, err := os.ReadFile(path)
	require.NoError(t, err)
	cfg, err := config.ParseConfigString(string(contents), &options.TerragruntOptions{
		TerragruntConfigPath:         path,
		OriginalTerragruntConfigPath: path,
		MaxFoldersToCheck:            options.DefaultMaxFoldersToCheck,
		Logger:                       logrus.NewEntry(logrus.New()),
	}, nil, path, nil)
	require.NoError(t, err)
	repr.Println(cfg)
}
