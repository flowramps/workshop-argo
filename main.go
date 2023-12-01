package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Definindo contadores para cliques em redes sociais e visualizações de página
var (
	instagramClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_instagram",
		Help: "Number of clicks on Instagram link",
	})

	linkedinClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_linkedin",
		Help: "Number of clicks on LinkedIn link",
	})

	githubClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_github",
		Help: "Number of clicks on GitHub link",
	})

	pageViews = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "page_views_total",
		Help: "Total number of page views",
	})

	metricsMtx sync.Mutex
)

func init() {
	// Registrando contadores no Prometheus
	prometheus.MustRegister(instagramClicks)
	prometheus.MustRegister(linkedinClicks)
	prometheus.MustRegister(githubClicks)
	prometheus.MustRegister(pageViews)
}

func main() {
	// Endpoint principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		// Bloqueando para garantir operações atômicas nos contadores
		metricsMtx.Lock()
		defer metricsMtx.Unlock()

		// Incrementando o contador de visualizações de página
		pageViews.Inc()

		// Escrevendo o HTML da página principal
		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<style>
						/* ... (o resto do seu estilo) ... */
					</style>
				</head>
				<body>
					<div class="navbar">
						<a href="/metrics" style="text-decoration: none; color: #333;">Metrics</a>
					</div>
					<div class="content">
						<h1>Workshop DevOps !!!!</h1>
						<div class="image-container">
							<img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/flowramps.jpg" alt="Imagem" style="width: 60%; height: auto;">
						</div>
						<div class="white-text">
							<p>Nome do Pod: ` + hostname + `</p>
						</div>
					</div>
					<div class="footer">
						<div class="social-links">
							<a href="https://www.instagram.com/flow.ramps/" title="Instagram" target="_blank" rel="noopener" onclick="recordSocialMediaClick('instagram')"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/instagram.png"></a>
							<a href="https://www.linkedin.com/in/rafaelrampasso/" title="LinkedIn" target="_blank" rel="noopener" onclick="recordSocialMediaClick('linkedin')"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/linkedin.png"></a>
							<a href="https://github.com/flowramps" title="Github" target="_blank" rel="noopener" onclick="recordSocialMediaClick('github')"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/github.png"></a>
						</div>
					</div>
					<script>
						function recordSocialMediaClick(socialMedia) {
							fetch('/click?network=' + socialMedia);
						}
					</script>
				</body>
			</html>
		`))
	})

	// Endpoint para registrar cliques em redes sociais
	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		// Obtendo o tipo de rede social a partir dos parâmetros da URL
		network := r.URL.Query().Get("network")

		// Bloqueando para garantir operações atômicas nos contadores
		metricsMtx.Lock()
		defer metricsMtx.Unlock()

		// Incrementando os contadores com base na rede social clicada
		switch network {
		case "instagram":
			instagramClicks.Inc()
		case "linkedin":
			linkedinClicks.Inc()
		case "github":
			githubClicks.Inc()
		}
	})

	// Expondo métricas para Prometheus
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

