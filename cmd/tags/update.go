/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	tagsClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tagUpdateCmd = &cobra.Command{
	Use:   "tag [tagId]",
	Short: "Updates a specific tag",
	Long:  `Updates the specific tag by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateTagRequest := *tagsClient.NewUpdateTagRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if TagName != "" {
				updateTagRequest.Name = &TagName
			}
			if TagColor != "" {
				updateTagRequest.Color = &TagColor
			}
		default:
			// from file / stdin
			var requestFromFile tagsClient.UpdateTagRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateTagRequest)
		}

		resp, httpResp, err := client.ApiClient().TagsApi.UpdateTag(context.Background(), tagId).UpdateTagRequest(updateTagRequest).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while updating tag")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify tagId")
		}
		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified tagId %v is not valid", args[0]))
		}
		tagId = tagId64
		if viper.GetString("name") != "" {
			TagName = viper.GetString("name")
		}
		if viper.GetString("color") != "" {
			TagColor = viper.GetString("color")
		}
		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(tagUpdateCmd)

	tagUpdateCmd.Flags().StringVarP(&TagName, "name", "n", "", `name of the tag`)
	viper.BindPFlag("name", tagUpdateCmd.Flags().Lookup("name"))
	viper.SetDefault("name", "")
	tagUpdateCmd.Flags().StringVarP(&TagColor, "color", "c", "", `color of the tag`)
	viper.BindPFlag("color", tagUpdateCmd.Flags().Lookup("color"))
	viper.SetDefault("color", "")
}
