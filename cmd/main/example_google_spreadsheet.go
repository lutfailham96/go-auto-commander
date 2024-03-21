package main

import (
	"log"

	"github.com/lutfailham96/go-auto-commander/internal/notificator"
	"github.com/lutfailham96/go-auto-commander/internal/spreadsheetmanager"
)

func main() {
	// This file is just an example for the Google Sheets integration.
	// The actual implementation is in the internal/spreadsheetmanager package.
	// The example below is a simplified version of the actual implementation.
	// The actual implementation is more complex and contains error handling.
	// The actual implementation also uses the Google Sheets API to read data from a spreadsheet.

	// Create a new Google Sheets spreadsheet.
	googleSpreadsheet := spreadsheetmanager.NewGoogleSpreadsheet("spreadsheetId", "credentialFile.json")

	// Read data from the Google Sheets spreadsheet.
	readRange := "Sheet1!A1:Z1000"
	googleSpreadsheet.ReadData(readRange)

	// Proceed the sheet data.
	proceedSheetData(googleSpreadsheet)
}

// proceedSheetData proceeds the sheet data.
func proceedSheetData(googleSpreadsheet *spreadsheetmanager.GoogleSpreadsheet) {
	// Proceed the sheet data here.
	for _, row := range googleSpreadsheet.SheetContents.Values {
		// Do something with the row.
		log.Println(row)
	}

	// Send a Discord message.
	tokenBot := "your_discord_bot_token"
	recepientId := "your_discord_user_id"
	message := "Hello, this is a test message."
	err := sendDiscordMessage(tokenBot, recepientId, message)
	if err != nil {
		log.Println(err)
	}
}

// sendDiscordMessage sends a message to a Discord user.
func sendDiscordMessage(tokenBot string, recepientId string, message string) error {
	// Create a new Discord notificator.
	discord, err := notificator.NewDiscord(tokenBot, recepientId)
	if err != nil {
		return err
	}

	return discord.SendMessage(message)
}
