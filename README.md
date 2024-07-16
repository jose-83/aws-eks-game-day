# LlamaPenPatroller

This service patrolls the llama pens.

## Getting Started

These instructions will cover usage information for the docker container

### Prerequisities

In order to build this container you'll need docker installed.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

### Building
Build the container image using the Dockerfile in this repository.

```
docker build -t penpatroller .
```

```
docker tag penpatroller <ACCOUNT_ID>.dkr.ecr.<REGION>.amazonaws.com/penpatroller
```

### Pushing the image to our repository
```
docker push <ACCOUNT_ID>.dkr.ecr.<REGION>.amazonaws.com/penpatroller
```

### NOTE
We need to add - name: "AWS_ACCESS_KEY_ID" and "AWS_SECRET_ACCESS_KEY" to the llama-shaver.