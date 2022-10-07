package main

import (
	"fmt"
	"log"

	tg_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	API_KEY     = "YOUR_BOTFATHER_API_KEY"
	GITHUB_ADDR = "github.com"
)

var (
	START_MSG = fmt.Sprintf("سلام \U0001F44B\n\n- این ربات اطلاعات EXIF عکس ها رو حدف میکنه. روش کارش هم اینطوری هست که عکس رو برای ربات میفرستی و ربات هم عکس رو با حذف کردن اطلاعات EXIF برات میفرسته \U0001F91B\n\n- عکس ها رو بصورت فایل بفرست, در غیر این صورت کیفیت عکس افت میکنه \U0001F91B \U0001F44C\n\n- حالا EXIF چی هست؟\n+ هر تصویری که با دوربین دیجیتال یا دوربین تلفن همراهت می‌گیری، به عنوان یک فایل (که عموما پسوند JPG دارد) در حافظه‌ی دستگاهت ثبت میشه. این تصاویر، در کنار خود تصویر اصلی، اطلاعات زیاد دیگری رو نیز ذخیره می‌کنند. این اطلاعات شامل تاریخ، زمان، تنظیمات دوربین و حتی گاهی اوقات اطلاعات مربوط به محلی که عکس درش گرفته شده است رو هم داره(GPS) \U0001F44E\n\n- امنیت من چی میشه؟\n+ کد این برنامه ازاد هست و هرکس از طریق لینک زیر میتونه کد برنامه رو ببینه یا حتی در بهتر شدنش کمک کنه \U0001F91B \U0001F44C\n %s\n", GITHUB_ADDR)
	HELP_MSG  = ""
)

func init() {
	var err error
	// START_MSG, err = strconv.Unquote(`
	// سلام \U0001F44B
	// این ربات اطلاعات EXIF عکس ها رو حدف میکنه. روش کارش هم اینطوری هست که عکس رو برای ربات میفرستی و ربات هم عکس رو با حذف کردن اطلاعات EXIF برات میفرسته \U0001F44C
	// عکس ها رو بصورت فایل بفرست, در غیر این صورت کیفیت عکس افت میکنه \U0001F91B \U0001F44C

	// - حالا EXIF چی هست؟
	// + هر تصویری که با دوربین دیجیتال یا دوربین تلفن همراهت می‌گیری، به عنوان یک فایل (که عموما پسوند JPG دارد) در حافظه‌ی دستگاهت ثبت میشه. این تصاویر، در کنار خود تصویر اصلی، اطلاعات زیاد دیگری را نیز ذخیره می‌کنند. این اطلاعات شامل تاریخ، زمان، تنظیمات دوربین و حتی گاهی اوقات اطلاعات مربوط به محلی که عکس در اون گرفته شده است رو هم داره(GPS).

	// - امنیت من چی میشه؟
	// + کد این برنامه ازاد هست و هرکس از طریق لینک زیر میتونه کد برنامه رو ببینه یا حتی در بهتر شدنش کمک کنه \U0001F91B \U0001F44C`)
	if err != nil {
		log.Fatalln("Error occurred during strconv.Unquote start message")
	}
	HELP_MSG = START_MSG
}
func main() {
	fmt.Println("If you think program didn't work well, check out log file")
	// Config logger
	f, err := openLogFile()
	if err != nil {
		log.Fatalf("Error occurred during open log file - %s\n", err.Error())
	}
	log.SetOutput(f)
	log.SetFlags(log.Llongfile)

	log.Println("Starting ...")

	// Create bot
	// Config how to update messages
	bot, updates, err := configBot()
	if err != nil {
		log.Fatalf("Error occurred during config bot - %s\n", err.Error())
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	// Runs on every update from bot
	for update := range updates {
		// User sent command
		if isCommand(update.Message.Text) {
			switch update.Message.Text {
			case "/start":
				_, err := bot.Send(tg_api.NewMessage(update.FromChat().ChatConfig().ChatID, START_MSG))
				if err != nil {
					sendError(bot, update.FromChat().ChatConfig().ChatID)
					log.Printf("Error occurred during send start message - %s\n", err.Error())
					continue
				}
			case "/help":
				_, err := bot.Send(tg_api.NewMessage(update.FromChat().ChatConfig().ChatID, HELP_MSG))
				if err != nil {
					sendError(bot, update.FromChat().ChatConfig().ChatID)
					log.Printf("Error occurred during send start message - %s\n", err.Error())
					continue
				}
			}
		}
		// User sent document/file
		if update.Message.Document != nil {
			// Extract document/file download url
			download_url, err := bot.GetFileDirectURL(update.Message.Document.FileID)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during extract document/file download url - %s\n", err.Error())
				continue
			}
			// Download photo
			file_bytes, err := downloadFileFromURL(download_url)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during downloading document/file - %s\n", err.Error())
				continue
			}
			photo_name := getFileNameFromURL(download_url)
			// Get file name
			photo_without_exif_bytes, err := removeExifFromPhoto(file_bytes)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during removing EXIF data - %s\n", err.Error())
				continue
			}

			_, err = bot.Send(tg_api.NewDocument(update.FromChat().ChatConfig().ChatID, tg_api.FileBytes{Name: photo_name, Bytes: photo_without_exif_bytes}))
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during send file - %s\n", err.Error())
				continue
			}
		}
		// User sent photo
		if update.Message.Photo != nil {
			// Get every sent photo and remove exif data form them
			// Each sent photo has four quality, so update.Message.Photo has four item, we need only the last one (original photo)
			// Extract last item (main item)
			main_photo_file_id := update.Message.Photo[len(update.Message.Photo)-1].FileID
			// Extract photo download url
			download_url, err := bot.GetFileDirectURL(main_photo_file_id)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during extract document/file download url - %s\n", err.Error())
				continue
			}
			// Download photo
			photo_bytes, err := downloadFileFromURL(download_url)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during downloading photo - %s\n", err.Error())
				continue
			}
			// Remove EXIF
			photo_without_exif_bytes, err := removeExifFromPhoto(photo_bytes)
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during removing EXIF data - %s\n", err.Error())
				continue
			}
			// Get photo name
			photo_name := getFileNameFromURL(download_url)
			// Send stripped photo (removed exif)
			_, err = bot.Send(tg_api.NewDocument(update.FromChat().ChatConfig().ChatID, tg_api.FileBytes{Name: photo_name, Bytes: photo_without_exif_bytes}))
			if err != nil {
				sendError(bot, update.FromChat().ChatConfig().ChatID)
				log.Printf("Error occurred during send file - %s\n", err.Error())
				continue
			}
		}
	}
}
