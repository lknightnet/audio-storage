FROM ubuntu

WORKDIR /audio

COPY config/config.yaml config/
COPY audio /audio

EXPOSE 8004

CMD ["./audio"]