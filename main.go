package main

import (
	"net/http"
	"os"
	// add prometheus
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
							display: flex;
							justify-content: center;
							align-items: center;
							border-radius: 15px; /* Adicione bordas arredondadas */
							overflow: hidden; /* Garante que as bordas arredondadas se apliquem corretamente */
							margin: 20px; /* Adicione margem ao redor da imagem */
						}
						.image-container img {
							border-radius: 15px; /* Mantenha as bordas arredondadas na imagem */
							max-width: 100%; /* Garante que a imagem não ultrapasse o contêiner */
							height: auto; /* Mantenha a proporção da imagem */
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
	// add promhttp.Handler()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

