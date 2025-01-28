package mecm

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/dustinestes/MECM_ClientCleaner/pkg/config"
	"github.com/dustinestes/MECM_ClientCleaner/pkg/logging"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/mgr"
)

const ccmDir string = `C:\Windows\ccm`
const cacheDir string = `C:\Windows\ccmcache`
const ccmsetupDir string = `C:\Windows\ccmsetup`

var keysAll []string

func ValidateWMI() {
	logging.Write("- Validate WMI Components", logging.LogFile, 4, config.AppOutConsole)

	// Service State
	logging.Write("Service", logging.LogFile, 8, config.AppOutConsole)
	svc := "winmgmt"

	// Open Connection
	m, err := mgr.Connect()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}
	defer m.Disconnect()

	// Open Service
	s, err := m.OpenService(svc)
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}
	defer s.Close()

	// Get Query
	svcStatus, err := s.Query()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}

	// Output State
	switch svcStatus.State {
	case windows.SERVICE_STOPPED:
		logging.Write("State: Stopped", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_START_PENDING:
		logging.Write("State: Start Pending", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_STOP_PENDING:
		logging.Write("State: Stop Pending", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_RUNNING:
		logging.Write("State: Running", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_CONTINUE_PENDING:
		logging.Write("State: Continue Pending", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_PAUSE_PENDING:
		logging.Write("State: Pause Pending", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_PAUSED:
		logging.Write("State: Paused", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_NO_CHANGE:
		logging.Write("State: No Change", logging.LogFile, 10, config.AppOutConsole)
	default:
		logging.Write("State: Unknown State Value", logging.LogFile, 10, config.AppOutConsole)
	}

	// Get Config
	svcConfig, err := s.Config()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
		return
	}

	// Output StartType
	switch svcConfig.StartType {
	case windows.SERVICE_BOOT_START:
		logging.Write("StartType: Boot Start", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_SYSTEM_START:
		logging.Write("StartType: System Start", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_AUTO_START:
		logging.Write("StartType: Auto Start", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_DEMAND_START:
		logging.Write("StartType: Demand Start", logging.LogFile, 10, config.AppOutConsole)
	case windows.SERVICE_DISABLED:
		logging.Write("StartType: Disabled", logging.LogFile, 10, config.AppOutConsole)
	default:
		logging.Write("StartType: Unknown StartType Value", logging.LogFile, 10, config.AppOutConsole)
	}

	// WMI Repository
	logging.Write("Repository", logging.LogFile, 8, true)
	exe, args := "C:\\Windows\\System32\\wbem\\winmgmt.exe", "/verifyrepository"
	timeout := 10 * time.Second

	// Run Command
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd, err := exec.CommandContext(ctx, exe, args).Output()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
	}

	// Process Output
	if strings.TrimSpace(string(cmd)) != "WMI repository is consistent" {
		logging.Write(fmt.Sprint("Error: ", string(cmd)), logging.LogFile, 10, true)
	} else {
		logging.Write(fmt.Sprint("Status: ", strings.TrimSpace(string(cmd))), logging.LogFile, 10, true)
	}
}

func UninstallClient() {
	logging.Write("- Uninstall MECM Client", logging.LogFile, 4, config.AppOutConsole)

	exe, args := fmt.Sprint(ccmsetupDir, `\ccmsetup.exe`), "/uninstall"
	timeout := 600 * time.Second
	logging.Write(fmt.Sprint("Timeout: ", timeout), logging.LogFile, 8, config.AppOutConsole)

	// Validate Binary & Execute
	_, err := os.Stat(exe)

	if err != nil {
		logging.Write(fmt.Sprint("Error: ", exe, " Not Exist"), logging.LogFile, 10, config.AppOutConsole)
	} else {
		// Run Command
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_, err := exec.CommandContext(ctx, exe, args).Output()
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
		}

		// Process Log
		filePath := "C:\\Windows\\ccmsetup\\Logs\\ccmsetup.log"
		pattern := `<!\[LOG\[(.*?)\]LOG\]!>`

		file, err := os.Open(filePath)
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var lastLine string
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
			return
		}

		// Match RegEx Pattern
		re, err := regexp.Compile(pattern)
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
			return
		}

		match := re.FindStringSubmatch(lastLine)
		logging.Write(fmt.Sprint("Status: ", match[1]), logging.LogFile, 10, config.AppOutConsole)

		// Allow Post Uninstall Changes
		sleep := 60 * time.Second
		logging.Write(fmt.Sprint("Sleep: ", sleep), logging.LogFile, 8, config.AppOutConsole)
		time.Sleep(sleep)
	}
}

func RemoveServices() {
	logging.Write("- Remove Services", logging.LogFile, 4, config.AppOutConsole)
	var svc string = "CcmExec"

	logging.Write(fmt.Sprint("Name: ", svc), logging.LogFile, 8, config.AppOutConsole)

	// Open Connection
	m, err := mgr.Connect()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}
	defer m.Disconnect()

	// Open Service
	s, err := m.OpenService(svc)
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}
	defer s.Close()

	err = s.Delete()
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}
	logging.Write("Status: Success", logging.LogFile, 10, config.AppOutConsole)
}

func RemoveRegistry() {
	logging.Write("- Get Registry Keys", logging.LogFile, 4, config.AppOutConsole)
	var mecmClientReg [4]string
	mecmClientReg[0] = `SOFTWARE\Microsoft\CCM`
	mecmClientReg[1] = `SOFTWARE\Microsoft\CCMSetup`
	mecmClientReg[2] = `SOFTWARE\Microsoft\SMS`
	mecmClientReg[3] = `SOFTWARE\WOW6432Node\Microsoft\CCM`

	for _, reg := range mecmClientReg {
		logging.Write(fmt.Sprint("Root: ", reg), logging.LogFile, 8, config.AppOutConsole)

		// Get Child Items Recursively
		getChildKeys(reg)
	}

	logging.Write("- Remove Registry Keys", logging.LogFile, 4, config.AppOutConsole)

	if len(keysAll) < 1 {
		logging.Write(fmt.Sprint("Skipped: Keys found was ", len(keysAll)), logging.LogFile, 10, config.AppOutConsole)
	} else {
		// Sort Slice
		slices.Sort(keysAll)
		slices.Reverse(keysAll)

		// Remove Keys
		removeKeys(keysAll)
	}
}

func getChildKeys(parent string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, parent, registry.ALL_ACCESS)
	if err != nil {
		if err.Error() == "The system cannot find the file specified." {
			logging.Write(fmt.Sprint("Skipped: ", err), logging.LogFile, 10, config.AppOutConsole)
			return
		} else {
			keysAll = append(keysAll, parent)
		}
	}
	defer k.Close()

	subKeys, err := k.ReadSubKeyNames(0)
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
	}

	for _, sub := range subKeys {
		path := fmt.Sprint(parent, "\\", sub)
		keysAll = append(keysAll, path)

		getChildKeys(path)
	}
}

func removeKeys(k []string) {
	for _, key := range k {
		logging.Write(fmt.Sprint("Key: ", key), logging.LogFile, 8, config.AppOutConsole)

		err := registry.DeleteKey(registry.LOCAL_MACHINE, key)
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		}

		logging.Write("Status: Success", logging.LogFile, 10, config.AppOutConsole)
	}
}

// func RemoveWMI() {
// 	logging.Write("- Remove WMI Namespace", logging.LogFile, 4, config.AppOutConsole)

// 	namespace := "root\\SmsDm"
// 	// wmi := [3]string{"root\\SmsDm", "root\\CCMVDI", "root\\CCM"}
// 	// timeout := 60 * time.Second
// 	// logging.Write(fmt.Sprint("Timeout: ", timeout), logging.LogFile, 8, config.AppOutConsole)

// 	// Initialize COM library
// 	ole.CoInitialize(0)
// 	defer ole.CoUninitialize()

// 	// Create WMI COM Object
// 	wmiSvc, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
// 	if err != nil {
// 		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
// 	}
// 	defer wmiSvc.Release()

// 	//
// 	wmi, err := wmiSvc.QueryInterface(ole.IID_IDispatch)
// 	if err != nil {
// 		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
// 	}
// 	defer wmi.Release()

// 	//
// 	_, err = wmi.CallMethod("DeleteNamespace", 0, []interface{}{namespace})
// 	if err != nil {
// 		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
// 	}

// 	// for _, n := range wmi {
// 	// 	logging.Write(n, logging.LogFile, 8, config.AppOutConsole)

// 	// 	// Run Command

// 	// 	// Process Output

// 	// }
// }

// func RemoveWMI() {
// 	logging.Write("- Remove WMI Namespace", logging.LogFile, 4, config.AppOutConsole)

// 	wmi := [3]string{"SmsDm", "CCMVDI", "CCM"}
// 	timeout := 60 * time.Second
// 	logging.Write(fmt.Sprint("Timeout: ", timeout), logging.LogFile, 8, config.AppOutConsole)

// 	for _, n := range wmi {
// 		logging.Write(fmt.Sprint("root\\", n), logging.LogFile, 8, config.AppOutConsole)

// 		// exe, args := "powershell.exe", fmt.Sprint(" ", "-ExecutionPolicy Bypass", " ", "-Command", " ", fmt.Sprint("Get-WmiObject -query \"Select * From __Namespace Where Name='", n, "'\" -Namespace \"root\" | Remove-WmiObject"))
// 		exe, args := "powershell.exe", fmt.Sprint(" ", "-ExecutionPolicy Bypass", " ", "-Command", " ", fmt.Sprint("Get-WmiObject -query \"Select * From __Namespace Where Name='", n, "'\" -Namespace \"root\""))

// 		logging.Write(fmt.Sprint("Comand: ", exe, args), logging.LogFile, 10, true)
// 		// fmt.Println("EXE: ", exe)
// 		// fmt.Println("ARGS: ", args)

// 		// Run Command
// 		ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 		defer cancel()
// 		cmd, err := exec.CommandContext(ctx, exe, args).CombinedOutput()

// 		if err != nil {
// 			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
// 			continue
// 		}
// 		logging.Write(fmt.Sprint("Status: ", cmd), logging.LogFile, 10, true)

// 		// Process Output
// 		// if string(cmd) != "WMI repository is consistent\n" {
// 		// 	logging.Write(fmt.Sprint("Error: ", string(cmd)), logging.LogFile, 10, true)
// 		// } else {
// 		// 	logging.Write(fmt.Sprint("Status: ", string(cmd)), logging.LogFile, 10, true)
// 		// }
// 	}
// }

func RemoveWMI() {
	logging.Write("- Remove WMI Namespace", logging.LogFile, 4, config.AppOutConsole)

	wmi := [3]string{"SmsDm", "CCMVDI", "CCM"}
	timeout := 60 * time.Second

	for _, w := range wmi {
		logging.Write(fmt.Sprint("root\\", w), logging.LogFile, 8, config.AppOutConsole)

		// Start Process
		cmd := exec.Command("powershell", "-Command", fmt.Sprint("Get-WmiObject -query \"Select * From __Namespace Where Name='", w, "'\" -Namespace \"root\" | Remove-WmiObject"))
		err := cmd.Start()
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, true)
		}
		// logging.Write(fmt.Sprint("Process: ", cmd.Process.Pid), logging.LogFile, 8, config.AppOutConsole)

		// Monitor Process
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
		}()

		// Handle Process Completion
		select {
		case err := <-done:
			if err != nil {
				logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
			} else {
				logging.Write("Status: Success", logging.LogFile, 10, config.AppOutConsole)
			}
		case <-time.After(timeout):
			logging.Write(fmt.Sprint("Status: Process Time Out (", timeout, ")"), logging.LogFile, 10, config.AppOutConsole)
		}
	}
}

func RemoveLogFiles() {
	logging.Write("- Remove CcmSetup Log Files", logging.LogFile, 4, config.AppOutConsole)

	logDir := fmt.Sprint(ccmsetupDir, `\Logs`)

	// Get Directory
	files, err := os.ReadDir(logDir)
	if err != nil {
		logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
		return
	}

	// Remove Contents
	if len(files) == 0 {
		logging.Write("Status: No Files Found", logging.LogFile, 10, config.AppOutConsole)
	} else {
		for _, file := range files {
			logging.Write(file.Name(), logging.LogFile, 8, config.AppOutConsole)

			err = os.Remove(fmt.Sprint(logDir, `\`, file.Name()))
			if err != nil {
				logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
				return
			}
			logging.Write("Status: Success", logging.LogFile, 10, config.AppOutConsole)
		}
	}
}

func RemoveDirectories() {
	logging.Write("- Remove Directories", logging.LogFile, 4, config.AppOutConsole)
	dirs := [3]string{ccmDir, cacheDir, ccmsetupDir}

	for _, dir := range dirs {
		logging.Write(dir, logging.LogFile, 8, config.AppOutConsole)

		// Test Directory
		d, err := os.Open(dir)
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
			continue
		}
		d.Close()

		// Remove Directory
		err = os.RemoveAll(dir)
		if err != nil {
			logging.Write(fmt.Sprint("Error: ", err), logging.LogFile, 10, config.AppOutConsole)
			continue
		}
		logging.Write("Status: Success", logging.LogFile, 10, config.AppOutConsole)
	}
}
