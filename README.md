Cine ![Version](http://img.shields.io/badge/status-beta-blue.png)
====

> This application is a work-in-progress.

Cine is an application for combing movie actor data with social network data
so that users can find actors by social influence.


## Getting Started

To run Cine, simply `go get` the application:

```sh
$ go get github.com/benbjohnson/cine/...
```

Then run `cine import` to get the latest data.

```
$ cine import
```

Once the import is complete, run `cine server`:

```
$ cine server
Cine listening on http://localhost:7000
```
