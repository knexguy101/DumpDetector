# DumpPlug

Uses Windows pathing to detect if new .DMP files are being created using the Task Manager

## Installation

```bash
go get -u github.com/knexguy101/dumpPlug
```

## Usage

```go

import "github.com/knexguy101/dumpPlug"

//errorThreshold - the amount of errors the detector can take before exiting (0 for unlimited)
//panicOnExit - panics if detected or threshold is met
//onDet - callback before panic (if panicOnExit = true) when new dump is found
//onThresh - callback before panic (if panicOnExit = true) when error threshold is met
dumpDetector.DetectDumps(5, true, func(){}, func(){})

```

## Contributing
Made for fun, fork and make better

## To Do
want to find registry address for .DMP location instead of trusting no one will change it

## License
[MIT](https://choosealicense.com/licenses/mit/)
