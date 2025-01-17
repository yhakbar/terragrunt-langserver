package terragrunt

import (
	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParseHCL(t *testing.T) {
	contents, err := os.ReadFile("../testdata/terragrunt-example/terragrunt.hcl")
	require.NoError(t, err)
	ast, err := ParseHCLFile("account.hcl", contents)
	require.NoError(t, err, ast)
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
	require.NoError(t, err, cfg)
}
