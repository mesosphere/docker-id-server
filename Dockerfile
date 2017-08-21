FROM busybox:1.27.1-glibc

COPY start /start

ENTRYPOINT []
CMD ["/start"]
