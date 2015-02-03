# quickstart-go

A simple Golang web app (using Martini) which can easily be deployed to Tutum.

This application support the [Getting Started with Golang on Tutum]() article - check it out

#Running locally

	$ git clone https://github.com/tutumcloud/quickstart-go
	$ cd quickstart-go
	$ docker build -t quickstart-go .
	$ docker run -d -p 80 quickstart-go -n goweb

Alternatively, you can run the dockerized version:

	$ docker run -d -p 80 tutum/quickstart-go

Your app should now be running

#Deploying to Tutum

[Install the Tutum CLI.](https://support.tutum.co/support/solutions/articles/5000049209-installing-the-command-line-interface-tool)

	$ tutum login
	$ tutum service run tutum/quickstart-go
	
Continue with this tutorial [here]().
	