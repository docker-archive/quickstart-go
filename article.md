#1. Introduction

This tutorial will have you deploying a Golang application to Tutum.

To follow this tutorial you must have :

- a free [Tutum account](https://dashboard.tutum.co/accounts/register/)
- one node running on your Tutum account. Read [the documentation](https://support.tutum.co/support/solutions/articles/5000523221-your-first-node)
- Docker installed on your own machine


Note: During this tutorial, *your-username* must be replaced by your own tutum username.

#2. Set up

In this section, you will install the Tutum CLI.

**Linux**

On your Terminal run the following command:

	$ pip install tutum
	
**Mac**

On OSX it is recomended to use Homebrew. If you don't have brew on your computer yet, you can find all the information about it on [http://brew.sh](http://brew.sh).

Installing Homebrew is really simple, you just need to run this command :

	$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

and run 

	$ brew install tutum

**Setting Up Tutum CLI**

Check that is has installed correctly :

	$ tutum -v
	0.11.2
	
Now log in using your Tutum account credentials with the command :

	$ tutum login

You will need your Tutum username all along this tutorial so it will be easier for you if you export it as a environment variable :

	$ export TUTUM_USER=your-username
	
Find the documentation for the Tutum CLI tool and API [here](https://docs.tutum.co/v2/api/?shell)


#3. Prepare the app

In this step, you will prepare a simple Golang application to deploy on Tutum. You can clone the repository to get this sample app :

	$ git clone https://github.com/tutumcloud/quickstart-go.git
	cd quickstart-go
	
This repository contains both Dockerfile, that will allow you to build the container, and the code of the small Golang web application, so you can take a look at what it does.

Next, you have to build this app. Execute the following command to do so and create a Docker image called "quickstart-go":

	$ tutum build --tag quickstart-python .
	
Note: The "." at the end of the previous command tells the Tutum CLI tool to build the image from the Dockerfile in the current directoy.

#4. Push the Docker

In this step you will push the Docker image you previously created. 

	$ tutum image push quickstart-go
	Pushing quickstart-go to Tutum private registry ...
	Tagging quickstart-go as tutum.co/maximeheckel/quickstart-go ...
	Sending image list
	Pushing repository tutum.co/maximeheckel/quickstart-go (1 tags)
	Image 511136ea3c5a already pushed, skipping
	Image 36fd425d7d8a already pushed, skipping
	Image aaabd2b41e22 already pushed, skipping
	Image cd9d7733886c already pushed, skipping
	Image 9c88ec810d47 already pushed, skipping
	Image 9ab0b8f2b291 already pushed, skipping
	Image be9b273aeea2 already pushed, skipping
	Image 3484dc8df496 already pushed, skipping
	Image 289cdbbbe0dd already pushed, skipping
	Image d28ff26cf2f1 already pushed, skipping
	Image 24b3fc7668cb already pushed, skipping
	Image 552e98335c47 already pushed, skipping
	Image 37d37ea476b5 already pushed, skipping
	927df14d7b9e: Image successfully pushed
	a14cdb81e7f5: Image successfully pushed
	b7ca02097fde: Image successfully pushed
	9fd4b48abdf5: Image successfully pushed
	Pushing tag for rev [9fd4b48abdf5] on {https://tutum.co/v1/repositories/your-username/quickstart-go/tags/latest}

Now that the image is pushed on the registry, you can check the image list by using:

	$ tutum image list
	NAME                                 DESCRIPTION
	tutum.co/your-username/quickstart-go
	
#5. Deploy the app as a Tutum service

In this step you will deploy the app as a Tutum Service. In order to do so you just have to run the following command:

	$ tutum service run -p 8080 --name quickstart-go tutum.co/your-username/quickstart-go

The **run** command will create and launch the service unsing the quickstart-go image we preiously built and pushed on the registry.
The **-p 8080** flag is used to publish the port 8080 that is exported in the Dockerfile.

The container will be running after a few minutes. By executing the following command you can check the status of the services you have deployed on Tutum:

	$ tutum service ps
	NAME           UUID      STATUS     IMAGE                             DEPLOYED
	quickstart-go  9d9381b6  ▶ Running  tutum.co/your-username/quickstart-go:latest  13 minutes ago
	
	
In order tog et the url of your Golang application, just run:

	$ tutum container ps
	NAME             UUID      STATUS     IMAGE                                       RUN COMMAND      EXIT CODE  DEPLOYED        PORTS
	quickstart-go-1  accba6a7  ▶ Running  tutum.co/maximeheckel/quickstart-go:latest                              15 minutes ago  quickstart-go-1.your-username.cont.tutum.io:49154->8080/tcp
	
You can now copy and paste the link into your browser or just use curl 
on that URL. In this example the URL is [quickstart-go-1.your-username.cont.tutum.io:49154]().

	$ curl quickstart-go-1.your-username.cont.tutum.io:49154
	<h1>hello, world</h1>
	I'm running on linux with an amd64 CPU
	
You have now deployed your first Golang service on Tutum.

#6. Scale the service

The service we've just deployed is running on a single container. In case this container goes down for some reason, this might be a problem. Tutum gives you the ability to scale your services,i.e. to put your app on more than one container, with the following command:

	$ tutum scale quickstart-go 2
	
Now if you list again your running services you have a second conatiner that has just been launched:

	$ tutum container ps           (master✱)
	NAME             UUID      STATUS     IMAGE                                       RUN COMMAND      EXIT CODE  DEPLOYED       PORTS
	quickstart-go-1  accba6a7  ▶ Running  tutum.co/your-username/quickstart-go:latest                              1 hour ago     quickstart-go-1.your-username.cont.tutum.io:49154->8080/tcp
	quickstart-go-2  7f79cc5c  ▶ Running  tutum.co/your-username/quickstart-go:latest                              2 minutes ago  quickstart-go-2.your-username.cont.tutum.io:49155->8080/tcp
	
#7. View logs

Thanks to the Tutum infrastructure, you can have access to the logs of your services. You can, for instance see the logs for both quickstart-go-1 and quickstart-go-2 (1st example) or just one of those containers (2nd example)

	$ tutum service logs quickstart-go
	
	$ tutum service logs quickstart-go-1
	
You can also have access to the logs through your browser from [tutum.co](tutum.co).

#8. Load balance the service

In this part, you will see how to deploy and use a loadbalancer with Tutum. A loadbalancer will allow you to distribute the incoming request to all the containers inside your service.
The quickstart-go service has 2 containers running, and thanks to [Tutum's HAProxy image](https://github.com/tutumcloud/tutum-docker-clusterproxy), you will be able to do loadbalance on it.

	$ tutum service run -p 80:80/tcp --role global --autorestart ALWAYS --link-service quickstart-go:quickstart-go --name golb tutum/haproxy

