# wavvy
Cal Poly KCPR tracks to your terminal

## Installation

Download the source code into your `$GOPATH/src` directory
```
$ go get github.com/corykitchens/wavvy
```

Change into the source code directory
```
$ cd $GOPATH/src/github.com/corykitchens/wavvy
```

Run `go install`
```
$ go install -v ./...
```

## Usage
Run the binary with no flags to retrieve the current track as a string
```
$ wavvy
Artist: <artist_name> Title: <title> Album: <album>
```
Use the --json=true flag to return the output as json

```
$ wavvy --json=true
{ artist: <artist_name>, title: <title>, album: <album>}
```