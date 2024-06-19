ARG SERVICE="sai-cyclone-interaction"

FROM golang as BUILD

ARG SERVICE

WORKDIR /src/

COPY ./ /src/

RUN go build -o sai-cyclone-interaction -buildvcs=false

FROM ubuntu

ARG SERVICE

WORKDIR /srv

COPY --from=BUILD /src/sai-cyclone-interaction /srv/sai-cyclone-interaction

RUN chmod +x /srv/sai-cyclone-interaction

CMD /srv/sai-cyclone-interaction start
