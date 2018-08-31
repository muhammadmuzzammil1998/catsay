
# catsay [![CircleCI](https://circleci.com/gh/muhammadmuzzammil1998/catsay.svg?style=svg)](https://circleci.com/gh/muhammadmuzzammil1998/catsay) [![CodeFactor](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/catsay/badge)](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/catsay)
**catsay** is a program that generates pictures of a cat holding a sign with a message. *LIEK COWSAY BUT WIF KATS*.
```
$ fortune | catsay
 _______________________________________
/ Q: Why don't lawyers go to the beach? \
\ A: The cats keep trying to bury them. /
 ---------------------------------------
  (\__/)||
  (•ㅅ•)||
  /    \っ
```

## Build
Requires [git](https://git-scm.com/download/win) to clone and [Go](https://golang.org/dl/) to build.
```
$ git clone https://github.com/muhammadmuzzammil1998/catsay.git
$ cd catsay
$ go build
```

## Installation
Move it to your `bin`
```
$ mv ./catsay /usr/bin
```

## Usage
Pipe the data you want your cat to say.
### Examples
```
$ fortune | catsay
 _____________________________________________________________________________
/ You will lose your present job and have to become a door to door mayonnaise \
\ salesman.                                                                   /
 -----------------------------------------------------------------------------
  (\__/)||
  (•ㅅ•)||
  /    \っ

$ echo "Hi there\!" | catsay
 ___________
< Hi there! >
 -----------
  (\__/)||
  (•ㅅ•)||
  /    \っ
```

## Seeing blocks instead of eyes, nose, or hand?
This is probably because your terminal doesn't support those characters. lulz.

Try `catsay -ascii` to use ASCII-only characters. 
```
$ echo "Hi there\!" | catsay -ascii
 ___________
< Hi there! >
 -----------
  (\__/) ||
  (*/\*) ||
  /    \ >
```

## Help
```
$ cowsay -help
Usage of catsay:
  -ascii
        Use this if you see blocks here "•ㅅ•っ"
  -version
        Check version of catsay
```
