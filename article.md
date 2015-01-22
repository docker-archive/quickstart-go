#1. Introduction

This tutorial will have you deploying a Golang application to Tutum.

To follow this tutorial you must have :

- a free [Tutum account](https://dashboard.tutum.co/accounts/register/)
- one node running on your Tutum account. Read [the documentation](https://support.tutum.co/support/solutions/articles/5000523221-your-first-node)
- Docker installed on your own machine


#2. Set up

In this section, you will install the Tutum CLI.

**Linux**

On your Terminal run the following command:

	pip install tutum
	
**Mac**

On OSX it is recomended to use Homebrew. If you don't have brew on your computer yet, you can find all the information about it on [http://brew.sh](http://brew.sh).

Installing Homebrew is really simple, you just need to run this command :

	ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

and run 

	brew install tutum

**Setting Up Tutum CLI**

Check that is has installed correctly :

	tutum -v
	0.11.2
	
Now log in using your Tutum account credentials with the command :

	tutum login

You will need your Tutum username all along this tutorial so it will be easier for you if you export it as a environment variable :

	export TUTUM_USER=your-username
	
Find the documentation for the Tutum CLI tool and API [here](https://docs.tutum.co/v2/api/?shell)


#3. Prepare the app

In this step, you will prepare a simple Golang application to deploy on Tutum. You can clone the repository to get this sample app :

	git clone https://github.com/tutumcloud/quickstart-go.git
	cd quickstart-go
	
This repository contains both Dockerfile, that will allow you to build the container, and the code of the small Golang application, so you can take a look at what it does.

Next, you have to build this app. Execute the following command to do so and create a Docker image called "quickstart-go":

	tutum build --tag quickstart-python .
	
Note: The "." at the end of the previous command tells the Tutum CLI tool to build the image from the Dockerfile in the current directoy.

#4. Push the Docker 
