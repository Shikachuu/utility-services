FROM golang:1.18 as build
COPY . /app
WORKDIR /app
RUN make && mv /app/bin/csv-service-$(dpkg --print-architecture) /app/bin/csv-service

FROM alpine:3.15
RUN adduser --disabled-password --gecos "" -g app app
COPY --from=build /app/bin/csv-service /home/app/csv-service
RUN chown -R app /home/app && chmod +x /home/app/csv-service
USER app
EXPOSE 8080
CMD ["/home/app/csv-service"]