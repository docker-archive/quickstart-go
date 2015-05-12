# quickstart-go

[![Deploy to Tutum](https://s.tutum.co/deploy-to-tutum.svg)](https://dashboard.tutum.co/stack/deploy/)

A simple Golang web app (using Martini) which can easily be deployed to Tutum.

This application support the [Getting Started with Golang on Tutum]() article - check it out

## Running locally


Build and run using Docker:

	$ git clone https://github.com/tutumcloud/quickstart-go
	$ cd quickstart-go
	$ docker build -t quickstart-go .
	$ docker run -d -p 80 quickstart-go 

Or using fig:

	$ git clone https://github.com/tutumcloud/quickstart-go
	$ cd quickstart-go
	$ fig up

Or run the pre-built/dockerized version:

	$ docker run -d --env AUTH=no --name mongo tutum/mongodb
	$ docker run -d -p 80 --link mongo:mongo tutum/quickstart-go

## Deploying to Tutum

[Install the Tutum CLI.](https://support.tutum.co/support/solutions/articles/5000049209-installing-the-command-line-interface-tool)

	$ tutum login
	$ tutum service run tutum/quickstart-go
	
Continue with this tutorial [here]().
	