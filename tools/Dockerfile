FROM ubuntu:20.04

RUN wget https://golang.org/dl/go1.15.3.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.15.3.linux-amd64.tar.gz \
    && echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc \
    && source ~/.bashrc \
    && wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2004-4.4.1.tgz \
    && mkdir -p /data/mongodb \
    && tar -C /data/mongodb -xzf mongodb-linux-x86_64-ubuntu2004-4.4.1.tgz \
    && cd /data/mongodb \
    && mv mongodb-linux-x86_64-ubuntu2004-4.4.1 mongodb \
    # 这里要放mongodb配置调整代码
    && nohup master/bin/mongod f master/config/master.conf & \


