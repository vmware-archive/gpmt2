/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/

// cobra command line

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCollectorCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GPDB Version number",
	Long:  `Greenplum Magic Tool version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version ALPHA 1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "gpmt",
	Short: "GPMT - diagnostic and data collection for Greemplum Database",
	Long: `Greenplum Magic Tool is a collection of diagnostic and data collection tools
				to assist in troubleshooting issues with Greenplum Database.
				Documentation and development information is available at
				https://github.com/pivotal-gss/gpmt2`,
	Run: func(cmd *cobra.Command, args []string) {
		// probably just help
	},
}

var logCollectorCmd = &cobra.Command{
	Use:   "gp_log_collector",
	Short: "easy log collection",
	Long: `gp_log_collector is used to automate Greenplum database log collection. 
				Run without options, gp_log_collector will gather today's master
				and standby logs`,
	Run: func(cmd *cobra.Command, args []string) {
		// log collect
		fmt.Println("I'll be a log collector one day")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
