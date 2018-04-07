FROM golang:latest

# get workdir
RUN mkdir /toDoList

# set workdir
WORKDIR /toDoList

# copy file
ADD  todolist /toDoList
ADD static /toDoList/static

# expose the application to 8081
EXPOSE 8080

# give a permission
RUN chmod +x todolist

# Set the entry point of the container to the application executable
ENTRYPOINT [ "./todolist" ]