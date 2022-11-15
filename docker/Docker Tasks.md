**Task-1:**

Run nginx on the docker and expose it to the port 8080 on localhost. Check in your browser at "http://localhost:8080". 

  * docker run --name nginx-server -p 8080:80 -d nginx
  * Screenshot is attached - https://github.com/priyanka25081999/devops-practice/blob/main/docker/nginx%20homepage.png
  
**Task-2:**

Now run the same nginx server from previous assignment but the index.html should be replaced with the file provided (index.html)
  
  * Screenshot is attached - https://github.com/priyanka25081999/devops-practice/blob/main/docker/nginx-webapp-ass2.png
  
**Task-3:**

Create a simple golang application which takes one input as argument and print that input. Containerize this application in such a way that you can provide argument "Welcome to Velotio" when you run "docker run" command. The container log should print the argument provided by you, if no argument is provided it should print "Hello World!!!"

  * Details are added here - https://github.com/priyanka25081999/simple-go-docker-app
