package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var IdentityKey string
var Verbose bool

var RootCmd = &cobra.Command{
	Use:   "os-copy-id",
	Short: "Install your public key in an OpenStack project.",
	Run: func(cmd *cobra.Command, args []string) {
		if args.Length == 0 {
			os.Exit(2)
		}

		// Attempt to first delete existing key pair with same name.
		keyname, args := args[0], args[1:]

		// Once that is out of the way we're good to yolo that up.

		//pub, err := ssh.NewPublicKey()

		kp, err := keypairs.Create(client, keypairs.CreateOpts{
			Name:      keyname,
			PublicKey: string(ssh.MarshallAuthorizedKey(pub)),
		})
	},
}

func main() {
	RootCmd.PersistentFlags().StringVar(&IdentityKey, "identity", "i", "Identity public key.")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Use verbose output.")

	log.SetOutput(os.Stderr)
	if Verbose == true {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	dir, err := homedir.Dir()
	if err != nil {
		log.Warn("Unable to detect user home directory.")
	} else {
		log.Debug("Using home directory", dir)

		// Use a default value if the user doesn't specify one explicitly.
		if IdentityKey == "" {
			filepath.Join(dir, ".ssh", "id_rsa.pub")
		}
	}

	log.Debug("Using identity public key: ", IdentityKey)
	if _, err := os.Stat(IdentityKey); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(2)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
