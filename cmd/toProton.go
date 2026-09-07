package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/crossRT/escape-authy/helper"
	"github.com/crossRT/escape-authy/model"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// toProtonCmd represents the toProton command
var toProtonCmd = &cobra.Command{
	Use:   "toProton",
	Short: "Convert Authy decrypted_tokens.json to proton format",
	Run: func(cmd *cobra.Command, args []string) {
		decryptedTokensFilePath, _ = rootCmd.PersistentFlags().GetString("decrypted-tokens-file-path")

		file, err := os.ReadFile(decryptedTokensFilePath)
		if err != nil {
			panic(err)
		}

		// Parse the JSON
		var dtf model.DecryptedTokensFile
		err = json.Unmarshal(file, &dtf)
		if err != nil {
			panic(err)
		}

		var protonEntries []model.ProtonEntry
		for _, token := range dtf.DecryptedAuthenticatorTokens {

			name := strings.TrimSpace(token.Name)
			timer := strconv.Itoa(helper.GetTimerByDigits(int(token.Digits)))
			digits := strconv.Itoa(token.Digits)
			issuer := helper.DetermineIssuer(token)

			uri := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s&algorithm=SHA1&digits=%s&period=%s", name, token.DecryptedSeed, issuer, digits, timer)

			protonEntry := model.ProtonEntry{
				Id:   uuid.New().String(),
				Content: model.ProtonContent{
					Uri: uri,
					EntryType:   "Totp",
					Name: name,
				},
				Note: nil,
			}

			protonEntries = append(protonEntries, protonEntry)
		}

		protonOutput := model.ProtonBase{
			Version: 1,
			Entries: protonEntries,
		}

		err = helper.ExportFile(protonOutput, "proton.json")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(toProtonCmd)
}
