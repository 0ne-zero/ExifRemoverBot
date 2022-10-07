package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	tg_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	exif_r "github.com/scottleedavis/go-exif-remove"
)

func openLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func configBot() (*tg_api.BotAPI, tg_api.UpdatesChannel, error) {
	bot, err := tg_api.NewBotAPI(API_KEY)
	if err != nil {
		return nil, nil, fmt.Errorf("error occurred during create new bot instance - %w", err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tg_api.NewUpdate(0)
	u.Timeout = 10
	updates := bot.GetUpdatesChan(u)
	return bot, updates, nil
}

func sendError(bot *tg_api.BotAPI, chat_id int64) {
	// U0001F91B = fist emoji
	err_str := "مشکلی پیش اومد, دوباره امتحان کن \U0001F91B \U0001F91B"
	_, err := bot.Send(tg_api.NewMessage(chat_id, err_str))
	if err != nil {
		log.Printf("Error occurred during send error message - %s\n", err.Error())
	}
}

func getFileNameFromURL(url string) string {
	if i := strings.LastIndex(url, "/"); i == -1 {
		return ""
	} else {
		return url[i+1:]
	}
}
func isCommand(text string) bool {
	return strings.HasPrefix(text, "/")
}
func removeExifFromPhoto(bytes []byte) ([]byte, error) {
	removed_exif_bytes, err := exif_r.Remove(bytes)
	if err != nil {
		return nil, fmt.Errorf("error occurred during removing exif from file")
	}
	return removed_exif_bytes, nil
}

// Returns file bytes
func downloadFileFromURL(url string) ([]byte, error) {
	// Download file
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while downloading %s - %w", url, err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	return bytes, err
}

// UNUSED; all files proccessed in memory without even saving on local disk
// func writeBytesToFile(path string, bytes []byte) error {
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	_, err = f.Write(bytes)
// 	return err
// }
// func readFileBytes(path string) ([]byte, error) {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	bytes, err := io.ReadAll(f)
// 	return bytes, err
// }
