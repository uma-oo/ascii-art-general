# Ascii-Art-Web-Dockerize:

## Description:
  This project involves containerizing the previous project it using Docker, and ensuring best practices for both the Go code and Docker configurations. The goal is to learn the fundamentals of web development (HTTP, HTML, receiving and sending data) while also gaining hands-on experience with Docker for containerization.

## Authors:
  **ooumayma**\
  **aayoubst**\
  **midbenke**



# Usage: (how to run)

## Getting Started

To build and run the project, follow these steps:
But before make sure that docker is installed if not please execute these commands:

```bash 
curl -fsSL https://get.docker.com/rootless | sh
export PATH=$HOME/bin:$PATH
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock

```

Clone the repository:

 ```bash
git clone https://learn.zone01oujda.ma/git/ooumayma/ascii-art-web-dockerize.git
```

``` bash
cd ascii-art-web-dockerize
```

 Check if the code is working without dockerizing it
```bash
 go run ./app
```
Build the Docker image:

```bash
docker build -t app .
```

Run the Docker container:

```bash
docker run -p 6500:6500 app
```

Access the web server: Open your browser and go to http://localhost:6500 to see the server in action.

## Conclusion
This project is a hands-on opportunity to learn about both web development and Docker, providing you with a solid foundation for containerizing applications and managing dependencies in a modern software development environment.
