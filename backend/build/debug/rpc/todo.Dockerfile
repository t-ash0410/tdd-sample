FROM golang:latest

# init
RUN apt-get update && \
  apt-get install -y curl git procps ssh

# git clone setting
RUN mkdir /root/.ssh
ADD .ssh /root/.ssh
RUN chmod 600 /root/.ssh/* && \
  touch /root/.ssh/known_hosts && \
  ssh-keyscan github.com >> /root/.ssh/known_hosts

# git clone 
WORKDIR /source
RUN git clone git@github.com:t-ash0410/tdd-sample.git
WORKDIR /source/tdd-sample
RUN git switch develop && git pull

# build
WORKDIR /source/tdd-sample/backend/cmd
RUN go mod tidy && \
  go build -o /bin/todo rpc/todo.go

CMD ["/bin/todo"]
