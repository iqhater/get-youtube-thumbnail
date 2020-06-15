#build our docker image with name iqhater/get-youtube-thumbnail
#docker build -t iqhater/get-youtube-thumbnail .

#run our docker container afterwards remove himself
#docker run --rm -it iqhater/get-youtube-thumbnail:latest

#Docker Remove All <none> images (only run in bash terminal)
#docker rmi $(docker images -f "dangling=true" -q)

################################################################



#name of base image
FROM golang:1.14

#need to enable to run tests!
ENV CGO_ENABLED=1

#create a folder where our program will be located
RUN mkdir -p /go/src/github.com/iqhater/get-youtube-thumbnail

#set a working directory with a created folder
WORKDIR /go/src/github.com/iqhater/get-youtube-thumbnail

#Copy all files from source to the Docker's path in the image's filesystem
COPY . /go/src/github.com/iqhater/get-youtube-thumbnail

#run test with coverage and goes to test_data folder. Must be empty 
CMD go test -race -v -cover ./...