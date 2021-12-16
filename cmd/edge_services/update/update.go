package update

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aziontech/azion-cli/cmd/edge_services/requests"
	"github.com/aziontech/azion-cli/pkg/cmdutil"
	"github.com/aziontech/azion-cli/utils"
	sdk "github.com/aziontech/edgeservices-go-sdk"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	// listCmd represents the list command
	updateCmd := &cobra.Command{
		Use:           "update",
		Short:         "Updates parameters of an edge service",
		Long:          `Updates parameters of an edge service`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := utils.ConvertIdsToInt(args[0])
			if err != nil {
				return utils.ErrorConvertingIdArgumentToInt
			}

			client, err := requests.CreateClient(f, cmd)
			if err != nil {
				return err
			}

			if err := updateService(client, id[0], cmd, args); err != nil {
				return err
			}

			return nil
		},
	}
	updateCmd.Flags().String("name", "", "<EDGE_SERVICE_NAME>")
	updateCmd.Flags().String("active", "", "<true|false>")
	updateCmd.Flags().String("variables-file", "", `<VARIABLES_FILE_PATH>
The format accepted for variables definition is one <KEY>=<VALUE> per line`)

	return updateCmd
}

func updateService(client *sdk.APIClient, id int64, cmd *cobra.Command, args []string) error {
	c := context.Background()
	api := client.DefaultApi

	serviceRequest := sdk.UpdateServiceRequest{}

	nameHasChanged := cmd.Flags().Changed("name")
	if nameHasChanged {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		serviceRequest.SetName(name)
	}

	activeHasChanged := cmd.Flags().Changed("active")
	if activeHasChanged {
		activeStr, err := cmd.Flags().GetString("active")
		if err != nil {
			return err
		}

		active, err := strconv.ParseBool(activeStr)
		if err != nil {
			return utils.ErrorConvertingStringToBool
		}
		serviceRequest.SetActive(active)
	}

	variablesHasChanged := cmd.Flags().Changed("variables-file")
	if variablesHasChanged {
		variablesPath, err := cmd.Flags().GetString("variables-file")
		if err != nil {
			return utils.ErrorHandlingFile
		}

		file, err := os.Open(variablesPath)
		if err != nil {
			return utils.ErrorHandlingFile
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		reName := regexp.MustCompile("^([^= ]+) *= *(.*)")
		reValue := regexp.MustCompile("^([^= ]+) *= *(.+)")
		v := []sdk.Variable{}
		for scanner.Scan() {
			line := scanner.Text()
			entry := strings.Split(line, "=")
			if len(entry) < 2 {
				return utils.ErrorInvalidVariablesFileFormat
			}
			varName := reName.FindStringSubmatch(strings.Trim(line, " "))[1]
			varValue := reValue.FindStringSubmatch(strings.Trim(line, " "))[2]
			variable := sdk.NewVariable(varName, varValue)
			v = append(v, *variable)
		}
		serviceRequest.SetVariables(v)

		if err := scanner.Err(); err != nil {
			return err
		}

	}

	resp, httpResp, err := api.PatchService(c, id).UpdateServiceRequest(serviceRequest).Execute()
	if err != nil {
		if httpResp.StatusCode >= 500 {
			return utils.ErrorInternalServerError
		}

		return err
	}

	fmt.Printf("ID: %d\tName: %s \n", resp.Id, resp.Name)

	return nil
}
