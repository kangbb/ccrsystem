FROM golang:latest

# get workdir
RUN mkdir /ccrSystem

# set workdir
WORKDIR /ccrSystem

# copy file
ADD  ccrsystem /ccrSystem
ADD static/ /ccrSystem/static
ADD views/ /ccrSystem/views
ADD data/ /ccrSystem/data

# expose the application to 8080
EXPOSE 8080

# give a permission
RUN chmod +x ccrsystem

# Set the entry point of the container to the application executable
ENTRYPOINT [ "./ccrsystem" ]