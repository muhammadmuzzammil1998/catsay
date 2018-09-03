# catsay [![CircleCI](https://circleci.com/gh/muhammadmuzzammil1998/catsay.svg?style=svg)](https://circleci.com/gh/muhammadmuzzammil1998/catsay) [![CodeFactor](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/catsay/badge)](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/catsay) [![Go Report Card](https://goreportcard.com/badge/github.com/muhammadmuzzammil1998/catsay)](https://goreportcard.com/report/github.com/muhammadmuzzammil1998/catsay) [![GitHub license](https://img.shields.io/github/license/muhammadmuzzammil1998/catsay.svg)](https://github.com/muhammadmuzzammil1998/catsay/blob/master/LICENSE) [![Twitter](https://img.shields.io/twitter/url/https/github.com/muhammadmuzzammil1998/catsay.svg?style=social)](https://twitter.com/intent/tweet?hashtags=catsay&text=Take%20a%20look%20at%20this!%20CatSay%20by%20@mmuzzammil1998&url=https://github.com/muhammadmuzzammil1998/catsay/)

**catsay** is a program that generates pictures of a cat holding a sign with a message. 
![image](https://user-images.githubusercontent.com/12321712/44936312-c7474500-ad91-11e8-87a5-e341f5f55170.png)

## Build
Requires [git](https://git-scm.com/download/win) to clone and [Go](https://golang.org/dl/) to build.
```bash
$ git clone https://github.com/muhammadmuzzammil1998/catsay.git
$ cd catsay
$ go build
```

## Installation
Follow the guide on the [releases](https://github.com/muhammadmuzzammil1998/catsay/releases) page for detailed instructions.
### Linux
Download `.deb` file for catsay from the [releases](https://github.com/muhammadmuzzammil1998/catsay/releases) page.
```bash
$ wget https://github.com/muhammadmuzzammil1998/catsay/releases/download/{version}/catsay-linux-{amd64/386}.deb
$ sudo dpkg -i catsay-linux-{amd64/386}.deb
```
### Windows
Start PowerShell as an admin
```ps
$ Invoke-WebRequest https://github.com/muhammadmuzzammil1998/catsay/releases/download/{version}/catsay-windows-{amd64/386}.exe -OutFile catsay.exe
$ mv .\catsay.exe C:\Windows\catsay.exe
```
### Other OSs
You can download the [binary already built](https://github.com/muhammadmuzzammil1998/catsay/releases) for your system or [build it yourself](https://github.com/muhammadmuzzammil1998/catsay#build).

### NOTE: This should be obvious but still:
 - Adapt {version} number. Check version number from [here](https://github.com/muhammadmuzzammil1998/catsay/releases).
 - Choose your architecture, `amd64` for 64 bit and `386` for 32 bit systems.

## Usage
Pipe the data you want your cat to say.
### Examples
![image](https://user-images.githubusercontent.com/12321712/44936315-c7dfdb80-ad91-11e8-9277-1377706b6da5.png)

## Seeing blocks instead of eyes, nose, or hand?
This is probably because your terminal doesn't support those characters. lulz.

Try `catsay -ascii` to use ASCII-only characters. 
![image](https://user-images.githubusercontent.com/12321712/44936316-c9110880-ad91-11e8-8e0f-05c07666e436.png)

## Help
![image](https://user-images.githubusercontent.com/12321712/44982118-6470cc00-af92-11e8-9df6-4642ff10700e.png)

## Uninstall
### Linux
```bash
$ sudo dpkg --remove catsay
```
### Windows
Start PowerShell as an admin
```ps
$ del C:\Windows\catsay.exe
```
