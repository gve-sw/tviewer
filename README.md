# Topology Viewer with gRPC and XR telemetry

[![Build Status](https://travis-ci.org/sfloresk/tviewer.svg?branch=master)](https://travis-ci.org/sfloresk/tviewer) 
[![codecov](https://codecov.io/gh/sfloresk/tviewer/branch/master/graph/badge.svg)](https://codecov.io/gh/sfloresk/tviewer) 
[![Go Report Card](https://goreportcard.com/badge/github.com/sfloresk/tviewer)](https://goreportcard.com/report/github.com/sfloresk/tviewer) 
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

Example of how to build a real time topology using gRPC, XR Telemetry and a very small Mongo database.

You can add devices directly from the web interface and they will be configured with the correct telemetry subscription. 

![Add Device](https://wwwin-gitlab-sjc.cisco.com/sfloresk/tviewer/blob/master/doc-images/AddDevice.png)

Then, you will see the devices coming up in the graph

![Topology](https://wwwin-gitlab-sjc.cisco.com/sfloresk/tviewer/blob/master/doc-images/Topology.png)

It uses the ISIS adjacency and interface IP information to build the links between devices. In order to get real time information without querying all the time to the server javascript web-sockets are used. 
The rest of the actions (e.g. get devices, add devices) are done with traditional get/post actions using angular JS

The database address needs to be added as an env variable called TELEMETRY_DB

## Usage

From your go path:

Get the code

* go get sfloresk/tviewer

Start the database as a container
* docker run --name tviewerdb -p 27017:27017 -d mongo

Set database env variable
* export TELEMETRY_DB=localhost

Compile project
* go install github.com/sfloresk/tviewer

Run
* ./bin/tviewer

There is a docker file in the repo that you can use to build a container if you like

## Current Limitations

* Only ISIS support
* Only IPv4 topologies 
