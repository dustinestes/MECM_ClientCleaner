<a id="readme-top"></a>

<!-- PROJECT HEADER -->
<br />
<div align="center">
  <a href="https://github.com/">
    <img src="github/assets/project_main.jpg" alt="Logo" width="580" height="435">
  </a>
  <br><br><br>

# MECM Client Cleaner

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Unlicense License][license-shield]][license-url]

  <p align="center" style="font-size: 1.1em; font-weight:bold">
    <br />
    <a href="https://github.com/">Documentation</a>
    &nbsp;-&nbsp;
    <a href="https://github.com/">Examples</a>
    &nbsp;-&nbsp;
    <a href="https://github.com/">Bugs</a>
    &nbsp;-&nbsp;
    <a href="https://github.com/">Features</a>
  </p>
</div>

<br>

## Table of Contents

- [MECM Client Cleaner](#mecm-client-cleaner)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
    - [Features](#features)
    - [Process Flow](#process-flow)
    - [Tech Stack](#tech-stack)
      - [Go](#go)
      - [PowerShell](#powershell)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
    - [Standard Execution](#standard-execution)
    - [Command Line Execution](#command-line-execution)
    - [Log Output](#log-output)
  - [Roadmap](#roadmap)
  - [Version History](#version-history)
    - [0.0.0 - Codename: Andesite](#000---codename-andesite)
      - [Changes](#changes)
  - [Contribution](#contribution)
    - [Top contributors:](#top-contributors)
  - [License](#license)
  - [Acknowledgments](#acknowledgments)

<br>

## About

This is a binary tool written in Go that can perform the cleanup and removal of stubborn MECM clients that are not communicating correctly. This tool removes all of the installations, files, services, and WMI namespaces related to the MECM client to provide you with a completely clean system environment. This can resolve issues where clients stop communicating or reinstallations fail to return the endpoint to a healthy state.

__IMPORTANT:__ This tool does not perform a WMI reset or repair. This is because that has never been a supported method for client repair from Microsoft and doing so could have far reaching implications beyond the scope of just the MECM client. I have found that simply removing the namespace from WMI clears out all of the referenced settings and data that the client would consume. There is no need to modify/restore the entire WMI just to have affect on 1 single namespace.

### Features

This tool provides the following features.

- MECM Client Removal
- File Cleanup
- WMI Cleanup
- Log Output
- CLI Parameters
- Silent Execution

### Process Flow

All cleanup steps will attempt to run and if they cannot or are not applicable they will handle that gracefully and provide output to the log.

| Phase           | Steps                               | Targets                                                                                                                                 |
|-----------------|-------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| Initialization  | Create program directories          | C:\ProgramData\MECM Client Cleaner                                                                                                      |
|                 | Initialize logging to file          | C:\ProgramData\MECM Client Cleaner\Logs\[YYYY-MM-DD]_[HHMMSS].log                                                                       |
| Validation      | None                                | N/A                                                                                                                                     |
| Execution       | Remove ccmsetup log files           | C:\Windows\ccmsetup\Logs                                                                                                                |
|                 | Uninstall the MECM Client           | C:\Windows\ccmsetup\ccmsetup.exe /uninstall                                                                                             |
|                 | Remove services                     | SMS Agent Host (CcmExec)                                                                                                                |
|                 | Remove Registry keys                | HKLM:\SOFTWARE\Microsoft\CCM, HKLM:\SOFTWARE\Microsoft\CCMSetup, HKLM:\SOFTWARE\Microsoft\SMS, HKLM:\SOFTWARE\WOW6432Node\Microsoft\CCM |
|                 | Remove WMI namespaces               | root\SmsDm, root\CCMVDI, root\CCM                                                                                                       |
|                 | Remove Directories                  | C:\Windows\CCM, C:\Windows\ccmcache, C:\Windows\ccmsetup                                                                                |
| Completion      | Finalizes log file                  | C:\ProgramData\MECM Client Cleaner\Logs\[YYYY-MM-DD]_[HHMMSS].log                                                                       |

### Tech Stack

This application utilizes the following languages, libraries, packages, or other toolsets:

[![Go][Go]][Go-url] [![PowerShell][PowerShell]][PowerShell-url]

#### Go

- [std](https://pkg.go.dev/std)
- [fmt](https://pkg.go.dev/fmt)
- [os](https://pkg.go.dev/os)
- [time](https://pkg.go.dev/time)
- [x/sys](golang.org/x/sys)

#### PowerShell

- Version: 5.1 or later
- Modules: Microsoft.PowerShell.Management
- Used For: WMI Namespace removal

TODO: Find method to convert this to more native Go code and remove dependency on another language.


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

In order to use this tool, you need to decide how you would like to run it in your environment. Go provides us the ability to run the .go files directly, or they can be compiled into a binary (exe) for distribution.

| Execution Style | Use Case                                                                                                                                                                            |
|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Source Code     | Helpful if you want to evaluate the code and run a test on a single device.                                                                                                         |
| Binary (exe)    | a better option for mass deployment, distribution through a deployment tool such as MECM, or to avoid installing the go toolset on every device you need to run the source code on. |

### Prerequisites

Source Code

- [Go](https://go.dev/)

Binary (exe)

- None

### Installation

There is no installation. The source code can be run from an IDE such as Visual Studio Code and the binary is standalone

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

In order to use the binary (exe), you just need to execute it from a command line or using a deployment tool such as Microsoft Endpoint Configuration Manager (MECM).

### Standard Execution

This is a standard execution using all of the default, built-in configurations and no customization. It will run silently and provide feedback only within the log file.

1. Double click the exe
2. Done

### Command Line Execution

> NOTE: Not executed yet. Here as a placeholder.

This section provides an example of how to run it from the command line along with the supported parameters.

| Parameter | Description                       | Example |
|-----------|-----------------------------------|---------|
| -noclient | Does not perform a client removal |         |
| -nofiles  | Does not perform the file cleanup |         |
| -nowmi    | Does not perform the WMI cleanup  |         |


```cmd
MECMClientCleaner.exe
```

### Log Output

Below is a sample of the output to the log file generated.

Location: C:\ProgramData\MECM Client Cleaner\Logs\YYYY-MM-DD_HHMMSS.log

```
---------------------------------------------------------------------------
  MECM - Client Cleaner
---------------------------------------------------------------------------
  Version: 0.1.0
  URL: https://github.com/dustinestes/MECM_ClientCleaner
  Description:Uninstall the MECM Client and clean remaining elements.
---------------------------------------------------------------------------
  Executed By: DESKTOP-G26FOF3\Test
  Start Time: 2025-01-28 21:49:53.476263 +0000 UTC
---------------------------------------------------------------------------

  Initialization
    - Create Program Directories
        C:\ProgramData\MECM Client Cleaner
          Status: Success
        C:\ProgramData\MECM Client Cleaner\Logs
          Status: Success
    - Create Log File
        C:\ProgramData\MECM Client Cleaner\Logs\2025-01-28_214953.log
          Status: Success
  Validation
    - Validate WMI Components
        Service
          State: Running
          StartType: Auto Start
        Repository
          Status: WMI repository is consistent
  Execution
    - Remove CcmSetup Log Files
        ccmsetup.log
          Status: Success
        client.msi.log
          Status: Success
    - Uninstall MECM Client
        Timeout: 10m0s
          Status: CcmSetup is exiting with return code 0
        Sleep: 1m0s
    - Remove Services
        Name: CcmExec
          Error: The specified service does not exist as an installed service.
    - Get Registry Keys
        Root: SOFTWARE\Microsoft\CCM
          Skipped: The system cannot find the file specified.
        Root: SOFTWARE\Microsoft\CCMSetup
        Root: SOFTWARE\Microsoft\SMS
        Root: SOFTWARE\WOW6432Node\Microsoft\CCM
          Skipped: The system cannot find the file specified.
    - Remove Registry Keys
        Key: SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution\VirtualAppPackages\AppV5XPackages
          Status: Success
        Key: SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution\VirtualAppPackages
          Status: Success
        Key: SOFTWARE\Microsoft\SMS\Mobile Client\Software Distribution
          Status: Success
        Key: SOFTWARE\Microsoft\SMS\Mobile Client\Reboot Management
          Status: Success
        Key: SOFTWARE\Microsoft\SMS\Mobile Client
          Status: Success
        Key: SOFTWARE\Microsoft\SMS\CurrentUser
          Status: Success
        Key: SOFTWARE\Microsoft\SMS
          Status: Success
        Key: SOFTWARE\Microsoft\CCMSetup\Prerequisites
          Status: Success
    - Remove WMI Namespace
        root\SmsDm
          Status: Success
        root\CCMVDI
          Status: Success
        root\CCM
          Status: Success
    - Remove Directories
        C:\Windows\ccm
          Status: Success
        C:\Windows\ccmcache
          Status: Success
        C:\Windows\ccmsetup
          Status: Success
  Completion
    - Finalize Log File

---------------------------------------------------------------------------
  Result: Add Result Logic
  Start Time: 2025-01-28 21:49:53
  End Time: 2025-01-28 21:52:39
  Total Time: 2m46.1366835s
---------------------------------------------------------------------------
  End of Log
---------------------------------------------------------------------------
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Roadmap

See the [open issues](https://github.com/[Author]/[ProjectName]/issues) for a full list of proposed features (and known issues).

- [x] Convert the code from PowerShell to Go
- [ ] Add in parameters for customized execution behavior
- [ ] Add check for admin rights for user who executes file
- [ ] Improve logging output to include error as parameter and process output based on this
- [x] Add Validation section and steps
  - [x] Check WMI integrity
- [ ] Add Result logic for footer


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Version History

This provides a brief review of the last two versions and their respective changes. For a detailed list of all changes, see the versionhistory.md file in the documentation.

### 0.0.0 - Codename: Andesite

This is the initial build of the tool using the original PowerShell code as a starting point to ensure that all functionality is mirrored from one to the other.

#### Changes

- Created GitHub repo
- Added files and directories from the Go Project template repo
- Updated the Readme and clearly defined the features and operation of this application

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contribution

TODO: All of this section

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Top contributors:

<a href="https://github.com/[Author]/[ProjectName]/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=[Author]/[ProjectName]" alt="contrib.rocks image" />
</a>

<br>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

TODO: Determine license before creating Git Repo

Distributed under the [LicenseName]. See `license.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Acknowledgments

TODO: Determine if anything needs to be in this section

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Title](https://google.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/[Author]/[ProjectName].svg?style=for-the-badge
[contributors-url]: https://github.com/[Author]/[ProjectName]/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/[Author]/[ProjectName].svg?style=for-the-badge
[forks-url]: https://github.com/[Author]/[ProjectName]/network/members
[stars-shield]: https://img.shields.io/github/stars/[Author]/[ProjectName].svg?style=for-the-badge
[stars-url]: https://github.com/[Author]/[ProjectName]/stargazers
[issues-shield]: https://img.shields.io/github/issues/[Author]/[ProjectName].svg?style=for-the-badge
[issues-url]: https://github.com/[Author]/[ProjectName]/issues
[license-shield]: https://img.shields.io/github/license/[Author]/[ProjectName].svg?style=for-the-badge
[license-url]: https://github.com/[Author]/[ProjectName]/blob/master/LICENSE.txt
[product-screenshot]: images/screenshot.png
[Go]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[Go-url]: https://go.dev/
[PowerShell]: https://img.shields.io/badge/PowerShell-%235391FE.svg?style=for-the-badge&logo=powershell&logoColor=white
[PowerShell-url]: https://learn.microsoft.com/en-us/powershell/
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com
