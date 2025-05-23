![Build](https://github.com/creativeprojects/resticprofile/workflows/Build/badge.svg)
[![codecov](https://codecov.io/gh/creativeprojects/resticprofile/branch/master/graph/badge.svg?token=cUozgF9j4I)](https://codecov.io/gh/creativeprojects/resticprofile)
[![Go Report Card](https://goreportcard.com/badge/github.com/creativeprojects/resticprofile)](https://goreportcard.com/report/github.com/creativeprojects/resticprofile)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=ncloc)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=bugs)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=creativeprojects_resticprofile&metric=security_rating)](https://sonarcloud.io/dashboard?id=creativeprojects_resticprofile)

# resticprofile
Configuration profiles manager for [restic backup](https://restic.net/)

**resticprofile** is the missing link between a configuration file and restic backup. Creating a configuration file for restic has been [discussed before](https://github.com/restic/restic/issues/16), but seems to be a very low priority right now.

With resticprofile:

* You no longer need to remember command parameters and environment variables
* You can create multiple profiles inside one configuration file
* A profile can inherit all the options from another profile
* You can run the forget command before or after a backup (in a section called *retention*)
* You can check a repository before or after a backup
* You can create groups of profiles that will run sequentially
* You can run shell commands before or after running a profile: useful if you need to mount and unmount your backup disk for example
* You can run a shell command if an error occurred (at any time)
* You can send a backup stream via _stdin_
* You can start restic at a lower or higher priority (Priority Class in Windows, *nice* in all unixes) and/or _ionice_ (only available on Linux)
* It can check that you have enough memory before starting a backup. (I've had some backups that literally killed a server with swap disabled)
* You can generate cryptographically secure random keys to use as a restic key file
* You can easily schedule backups, retentions and checks (works for *systemd*, *crond*, *launchd* and *windows task scheduler*)
* You can generate a simple status file to send to some monitoring software and make sure your backups are running fine 
* You can use a template syntax in your configuration file
* You can generate scheduled tasks using *crond*
* Get backup statistics in your status file
* Automatically clear up stale locks
* Export a **prometheus** file after a backup, or send the report to a push gateway automatically
* **[new for v0.17.0]** Run shell commands in the background when non fatal errors are detected from restic
* **[new for v0.18.0]** Send messages to HTTP hooks before, after a successful or failed job (backup, forget, check, prune, copy)
* **[new for v0.18.0]** Automatically initialize the secondary repository using `copy-chunker-params` flag
* **[new for v0.18.0]** Send resticprofile logs to a syslog server
* **[new for v0.19.0]** Preventing your system from idle sleeping
* **[new for v0.21.0]** See the help from both restic and resticprofile via the `help` command or `-h` flag
* **[new for v0.24.0]** Don't schedule a job when the system is running on battery
* **[new for v0.29.0]** Scheduling a group of profiles is finally available (configuration `v2` only)

The configuration file accepts various formats:
* [TOML](https://github.com/toml-lang/toml) : configuration file with extension _.toml_ and _.conf_ to keep compatibility with versions before 0.6.0
* [JSON](https://en.wikipedia.org/wiki/JSON) : configuration file with extension _.json_
* [YAML](https://en.wikipedia.org/wiki/YAML) : configuration file with extension _.yaml_
* [HCL](https://github.com/hashicorp/hcl): configuration file with extension _.hcl_



<!--ts-->

<!--te-->

We recommend you start by reading the [getting started](https://creativeprojects.github.io/resticprofile/configuration/getting_started/index.html) section

# Using resticprofile

The full documentation has been moved to [creativeprojects.github.io](https://creativeprojects.github.io/resticprofile/)

# Survey

What are your most important features?
Please fill in the [survey](https://github.com/creativeprojects/resticprofile/issues/415)

