FROM ubuntu:20.04

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update -y && apt-get install -y wget gnupg2 docker docker.io

RUN echo 'deb [arch=amd64] https://download.01.org/intel-sgx/sgx_repo/ubuntu focal main' | tee /etc/apt/sources.list.d/intel-sgx.list
RUN wget -qO - https://download.01.org/intel-sgx/sgx_repo/ubuntu/intel-sgx-deb.key | apt-key add -

RUN echo "deb http://apt.llvm.org/focal/ llvm-toolchain-focal-11 main" | tee /etc/apt/sources.list.d/llvm-toolchain-focal-11.list
RUN wget -qO - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -

RUN echo "deb [arch=amd64] https://packages.microsoft.com/ubuntu/20.04/prod focal main" | tee /etc/apt/sources.list.d/msprod.list
RUN wget -qO - https://packages.microsoft.com/keys/microsoft.asc | apt-key add -

RUN wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

RUN apt-get update -y
RUN apt-get install -qq -y snapd build-essential libssl-dev clang-11 libssl-dev gdb libsgx-enclave-common libsgx-quote-ex libprotobuf17 libsgx-dcap-ql libsgx-dcap-ql-dev az-dcap-client open-enclave software-properties-common 

RUN wget https://github.com/edgelesssys/ego/releases/download/v1.5.3/ego_1.5.3_amd64_ubuntu-20.04.deb
RUN apt-get install ./ego_1.5.3_amd64_ubuntu-20.04.deb

RUN apt-get install -y ca-certificates curl


# if you do not use azure
# ego install libsgx-dcap-default-qpl
