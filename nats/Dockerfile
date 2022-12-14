FROM ubuntu:22.10
RUN apt-get update
RUN apt-get --assume-yes install build-essential maven openjdk-17-jdk python3-dev python3-pip
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64
ENV PIP_NO_BINARY=jpy
ENV PATH=$PATH:~/.local/bin
RUN pip3 install -U Jinja2
RUN pip3 install ansible ansible-rulebook ansible-runner nats-py wheel
