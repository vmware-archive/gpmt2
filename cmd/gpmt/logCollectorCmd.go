package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogCollectorOptions define the options/flag for the gp_log_collector command
type LogCollectorOptions struct {
	failedOnly bool
	freeSpace  int
	contentIds []string
	hostfile   string
	hostnames  []string
	startDate  string
	endDate    string
	noPrompt   bool
	workingDir string
	segmentDir string
	osOnly     bool
	standby    bool
}

// Sub Command: Log Collector
// This command line arguments helps to obtain the logs from the greenplum database
var logCollectorCmd = &cobra.Command{
	Use:   "gp_log_collector",
	Short: "easy log collection",
	Long: "\ngp_log_collector is used to automate Greenplum database log collection. \n" +
		"Run without options, gp_log_collector will gather today's master and standby logs",
	Run: func(cmd *cobra.Command, args []string) {
		// log collect
		fmt.Println("I'll be a log collector one day")

	},
}

// All the usage flags of the log collector
func flagsLogCollector() {
	logCollectorCmd.Flags().BoolVar(&lcOpts.failedOnly, "failed-segs", false, "Query gp_configuration_history for list of faulted content ids")
	logCollectorCmd.Flags().IntVar(&lcOpts.freeSpace, "free-space", 10, "default=10  Free space threshold which will abort log collection if reached")
	logCollectorCmd.Flags().StringArrayVar(&lcOpts.contentIds, "c", nil, "Space seperated list of content ids")
	logCollectorCmd.Flags().BoolVar(&lcOpts.noPrompt, "no-prompts", false, "Accept all prompts")
	logCollectorCmd.Flags().StringVarP(&lcOpts.hostfile, "hostfile", "f", "", "Read hostnames from a hostfile")
	logCollectorCmd.Flags().StringArrayVarP(&lcOpts.hostnames, "hostnames", "n", nil, "Space seperated list of hostnames")
	// FIXME: If date is empty string startDate and endDate it should default to current date
	logCollectorCmd.Flags().StringVar(&lcOpts.startDate, "start", "", "Start date for logs to collect (defaults to current date)")
	logCollectorCmd.Flags().StringVar(&lcOpts.endDate, "end", "", "End date for logs to collect (defaults to current date)")
	// FIXME: If workingDir is empty string it should default to cwd
	logCollectorCmd.Flags().StringVar(&lcOpts.workingDir, "dir", "", "Working directory (defaults to current directory)")
	// FIXME: If segmentDir is empty string it should default to /tmp
	logCollectorCmd.Flags().StringVar(&lcOpts.segmentDir, "segdir", "", "Segment temporary directory (defaults to /tmp)")
	logCollectorCmd.Flags().BoolVar(&lcOpts.osOnly, "os-only", false, "Only collect minimal infrastucture information")
	logCollectorCmd.Flags().BoolVar(&lcOpts.standby, "collect-standby", false, "Collect information from the standby master")
}
