# Getting Started with Mobile Apps in Go

>March 25, 2021 - Curitiba, BR.

Over a year ago I started learning Go and had an amazing opportunity to use this language in a real production application at my current company. Now I want to keep studying and using Go but at the same time I wanted to build an open-source application that could reach more users.

That's when I decided to build a mobile app using Go! In this document I will share some of my experiences with it.

## Getting the Tools

It may sound obvious, but to get started with some real-world Go app you need to install the SDK in your machine. To do that, you will find all the information you need in the official download page:

https://golang.org/dl/

The other tools you will need to use is `gomobile`. Once you have installed the Go SDK, you can install this second tool by running the following commands on your terminal:

```bash
$ go get golang.org/x/mobile/cmd/gobind
$ go get golang.org/x/mobile/cmd/gomobile
```

>NOTE: In order to get `gomobile` to work on my machine I had to turn Go modules off by setting this environment variable: GO111MODULE=off

## Installing the Android NDK



## Getting and Running a Sample app