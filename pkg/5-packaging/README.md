# install fyne cmd

go >= 1.16 
```go
go install fyne.io/fyne/v2/cmd/fyne@latest
```

go < 1.16
```go
go get fyne.io/fyne/v2/cmd/fyne
```

# create a package install

linux
```go
fyne package -os linux -icon myapp.png
``` 

windows
```go
fyne package -os windows -icon myapp.png
```

darwin (mac)
```go
fyne package -os darwin -icon myapp.png
```

# install app

darwin (mac)
```go
fyne install -icon myapp.png
```

linux
```cmd
tar xf <package>.tar.xz
make
sudo make install
```