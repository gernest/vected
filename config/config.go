package config

import (
	"go/build"
	"path/filepath"

	"github.com/urfave/cli"
)

// Config contains configuration testails about the test running environment.
type Config struct {
	// Information about the package.
	Info *build.Package
	// URL where the test runner weservice is running.
	ServerURL string

	// This is absolute path to the root of the package being tested,
	Root string

	// The directory where test files stay.
	TestDirName string

	TestPath string

	TestUnitDir string
	TestUnitPkg string

	OutputDirName string

	OutputPath string

	OutputMainPkg string

	Build bool
}

// FLags returns configuration flags.
func FLags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "server_url",
			EnvVar: "PEST_SERVER_URL",
			Value:  "http://localhost:1955",
		},
		cli.StringFlag{
			Name:  "root",
			Usage: "the root path of the package",
		},
		cli.StringFlag{
			Name:  "test_dir",
			Usage: "relative path to the tests directory",
			Value: "test",
		},
		cli.StringFlag{
			Name:  "output_dir",
			Usage: "relative path to the generated tests directory",
			Value: "promtest",
		},
		cli.BoolFlag{
			Name: "build",
		},
	}
}

// Load returns *Config instance with values populated from ctx.
func Load(ctx *cli.Context) (*Config, error) {
	c := &Config{
		ServerURL:     ctx.String("server_url"),
		Root:          ctx.String("root"),
		TestDirName:   ctx.String("test_dir"),
		OutputDirName: ctx.String("output_dir"),
		Build:         ctx.Bool("build"),
	}
	if !filepath.IsAbs(c.Root) {
		p, err := filepath.Abs(c.Root)
		if err != nil {
			return nil, err
		}
		c.Root = p
	}
	pkg, err := build.ImportDir(c.Root, build.FindOnly)
	if err != nil {
		return nil, err
	}
	c.Info = pkg
	c.TestPath = filepath.Join(c.Info.Dir, c.TestDirName)
	c.OutputPath = filepath.Join(c.Info.Dir, c.OutputDirName)
	c.TestUnitDir = filepath.Join(c.OutputPath, c.TestDirName)
	c.TestUnitPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName, c.TestDirName)
	c.OutputMainPkg = filepath.Join(c.Info.ImportPath, c.OutputDirName)
	return c, nil
}

func (c *Config) GetOutDir() string {
	return filepath.Join(c.Info.Dir, c.OutputDirName)
}

func (c *Config) GetTestDirName() string {
	return filepath.Join(c.Info.Dir, c.TestDirName)
}