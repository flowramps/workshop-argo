package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
				<head>
					<style>
						body {
							display: flex;
							justify-content: center;
							align-items: center;
							height: 100vh;
							background-color: #00FF00; /* Cor de fundo verde claro */
							flex-direction: column;
						}
						h1, p {
							text-align: center;
							color: #333; /* Defina a cor do texto que desejar */
						}
						.image-container {
							display: block;
							-webkit-user-select: none;
							margin: auto;
							cursor: zoom-in;
							background-color: hsl(0, 0%, 90%);
							transition: background-color 300ms;
						}
						.image-container img {
							display: block;
							margin: auto;
							max-width: 654px; /* Largura desejada */
							max-height: 654px; /* Altura desejada */
							width: auto;
							height: auto;
							border-radius: 15px; /* Bordas arredondadas */
						}
					</style>
				</head>
				<body>
					<h1>Workshop DevOps !!!!</h1>
					<div class="image-container">
						<img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/flowramps.jpg" alt="Imagem">
					</div>
					<div>
						<p>Nome do Pode: ` + hostname + `</p>
					</div>
				</body>
			</html>
		`))
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

