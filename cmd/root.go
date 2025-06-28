package cmd

import (
	"context"
	"io/fs"

	"github.com/Fuabioo/altalune/internal/server"
	cliutils "github.com/Fuabioo/altalune/pkg/cliutls"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var assets fs.FS

var rootCmd = &cobra.Command{
	Use:   "altalune",
	Short: "Navigate your epics among the stars",
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		cliutils.LoadEnvFiles()
	},
	Run: func(cmd *cobra.Command, args []string) {

		// Start HTTP server
		srv, err := server.NewServer(
			server.ServerAssets(assets),
			server.ServerAddress(
				viper.GetString("server-host"),
				viper.GetUint("server-port"),
			),
			server.ServerJira(
				viper.GetString("workspace"),
				viper.GetString("email"),
				viper.GetString("token"),
			),
			server.ServerSuperDebug(viper.GetBool("super-debug")),
		)
		if err != nil {
			log.Fatal(err)
		}
		defer srv.Close()

		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func initConfig() {
	viper.SetEnvPrefix("JIRA_EPIC")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().String("host", "", "JIRA host")
	rootCmd.Flags().String("email", "", "JIRA email")
	rootCmd.Flags().String("token", "", "JIRA token")
	rootCmd.Flags().String("server-host", "0.0.0.0", "Server host")
	rootCmd.Flags().Uint("server-port", 3002, "Server port")
	rootCmd.Flags().Bool("verbose", false, "Verbose logging")
	rootCmd.Flags().Bool("super-debug", false, "Super debug logging")

	viper.BindPFlag("host", rootCmd.Flags().Lookup("host"))
	viper.BindPFlag("email", rootCmd.Flags().Lookup("email"))
	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	viper.BindPFlag("server-host", rootCmd.Flags().Lookup("server-host"))
	viper.BindPFlag("server-port", rootCmd.Flags().Lookup("server-port"))
	viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	viper.BindPFlag("super-debug", rootCmd.Flags().Lookup("super-debug"))
}

func Execute(staticFiles fs.FS) {
	assets = getFileSystem(staticFiles)
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		log.Fatal(err)
	}
}

func getFileSystem(assets fs.FS) fs.FS {
	log.Debug("Using embed mode for the frontend static assets 📦")
	fsys, err := fs.Sub(assets, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return fsys
}
