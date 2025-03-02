package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/brucelay/oapi-viewer/stoplightelements"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var (
	specPath      string
	openInBrowser bool
	rootCmd       = &cobra.Command{
		Use:   "oapi-viewer",
		Short: "oapi-viewer is a tool for viewing OpenAPI specifications in the browser",
		Run: func(cmd *cobra.Command, args []string) {
			port := 7626
			var l net.Listener
			for l == nil {
				var err error
				l, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
				if err != nil {
					if !strings.Contains(err.Error(), "address already in use") {
						log.Fatal("Failed to start HTTP server: ", err)
					}
					port++
				}
			}

			filename := filepath.Base(specPath)

			http.HandleFunc("/"+filename, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Access-Control-Allow-Origin", "*")
				http.ServeFile(w, r, specPath)
			})

			html := stoplightelements.HtmlFromSpec(
				specPath,
				fmt.Sprintf("http://localhost:%d/%s", port, filename),
			)

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(html)
			})

			if openInBrowser {
				browser.OpenURL(fmt.Sprintf("http://localhost:%d", port))
			}

			fmt.Printf("Listening at http://localhost:%d", port)
			if err := http.Serve(l, nil); err != nil {
				log.Fatal("Failed to serve using listener: ", err)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&specPath, "path", "p", "", "path to OpenAPI specification")
	rootCmd.MarkFlagRequired("path")

	rootCmd.Flags().BoolVarP(&openInBrowser, "open", "o", false, "open in browser")
}
