package worker

import "github.com/consolelabs/indexer-api/pkg/logger"

type ICommand interface {
	SetName(name string) ICommand
	SetUsage(usage string) ICommand
	SetAliases(aliases ...string) ICommand
	SetRun(run func(args ...string) error) ICommand
	Name() string
	Usage() string
	Aliases() []string
	Run() error
}

type cmd struct {
	name    string
	aliases []string
	usage   string
	run     func(args ...string) error
	log     logger.Logger
}

func NewCmd(log logger.Logger) ICommand {
	return &cmd{log: log}
}

func (c *cmd) SetName(name string) ICommand {
	c.name = name
	return c
}

func (c *cmd) SetUsage(usage string) ICommand {
	c.usage = usage
	return c
}

func (c *cmd) SetAliases(aliases ...string) ICommand {
	c.aliases = aliases
	return c
}

func (c *cmd) SetRun(run func(args ...string) error) ICommand {
	c.run = run
	return c
}

func (c *cmd) Name() string {
	return c.name
}

func (c *cmd) Usage() string {
	return c.usage
}

func (c *cmd) Aliases() []string {
	return c.aliases
}

func (c *cmd) Run() error {
	return c.run()
}
