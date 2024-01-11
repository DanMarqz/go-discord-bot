FROM golang:1.22rc1-alpine3.19
WORKDIR /app

ARG DISCORD_TOKEN
ENV DISCORD_TOKEN=${DISCORD_TOKEN}

RUN touch .env
RUN echo ${DISCORD_TOKEN} >> .env

COPY . .

CMD ["go", "run", "./main.go"]