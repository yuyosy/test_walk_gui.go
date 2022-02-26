# test-walk-gui.go

Test walk GUI App

## Build

### Compile Manifest File
Compile the manifest using the [rsrc](https://github.com/akavel/rsrc)

```
rsrc -manifest main.manifest -o main.syso
```

### Build (with-console)
```
go build
```

### Build (without-console)
Building an application without a Console window
```
go build -ldflags="-H windowsgui"
```

## License
MIT License Copyright (c) 2022 yuyosy

This software includes the work that is distributed in Copyright (c) 2010 The Walk Authors (BSD-style license).