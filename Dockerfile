FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY release/ .

RUN chmod +x ./HtmxBlog

EXPOSE 9590

ENTRYPOINT ["./HtmxBlog"]
