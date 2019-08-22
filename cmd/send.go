/*
Copyright © 2019 Jesus Guerrero <jesusant.guerrero@gmail.com>

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
	"flag"
	"fmt"
	"net/smtp"
	"os"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

var email, to, pass, userName, message, order string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send configured email or custom one",
	Long: `Send confugured email to a custom email
		example giveme send --email "jesusant@mctekk.com" --to jesusant.guerrero@gmail.com -o tradicional -s "Pedido JesusMCTekk" -m "Menu tradicional para hoy por favor"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			message = args[0]
		}

		msgs := getMessage(message, userName, order)
		fmt.Println("Your email password:")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		pass = string(bytePassword)
		handleError(err)
		sendEmail(email, pass, []string{to}, msgs)

		fmt.Println("message sent")
	},
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func getMessage(message string, userName string, order string) []byte {
	greet := "Buenos dias"
	bye := "Feliz resto del dia"

	now := time.Now()
	hours := now.Hour()
	customMessage := "Menu " + order + " por favor"

	if hours > 12 && hours != 0 {
		greet = "Buenas tardes"
		bye = "Saludos"
	}

	if len(message) > 0 {
		customMessage = message
	}

	return []byte("To: " + to + "\r\n" +
		"Subject: Pedido " + userName + " (" + order + ")\r\n" +
		"\r\n" +
		greet + "\r\n\n" +
		customMessage + " \r\n" +
		"\r\n" +
		"\r\n gracias, " + bye)
}

func sendEmail(email string, pass string, to []string, msgs []byte) {
	auth := smtp.PlainAuth("", email, pass, os.Getenv("emailhost"))
	err := smtp.SendMail("smtp.gmail.com:587", auth, email, to, msgs)
	handleError(err)
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	sendCmd.Flags().StringVarP(&email, "email", "e", "", "sender email (required)")
	sendCmd.Flags().StringVarP(&to, "to", "t", "", "receiver email (required)")
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "message")
	sendCmd.Flags().StringVarP(&userName, "user", "u", "", "userName (required)")
	sendCmd.Flags().StringVarP(&order, "order", "o", "", "order (required)")

	sendCmd.MarkFlagRequired("email")
	sendCmd.MarkFlagRequired("to")
	sendCmd.MarkFlagRequired("user")
	sendCmd.MarkFlagRequired("order")

	flag.Parse()
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
