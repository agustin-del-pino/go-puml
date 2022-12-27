package cmd

import (
	"fmt"
	"github.com/agustin-del-pino/gopuml/pkg/client"
	"github.com/spf13/cobra"
	"os"
)

var renderType string
var url = "https://www.plantuml.com/plantuml"

var renderCmd = &cobra.Command{
	Use:        "render",
	Short:      "render a diagram",
	ArgAliases: []string{"srcfile", "destfile"},
	Run: func(cmd *cobra.Command, args []string) {
		if t := client.PlantUMLRender(renderType); t == "" {
			panic(fmt.Errorf("invalid render type: %s", renderType))
		} else {
			if c, err := client.NewPlantUMLClient(url); err != nil {
				panic(err)
			} else {
				if u, b, err := c.RenderFile(t, args[0]); err != nil {
					panic(err)
				} else {
					println(u)
					p := fmt.Sprintf("%s.%s", args[1], renderType)
					if err := os.WriteFile(p, b, 0777); err != nil {
						panic(err)
					} else {
						fmt.Printf("the diagram was saved at: %s\n", p)
					}
				}
			}
		}
	},
}

func init() {
	renderCmd.Flags().StringVarP(&renderType, "render-type", "t", string(client.PNG), "specifies the type of diagram render")
	renderCmd.Flags().StringVarP(&url, "url", "u", url, "specifies the url to use")
	cli.AddCommand(renderCmd)
}
