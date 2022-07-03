FROM alpine:latest

RUN mkdir /app

COPY stockServiceApp /app

CMD [ "/app/stockServiceApp" ]