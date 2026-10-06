FROM gomicro/goose:3.24.1
WORKDIR /migrations
COPY *.sql entrypoint.sh ./
RUN chmod 0555 /migrations/entrypoint.sh
ENTRYPOINT ["/migrations/entrypoint.sh"]
