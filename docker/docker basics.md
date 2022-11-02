
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
