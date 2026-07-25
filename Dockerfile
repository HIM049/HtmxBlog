FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY release/ .

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x ./HtmxBlog /entrypoint.sh

EXPOSE 9590

ENTRYPOINT ["/entrypoint.sh"]
CMD ["./HtmxBlog"]
