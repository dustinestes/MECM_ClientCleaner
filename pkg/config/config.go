package config

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

const AppName string = "MECM - Client Cleaner"
const AppVersion string = "0.1.0"
const AppURL string = "https://github.com/dustinestes/MECM_ClientCleaner"
const AppDescription string = "Uninstall the MECM Client and clean remaining elements."
const AppOutConsole bool = true
const appDir string = `C:\ProgramData\MECM Client Cleaner`
const appLogDir string = `C:\ProgramData\MECM Client Cleaner\Logs`

var StartTime time.Time = time.Now().UTC()
var timeStamp string = StartTime.Format("2006-01-02_150405")
var LogFilePath string = fmt.Sprint(appLogDir, "\\", timeStamp, ".log")
var UserCurrent, _ = user.Current()

func InitEnv() {
	err := os.MkdirAll(appDir, 0750)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = os.MkdirAll(appLogDir, 0750)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
