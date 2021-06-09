package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:     "add  [ip] [fqdn] [comment]",
		Aliases: []string{"a", "ad"},
		Short:   "Add file data",
		Long: `Add hosts data to file.
[ip] and [fqdn] are required and positional arguments.
[comment] can be omitted`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 || len(args) > 3 {
				return errors.New("invalid number of arguments")
			}

			return nil
		},
		ValidArgs: []string{"ip", "fqdn", "comment"},
		Run: func(cmd *cobra.Command, args []string) {
			var fqdn, comment string
			var hfline *libhosty.HostsFileLine
			hfl := make([]*libhosty.HostsFileLine, 0)
			var err error

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			ip := net.ParseIP(args[0])
			fqdn = string(args[1])
			if len(args) >= 3 {
				comment = string(args[2])
			} else {
				comment = ""
			}

			if ip != nil {
				_, hfline, err = hf.AddHostsFileLine(ip.String(), fqdn, comment)
				cobra.CheckErr(err)
				hfl = append(hfl, hfline)

				if backupFile {
					err = Backup()
					cobra.CheckErr(err)
				}
				err = hf.SaveHostsFile()
				cobra.CheckErr(err)

				PrintOutput("add", details, hfl)
				return
			}

			fmt.Printf("invalid ip address %s as argument", args[0])
		},
	}
)
