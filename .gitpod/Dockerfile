FROM gitpod/workspace-full
USER root

RUN mkdir $HOME/bin
RUN npm install -g @bazel/bazelisk
RUN npm install -g @bazel/ibazel
RUN npm install --global smee-client

ARG TERRAFORM_URL="https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_amd64.zip"
RUN curl -sSL "${TERRAFORM_URL}" -o terraform.zip \
    && unzip terraform.zip \
    && rm terraform.zip \
    && chmod +x terraform \
    && mv terraform /usr/local/bin/terraform

ARG BAZEL_BUILDTOOLS_URL="https://github.com/bazelbuild/buildtools/releases/download"
RUN curl -sSL "${BAZEL_BUILDTOOLS_URL}/5.1.0/buildozer-linux-amd64" > /usr/local/bin/buildozer && \
    curl -sSL "${BAZEL_BUILDTOOLS_URL}/5.1.0/buildifier-linux-amd64" >  /usr/local/bin/buildifier && \
    chmod +x /usr/local/bin/buildozer && \
    chmod +x /usr/local/bin/buildifier 

USER gitpod
