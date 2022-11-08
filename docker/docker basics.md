
**Docker**

The main purpose of the docker is to package and containerized applications and to ship them and run them anytime and anywhere as many times you want. 

**Basic commands**

1. **docker run** : It is used to run a container from an image. Running the "docker run nginx" command will run an instance of nginx application on the docker host. If the image is not exists on the host then it will go to the docker hub and pull the image. This is done only for the first time, for the subsequent executions same image will be reused.

2. **docker ps** : It lists all the running containers and some basic info about them such as container id, current status, name etc. Each container get random id/name created by docker

3. **docker ps -a** : Shows all running and previously running containers.

4. **docker stop** : To stop the running container. We should pass either container id/name in the command.

5. **docker rm** : To remove the stopped or exited container permanently.
6. **docker exec** : To execute a command on docker container.
7. **docker images** : To see list of all available images and their size.
8. **docker rmi** : To remove the available image. Always make sure that no container is running of that image. You must stop and delete all dependent containers to be able to delete an image.
9. **docker pull <img_name>** : Just to pull the image, but not run the container.
10. **docker attach <container_id>** : To attach the container running in background to foreground.

**Docker images**

Creating a new image (by own):

1. Identify the base os
2. update the source repository using sudo apt repo command in case of ubuntu OS
3. install dependancies using apt command
4. install python dependancies using pip command
5. copy the source code to /opt directory
6. run the web server

**Dockerfile**

Now, create a dockerfile as "Dockerfile" and then write down the instructions for setting up the application in it such as installing dependancies, where to copy the source code from and to and what is the entry point for an application, etc.
Once done, use the docker build command and specify the dockerfile as input as well as the tag name for an image. This will create an image locally on your system. 
To make it available to public docker hub repository, run the docker push command and specify the name of the image you just created. 


    FROM -> start from a base OS or another image
    RUN -> install all dependancies
    COPY -> copies files from the local system onto docker image
    ENTRYPOINT -> specify entrypoint that allows us to specify a command that will be run when the image is run as a container.

**IMP**: When docker builds the image, it builds these in a layered archiecture. Each line of instruction creates a new layer in the docker image which (just) changes in the previous layer. If any steps fails then it will reuse the previous layers from cache and continue to build the remaining layers.

**Creating a new Docker Image**

1. Demo:

    **vim index.html** 

        <!DOCTYPE html>
            <html>
                    <head>
                            <meta charset="UTF-8" />
                            <meta http-equiv="X-UA-Compatible" content="IE=edge" />
                            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
                            <link rel="icon" href="./logo.jpg" type="image/x-icon" />
                            <title> Typing Animation </title>
                            <link rel="stylesheet" href="./src/index.css" />
                    </head>

                    <body>
                            <div class="container">
                                    <div class="title"> Hey, I am Priyanka Salunke </div>
                                    <div class="subtitle">
                                            <div class="text1"> I am a Developer! </div>
                                    </div>
                            </div>

                            <!-- Javascript link -->
                            <script src="./src/index.js"></script>
                    </body>
            </html>

    You can see the output using -> **lynx index.html**  (Run code on ubuntu terminal-browser called lynx)

**Steps to run this application on localhost:**
1. install apache server -> sudo apt install apache2
2. Start the apache service -> sudo service apache2 start
3. Move the index.html to /var/www/html directory -> sudo mv index.html /var/www/html
4. start the apache service again -> sudo service apache2 start
5. Open the browser and start the app -> http://172.17.0.2/index.html

Now, let's dockerized this applciation
1. Write down all above commands on the notepad
2. Create a Directory and create a Dockerfile

        FROM ubuntu

        RUN apt-get update
        RUN apt install -y apache2
        RUN service apache2 start

        COPY index.html /var/www/html
        ENTRYPOINT service apache2 start
        
3. Build the image using -> docker build . -t my-web-app
-t is used to give the name/tag to an image

4. Now, if you do "docker images", you can see my-web-app appeared there. You can run it using "docker run web-app"

5. To push it to docker repository, we can use the command "docker push web-app".

6. If you are doing first time, then we need to tag the image with our username. So, build the image again -> "docker build . -t priyankass99/my-web-app"

7. Then do "docker push priyankass99/my-web-app"

Now, you can login to docker hub and check your repository with new image!

**Environment variables in Docker:**

1. docker run -e APP_COLOR=blue web-app --> so it will apply blue color to the web application. In the code, we need to add color = os.environ.get("APP_COLOR")
2. You can find list of set environment variables in config section of docker inspect command

**command vs entrypoint:**
1. When you run ubuntu image, it will start and stop immediatly. Why? Bcz unlike VM, docker containers are not meant to host an operating system (OS). Containers are meant to run a specific task or a process such as to host an instance of a web server or application server or a database or simply to carry out some kind of computation or analysis. Once the task is complete the container exits, the container only lives as long as the process inside it is alive. If the web service inside the container is stopped or crashes, then the container exits. 
2. If we see the dockerfile of any application then the "cmd" which stands for command that defines, the program that will be run within the container when it's starts. For the nginx image, it is the nginx command, for the mysql image, it is the mysql command. By default it uses bash. Bash is not a db server or a web server. It is a shell that listens for inputs from a terminal and if it cannot find a terminal, then it exits. If you want your container sleeps for 5sec when it started always, then just add "cmd sleep 5" in the Dockerfile.
3. But what if we want to specify the seconds (above we added 5secs) to the command runtime, like docker run ubuntu sleep 10. There is where entrypoints comes into the picture
4. The entrypoint instructions is like the command instruction and whenever you specify anything from the terminal then it will get appended to the entrypoint like ENTRYPOINT ["sleep"] - add this line in the docketfile and then sleep will be replaced with 10 if we pass 10 (sec) from the terminal.
5. But now, what if the you don't specify the 10 or any seconds from the terminal, then it will throw an error as the program was expecting some values. So, in that case, we always use command and ENTRYPOINT together. You can specify the default value. So, in the case where you do not specify the input from terminal, then the entrypoint will be appended to the command instruction. i.e 
FROM Ubuntu
ENTRYPOINT ["sleep"]
CMD ["5"]
So, this both will be appended as: docker run ubuntu sleep 5

**Docker compose:**
1. With docker compose, we can create configuration file in YAML format (called docker-compose.yaml) and put different services together to running them in this file.
2. Then we can simply run "docker compose up" command to bring up the entire application stack. 
3. This is easier to maintain, implement and run as all changes are always stored in the docker compose configuration file. But, this is all applicable to running containers on a single docker host.
	a. **--link** option -> link is a command line option which can be used to link two containers together
		Example: docker run -d --name=vote -p 5000:80 --link redis:redis voting-app
		Here, redis is the service name followed by colon and name of the host that the voting app is looking for. So, internally it creates an entry in the /etc/hosts file on the voting-app container, it adds an entry with redis with the internal IP of redis container.
	b. We can also do "docker run -d --name=vote -p 5000:80 --link redis:redis --link db:db voting-app", but this will soon to be removed from docker as docker swarm and other networking options provides better way of achieving this

4. Now, let's start creating a docker-compose.yml file:
	a. first we will add container names as key in the yml file and then add name of the image as it's value (key is the image and value is the name of that image to use). If we have used ports to expose for any container then add it under the respective containers. Also, add links. Whichever container requires the link, create a property under it called links and provides an array of links such as redis, db (see above example).
5. Instead of pulling an existing image from the docker registery, we can also build our own image by adding build command in the docker compose.
6. Docker compose versions are important to understand. Currently the latest version is 3.
7. Best example : https://github.com/dockersamples/example-voting-app
8. Note: We need to install docker compose separately
9. Example of docker-compose.yml file (older version-1)
		redis:
		  image: redis
		db:
		  image: postgres:9.4

		vote:
		  image: voting-app
		  ports:
			- 5000:80
		  links:
			- redis
		worker:
		  image: worker-app
		  links:
			- db

		result:
		  image: result-app
		  ports:
			- 5001:80
		  links:
			- db
10. The newer versions of docker-compose establishes the default network, so we do not require to add links option in the services. But we need to add version section in the compose file. 

    Example of docker-compose.yml file (newer version-3)

		version: "3"
		services:
			redis:
			  image: redis

			db:
			  image: postgres:9.4

			vote:
			  image: voting-app
			  ports:
				- 5000:80

			worker:
			  image: worker-app

			result:
			  image: result-app
			  ports:
				- 5001:80


**Docker registry:**
1. It is a central repository of all docker images. The image name can be given as nginx/nginx or kodekloud/click-counter-app. First name is the user account or the orgnization name and the second name is the repository name. If we didn't specify username then docker assumes it as the same as repository/image name. Also, it is by default assumed that all images are stored and pulled from the docker's default register - docker hub. The DNS for this is docker.io, i.e docker.io/kodekloud/click-counter-app or docker.io/nginx/nginx
2. We can also publish an image privately. Many cloud service providers such as GCP, AWS, Azure provide a private registry by default when we open an account with them. 
3. We can also deploy the private registry locally using registry image. And then we can access that image and also push the new image anytime within this network using either localhost if you are on same host or using an IP address if you are on different hosts in that environment. By default docker engine interacts with docker hub. DockerHub is a hosted registry solution by Docker Inc.
4. Example to push on the private local registry:

		$ docker pull nginx:latest
		$ docker image tag nginx:latest localhost:5000/nginx:latest
		$ docker push localhost:5000/nginx:latest

	This will add images to our local private registry. In case, if we want to pull the nginx image on accidental deletion, then we can pull it from our local registry using "docker pull localhost:5000/nginx:latest" command.

**docker engine:**
1. When we install a docker on linux host, it will install 3 different components: a. Docker CLI b. REST API c. Docker Daemon

	a. **Docker Daemon** : It is the background process that manages the docker objects such as the images, containers, volumes and networks.

	b. **Docker REST API server** : It is the API interface that can use to talk to the daemon and provide instructions. 

	c. **Docker CLI** : It is the command line interface used to perform actions such as running containers, stopping containers, destroying images etc. It uses REST API to interact with the docker daemon.

2. Docker uses namespace to isolate the workspace. Process ID, network, mount, interprocess communication, unix timesharing are created in their own namespace thereby providing isolation between containers.

3. **Namespaces - PID:** When we create a container on the linux system then it can be treated as creating a child subsystem which conists of it's bunch of own processes orginating from root process with process id of one (init process). But two processes cannot have same process id. This is where namespaces comes into the picture. So, the container's first process always get an id-1 on that namespace but in the main system it will get a next available process id as the 1st process is always init (which has id-1).

	a. Example: If you run any container and assume that there is only one process is running. Now, if you do "docker exec <container_id> ps -eaf" this will list down all the running processes inside that container. Then you will probably get process-id as 1. 

	b. Now, if you check the same running process outside of the docker container ie. on the host using "ps -aef | grep <process_name>" (you can get the process name from above exec cmd's response), then you can see some different process id assigned to it. 

	c. So, with namespaces we can able to give multiple process id's to the same process, thereby making a container think that it is a root process on that container. Whereas in fact it is just another process running on the underlying docker host. And containerization technology allows this process to be isloated and run inside a container using namespaces.

4. Docker host as well as docker containers share the same system resources such as CPU and memory. By default there is no restriction as to how much of a resource a container can use and hence containers may end up utilizing all of the resources on the underlying host. Docker uses 3 groups to restrict the amount of hardware resources allocated to each container. This can be done using --cpus option of docker run command.

	**Example:**
		
		a. docker run --cpus=.5 ubuntu
		This will ensure that the container does not take up more than 50% of the host CPU at any given time. 
		
		b. docker run --memory=100m ubuntu
		This will ensure that the container does not take up more than 100megabytes of the host memory at any given time.
