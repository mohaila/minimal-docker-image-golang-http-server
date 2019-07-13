# Minimal Docker image for Golang Hello Http Server
A simple Http server in Golang using http package.<br>
The server use PORT environment variable, otherwise will run on port 3000.
## Build the server
```bash
bash
$ go build
```

## Run the server
```bash
bash
$ ./hellohttp
2019/07/12 23:49:32 Server starting on port 3000
```

## Test the /hello
```bash
bash
$ curl http://localhost:3000/hello
Hello Http!
```

## For Docker build a static binary
You can cross compile for Linux (on macOS (tested) and Windows (untested))
```bash
bash
$ ./build.sh
```

## Verify the binary type
```bash
bash
$ ldd hellohttp
not a dynamic executable
$ file hellohttp
hellohttp: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

## Build Docker image
```bash
bash
$ docker build -t hellohttp:latest .
Sending build context to Docker daemon  5.343MB
Step 1/5 : FROM scratch
 --->
Step 2/5 : ENV PORT 5000
 ---> Running in 7705e6f9a02d
Removing intermediate container 7705e6f9a02d
 ---> 6b848611646f
Step 3/5 : COPY hellohttp /hellohttp
 ---> 2cb0a5a3450c
Step 4/5 : EXPOSE 5000/tcp
 ---> Running in b17706e4d955
Removing intermediate container b17706e4d955
 ---> aca9fe74abba
Step 5/5 : CMD ["/hellohttp"]
 ---> Running in f4d76ac34c4c
Removing intermediate container f4d76ac34c4c
 ---> 5f9f08cef176
Successfully built 5f9f08cef176
Successfully tagged hellohttp:last
```

## Verify the Docker image
```bash
bash
$ docker image ls
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hellohttp           last                5f9f08cef176        6 minutes ago       5.34MB
```

## Use your machine, Kubernetes or Docker Swarm cluster to run the app