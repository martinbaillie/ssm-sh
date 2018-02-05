package command

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var Cmd RootCommand

type RootCommand struct {
	List    ListCommand  `command:"list" alias:"ls" description:"List managed instances."`
	Shell   ShellCommand `command:"shell" alias:"sh" description:"Start an interactive shell."`
	Run     RunCommand   `command:"run" description:"Run a command on the targeted instances."`
	AwsOpts AwsOptions   `group:"AWS Options"`
}

type AwsOptions struct {
	Profile string `short:"p" long:"profile" description:"AWS Profile to use. (If you are not using Vaulted)."`
	Region  string `short:"r" long:"region" description:"Region to target." default:"eu-west-1"`
}

type SsmOptions struct {
	Targets   []string `short:"t" long:"target" description:"One or more instance ids to target," required:"1"`
	Frequency int      `short:"f" long:"frequency" description:"Polling frequency (millseconds to wait between requests)." default:"500"`
	Timeout   int      `short:"i" long:"timeout" description:"Seconds to wait for command result before timing out." default:"30"`
}

func newSession() (*session.Session, error) {
	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if Cmd.AwsOpts.Profile != "" {
		opts.Profile = Cmd.AwsOpts.Profile
	}
	sess, err := session.NewSessionWithOptions(opts)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
