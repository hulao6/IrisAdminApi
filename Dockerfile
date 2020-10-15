# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.

FROM centos

LABEL maintainer="snowlyg <569616226@qq.com>"

COPY ./cmd/main_lin /go/src/github.com/snowlyg/IrisAdminApi/cmd/main_lin
COPY ./application.example.yml /go/src/github.com/snowlyg/IrisAdminApi/cmd/application.yml
COPY ./rbac_model.conf /go/src/github.com/snowlyg/IrisAdminApi/cmd/rbac_model.conf


# Run the command by default when the container starts.
ENTRYPOINT /go/src/github.com/snowlyg/IrisAdminApi/cmd/main_lin

# Document that the service listens on port 80
EXPOSE 80
