package config

import (
	"fmt"
	"log"
)

const helpText = "For more info visit: https://github.com/RA341/ctrl-srv"

func validateQbitSection() {
	if config.Qbit.Enable {
		message := "Qbit is enabled, the following fields must be filled"

		if config.Qbit.Host == "" {
			message += "\nHost is empty\n"
		}
		if config.Qbit.Pass == "" {
			message += "\nPassword is empty\n"
		}
		if config.Qbit.User == "" {
			message += "\nUsername is empty\n"
		}

		if message != "Qbit is enabled, the following fields must be filled" {
			// default message has been changed therefore some error has occurred
			errorMessage(message)
		}
	}
}

func validateDiscordSection() {
	if config.DiscordNotif.Enable {
		message := "Discord notifications are enabled, the following fields must be filled"

		if config.DiscordNotif.WebhookURL == "" {
			message += "\nWebhookURL is empty\n"
		}
		if config.DiscordNotif.Username == "" {
			message += "\nUsername is empty\n"
		}

		if message != "Discord notifications are enabled, the following fields must be filled" {
			// default message has been changed therefore some error has occurred
			errorMessage(message)
		}
	}
}

func validateNetworkSection() {
	message := "Error in Network section, the following fields must be filled"

	if config.Network.Host == "" {
		message += "\nHost is empty\n"
	}

	if message != "Error in Network section, the following fields must be filled" {
		// default message has been changed therefore some error has occurred
		errorMessage(message)
	}
}

func errorMessage(message string) {
	log.Fatal(fmt.Errorf(message+"\n"), helpText)
}
