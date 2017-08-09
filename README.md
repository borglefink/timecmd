## Description

*timecmd* is a tiny command line utility that takes a command as a parameter and executes it,
displaying start time, end time, and duration after the executed command has finished.

The duration status will look similar to the following.
```
[first the executing command's output (if any), then:]
===================================================
Start time: 2017-08-08 13:12:02.245842894 +0200 CEST
End time  : 2017-08-08 13:12:02.250388331 +0200 CEST
Duration  : 4.545437ms
```

## Usage

```
timecmd [command] [optional parameters to command]
```

## Install

Clone the repository into your GOPATH somewhere and do a **go install**.
