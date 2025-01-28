package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dustinestes/MECM_ClientCleaner/pkg/config"
)

var LogFile *os.File = CreateFile(config.LogFilePath)

func CreateFile(file string) *os.File {
	LogFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return LogFile
}

func Write(msg string, file *os.File, indent int, console bool) {
	file.WriteString(fmt.Sprint(strings.Repeat(" ", indent), msg, "\n"))

	if console {
		fmt.Println(fmt.Sprint(strings.Repeat(" ", indent), msg))
	}
}

func WriteHeader() {
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write(config.AppName, LogFile, 2, config.AppOutConsole)
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write(fmt.Sprint("Version: ", config.AppVersion), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("URL: ", config.AppURL), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("Description: ", config.AppDescription), LogFile, 2, config.AppOutConsole)
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write(fmt.Sprint("Executed By: ", config.UserCurrent.Username), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("Start Time: ", config.StartTime), LogFile, 2, config.AppOutConsole)
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write(" ", LogFile, 0, config.AppOutConsole)
}

func WriteInitial() {
	Write("Initialization", LogFile, 2, config.AppOutConsole)
	Write("- Create Program Directories", LogFile, 4, config.AppOutConsole)
	Write("C:\\ProgramData\\MECM Client Cleaner", LogFile, 8, config.AppOutConsole)
	Write("Status: Success", LogFile, 10, config.AppOutConsole)
	Write("C:\\ProgramData\\MECM Client Cleaner\\Logs", LogFile, 8, config.AppOutConsole)
	Write("Status: Success", LogFile, 10, config.AppOutConsole)
	Write("- Create Log File", LogFile, 4, config.AppOutConsole)
	Write(config.LogFilePath, LogFile, 8, config.AppOutConsole)
	Write("Status: Success", LogFile, 10, config.AppOutConsole)
}

func WriteFinal() {
	Write("- Finalize Log File", LogFile, 4, config.AppOutConsole)
	Write("", LogFile, 0, config.AppOutConsole)
}

func WriteFooter() {
	end := time.Now().UTC()
	totalTime := end.Sub(config.StartTime)

	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write(fmt.Sprint("Result: ", "Add Result Logic"), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("Start Time: ", config.StartTime.Format("2006-01-02 15:04:05")), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("End Time: ", end.Format("2006-01-02 15:04:05")), LogFile, 2, config.AppOutConsole)
	Write(fmt.Sprint("Total Time: ", totalTime), LogFile, 2, config.AppOutConsole)
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)
	Write("End of Log", LogFile, 2, config.AppOutConsole)
	Write(strings.Repeat("-", 75), LogFile, 0, config.AppOutConsole)

	LogFile.Close()
}
