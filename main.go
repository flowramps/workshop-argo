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
						h1 {
							text-align: center;
							color: #333; /* Defina a cor do texto que desejar */
						}
						p {
							text-align: center;
							color: #333; /* Defina a cor do texto que desejar */
						}
					</style>
				</head>
				<body>
					<h1>Workshop DevOps !!!!</h1>
					<div style="display: flex; flex-direction: column; align-items: center;">
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

