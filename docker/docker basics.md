
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
