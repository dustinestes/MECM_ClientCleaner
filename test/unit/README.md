# MECM - Client Cleaner - Test - Unit

This content will be used to facilitat the unit testing for this application.

## PowerShell Method

This needs to be converted into a Go method for performing the restoration and cleanup of MECM content.

```powershell
# Folders
    ## Create
        New-Item -Path "C:\Windows\ccm" -ItemType Directory
        New-Item -Path "C:\Windows\ccmsetup" -ItemType Directory
        New-Item -Path "C:\Windows\ccmsetup\Logs" -ItemType Directory
        New-Item -Path "C:\Windows\ccmcache" -ItemType Directory

# Log Files
    ## Create
        for ($i = 0; $i -lt 10; $i++) {
            New-Item -Path "C:\Windows\ccmsetup\Logs" -Name "Logging_$($i).log" -ItemType File
        }

# Client
    ## Install
        Start-Process -FilePath ".\ccmsetup.exe" # -ArgumentList ""
    ## Uninstall
        Start-Process -FilePath ".\ccmsetup.exe" -ArgumentList "/uninstall"

# Services
    ## Create
        New-Service -Name "CcmExec" -BinaryPathName "C:\Windows\ccm\ccmexec.exe" -DisplayName "SMS Agent Host" -StartupType Automatic
    ## Remove
        $service = Get-WmiObject -Class Win32_Service -Filter "Name='CcmExec'"
        $service.delete()

# Registry
    ## Create
        New-Item -Path "HKLM:\SOFTWARE\Microsoft\CCM" | Out-Null
        New-Item -Path "HKLM:\SOFTWARE\Microsoft\CCMSetup" | Out-Null
        New-Item -Path "HKLM:\SOFTWARE\WOW6432Node\Microsoft\CCM" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\CurrentUser" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\Mobile Client" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\Mobile Client\Reboot Management" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution\VirtualAppPackages" | Out-Null
        New-Item -Path  "HKLM:\SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution\VirtualAppPackages\AppV5XPackages" | Out-Null

# Directories
    ## Create
        New-Item -Path "C:\Windows\CCM"
        New-Item -Path "C:\Windows\ccmcache"
        New-Item -Path "C:\Windows\ccmsetup"

# WMI
    ## Create
        $Namespaces = @("SmsDm", "CCMVDI", "CCM")
        foreach ($N in $Namespaces) {
            New-CimInstance -Namespace "root" -ClassName "__Namespace" -Property @{ Name = $N }
        }
```