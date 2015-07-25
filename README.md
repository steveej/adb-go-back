# adb go back

`adb` based backup/restore utility for Android apps written in Go. 

**Be warned: this is a toy**

## Features
* Retrieve and save a list of installed user apps
* Backup apps according to saved list
* Restore apps from backup file


## Dependencies
* go
* `adb` binary must be in *$PATH*
* TODO: list go dependencies

## Build
```
$ ./build
```

## Usage
Please take a look at the help message for details
```
$ ./bin/adbgb --help
```

### Backup
1. Connect your device and make sure adb is working
1. `$ ./bin/adbdb dump -o /tmp/apps.yaml`
1. Edit /tmp/apps.yaml. keep only the entries you want to backup
1. `$ ./bin/adbdb backup -i /tmp/apps.yaml -o apps.bkp`
1. At this point your device should display a message and ask for backup
   confirmation

### Restore
1. Connect your device and make sure adb is working
1. `$ ./bin/adbdb restore -i apps.bkp`
1. At this point your device should display a message and ask for restore
   confirmation

## Roadmap
Please consult and/or use the project's issue list.
