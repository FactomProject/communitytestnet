# Golang ontop of ubuntu
FROM emyrk/ubuntu-golang:1.9.2

# Get git, ssh, and other tools
RUN apt-get update \
    && apt-get -y install openssh-server curl git vim iputils-ping net-tools netcat ufw\
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Setup ssh
RUN mkdir /var/run/sshd
# Password for root if we need debugging without ssh-key access
# RUN echo 'root:screencast' | chpasswd
# RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

# Setup ssh keys
RUN mkdir /root/.ssh
#	Can add from URL?
ADD ./authorized_keys /root/.ssh

# Install Factomd
	# RUN mkdir -p /root/go/bin
	# ENV GOPATH=/root/go
	# # Get glide
	# RUN go get github.com/Masterminds/glide
	# COPY factomd /root/go/src/github.com/FactomProject/factomd


	# # Where factomd sources will live
	# WORKDIR $GOPATH/src/github.com/FactomProject/factomd

	# # Install dependencies
	# RUN /root/go/bin/glide install -v

	# # Build and install factomd
	# ARG GOOS=linux
	# RUN go install -ldflags "-X github.com/FactomProject/factomd/engine.Build=`git rev-parse HEAD` -X github.com/FactomProject/factomd/engine.FactomdVersion=`cat VERSION`"

ADD ./bin/* /root/

# Inserting binaries instead

# Setup the cache directory
RUN mkdir -p /root/.factom/m2
# COPY $GOPATH/src/github.com/FactomProject/factomd/factomd.conf /root/.factom/m2/factomd.conf

## Expose and start
EXPOSE 22 8108 8109 8110 8088 8090 9876
CMD ["/usr/sbin/sshd", "-D"]