go build -o bin\wio-init.exe ..\wio-init
go build -o bin\wio-generate.exe ..\wio-generate

$old = $env:PATH
$cwd = (Convert-Path .)
$env:PATH = "$cwd\bin;$env:PATH"
wio-generate
$env:PATH = $old

go build
