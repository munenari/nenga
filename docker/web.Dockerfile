FROM golang:1.19-alpine

WORKDIR /code

RUN apk update && \
    apk --no-cache add git
RUN go install github.com/pilu/fresh@latest

ENV PATH /go/bin:$PATH
ENV GOPRIVATE git.arcadia.co.jp

EXPOSE 80

CMD ["fresh"]
