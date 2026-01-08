package cmd

import "github.com/spf13/cobra"

var getPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCube(cmd, args)
	},
}

var getPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCube(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getPodsCmd)
	getCmd.AddCommand(getPodCmd)

	getPodsCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getPodsCmd.Flags().Bool("watch", false, "Auto-rotate cube")

	getPodCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getPodCmd.Flags().Bool("watch", false, "Auto-rotate cube")

}
