FROM golang:1.6-alpine
MAINTAINER Jose L. Vazquez <josvazg@gmail.com>
ADD sparses*.go sparses/
WORKDIR sparses/
