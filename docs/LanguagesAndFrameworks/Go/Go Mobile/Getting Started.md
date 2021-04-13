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

To install the Android Native Development Kit (NDK), follow these steps:

1. Install Android Studio
2. Once installed, load any project, or create a new one.
3. With the project opened, click on `Tools` >> `SDK Manager` >> `SDK Tools` >> and check `NDK (Side by side)` and `CMake`.
4. Click `Apply` and `Ok`.

## Create an Environment Variable for NDK

After installing the NDK, create the following environment variable:

```
$ANDROID_NDK_HOME = /c/Users/my-user/AppData/Local/Android/Sdk/ndk/22.1.7171670/
```

## Installing ADB on Windows

1. Download the latest version from here: https://dl.google.com/android/repository/platform-tools-latest-windows.zip
2. Extract the contents to an easily accessible folder
3. In your command prompt, navigate to that folder
4. Enable developer mode on your phone and then connect it to your PC
5. In the command prompt enter the following command: `adb devices`
6. Allow USB debugging access from your phone.
7. Reenter the previous command. You should see a message message like:

```
List of devices attached
9d1ds14156d     device
```

## Install gomobile

```bash
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init
```

## Grab the Demo App

```bash
$ go get -d golang.org/x/mobile/example/basic
```

## Building

```bash
$ gomobile build -target=android golang.org/x/mobile/example/basic
```

## Deploying to Android

The above command is going to generate a file with the extension `.apk`, which can be installed on an Android device.

To install it, with your device connected, run the following command:

```bash
$ adb install <path_to_apk>
```

OR

```bash
$ gomobile install golang.org/x/mobile/example/basic
```

After this, you will see a new icon on your Android device.