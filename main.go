package main

import (
	"github.com/dustinestes/MECM_ClientCleaner/pkg/config"
	"github.com/dustinestes/MECM_ClientCleaner/pkg/logging"
	"github.com/dustinestes/MECM_ClientCleaner/pkg/mecm"
)

func main() {
	// Environment
	config.InitEnv()

	// Initialization
	logging.WriteHeader()
	defer logging.WriteFooter()
	logging.WriteInitial()

	// // Validation
	logging.Write("Validation", logging.LogFile, 2, config.AppOutConsole)
	mecm.ValidateWMI()

	// // Execution
	logging.Write("Execution", logging.LogFile, 2, config.AppOutConsole)
	mecm.RemoveLogFiles()
	mecm.UninstallClient()
	mecm.RemoveServices()
	mecm.RemoveRegistry()
	mecm.RemoveWMI()
	mecm.RemoveDirectories()

	// Completion
	logging.Write("Completion", logging.LogFile, 2, config.AppOutConsole)
	logging.WriteFinal()
}
