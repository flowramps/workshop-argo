FROM golang:1.21 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/server .
COPY flowramps.jpg .  # Copia a imagem para o diretório do container
CMD ["./server"]
