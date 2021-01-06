FROM ubuntu:20.04

# base
RUN apt-get update \
    && apt install wget \
    && apt install unzip
    && mkdir -p /home/work \
    && mkdir -p /home/work/bin
    && cd /home/work

# go
RUN wget https://golang.org/dl/go1.15.3.linux-amd64.tar.gz -O go1.15.tar.gz \
    && tar -C /usr/local -xzf go1.15.tar.gz \
    && echo 'PATH=$PATH:/usr/local/go/bin' >> /etc/profile \
    && source /etc/profile

# mongodb
RUN wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2004-4.4.1.tgz -O mongodb.tgz\
    && tar -C /home/work/bin -xzf mongodb.tgz \
    && mv /home/work/bin/mongodb-linux-x86_64-ubuntu2004-4.4.1  /home/work/bin/mongodb \

# rocketmq
RUN wget https://apache.claz.org/rocketmq/4.8.0/rocketmq-all-4.8.0-bin-release.zip -O rocketmq.zip \
    && unzip rocketmq.zip -d  /home/work/bin \
    && mv /home/work/bin/rocketmq-all-4.8.0-bin-release /home/work/bin/rocketmq \
    && apt-get install openjdk-8-jre \




