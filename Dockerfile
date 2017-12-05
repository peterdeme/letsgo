FROM scratch
WORKDIR /app
ADD letsgo .
EXPOSE 5000
ENTRYPOINT ["/app/letsgo"]