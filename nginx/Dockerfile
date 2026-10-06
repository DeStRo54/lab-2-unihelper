FROM alpine:3.21 AS cert-builder
WORKDIR /certs
# Only the explicitly exported TLS artifacts are kept after the secret mounts disappear.
RUN --mount=type=secret,id=cert_key,required=true,target=/tmp/server.key \
    --mount=type=secret,id=cert_crt,required=true,target=/tmp/server.crt \
    mkdir -p /out/ssl \
    && cp /tmp/server.key /out/ssl/server.key \
    && cp /tmp/server.crt /out/ssl/server.crt \
    && chmod 600 /out/ssl/server.key \
    && chmod 644 /out/ssl/server.crt

FROM nginx:1.28.0-alpine AS runtime
RUN apk add --no-cache curl
COPY default.conf /etc/nginx/conf.d/default.conf
COPY --from=cert-builder --chown=nginx:nginx /out/ssl/ /etc/nginx/ssl/
EXPOSE 8443
