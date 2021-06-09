FROM alpine:latest

COPY backend_example backend_example
COPY config.toml config.toml
RUN chmod 777 backend_example
CMD ./backend_example