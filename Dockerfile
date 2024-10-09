FROM golang:1.23.1-alpine

WORKDIR /home/Pallina_Di_Gelato_app


COPY .  /home/Pallina_Di_Gelato_app

ENV PORT=:8080

EXPOSE 8080

RUN go build .

CMD [ "/home/Pallina_Di_Gelato_app/Palline_Di_Gelato" ]

