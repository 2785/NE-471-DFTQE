package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/2785/n471-proj-carrot/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Loading Config...")
	if !(viper.IsSet("inputdir") && viper.IsSet("outputdir")) {
		return errors.New("input or output not set")
	}

	fmt.Println("  done")

	fmt.Printf("Start Loading Input Files...\n")
	inputdir := viper.GetString("inputdir")

	files, err := ioutil.ReadDir(inputdir)
	if err != nil {
		return fmt.Errorf("cannot read input directory: %w", err)
	}

	fmt.Printf("discovered %v files\n", len(files))

	simulations := make([]*model.Simulation, len(files))

	_ = simulations

	return nil
}
