package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
	prometheus.MustRegister(instagramClicks)
	prometheus.MustRegister(linkedinClicks)
	prometheus.MustRegister(githubClicks)
	prometheus.MustRegister(pageViews)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		metricsMtx.Lock()
		defer metricsMtx.Unlock()

		pageViews.Inc()

		w.Write([]byte(`
            <!DOCTYPE html>
            <html lang="en">
                <head>
                    <meta charset="UTF-8">
                    <meta name="viewport" content="width=device-width, initial-scale=1.0">
                    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
                    <script>
                    $(document).ready(function() {
                        $(".social-links a").click(function() {
                            var socialMedia = $(this).attr("title").toLowerCase();
                            $.get("/increment-" + socialMedia + "-counter", function(data) {
                                console.log("Contador do " + socialMedia + " incrementado");
                            });
                        });
                    });
                    </script>
                    <style>
                    /* Adicione os estilos CSS conforme necessário */
                    </style>
                </head>
                <body>
                    <div class="navbar">
                        <a href="/metrics" style="text-decoration: none; color: #333;">Metrics</a>
                        <!-- Adicione aqui o conteúdo da barra de navegação -->
                    </div>
                    <div class="content">
                        <h1>Workshop DevOps - Ri Happy !!!!</h1>
                        <div class="image-container">
                            <img class="responsive-image" src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/flowramps.jpg" alt="Imagem">
                        </div>
                        <div class="white-text">
                            <p>Nome do Pod: ` + hostname + `</p>
                        </div>
                    </div>
                    <div class="footer">
                        <div class="social-links">
                            <a href="https://www.instagram.com/flow.ramps/" title="Instagram" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/instagram.png"></a>
                            <a href="https://www.linkedin.com/in/rafaelrampasso/" title="LinkedIn" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/linkedin.png"></a>
                            <a href="https://github.com/flowramps/" title="Github" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/github.png"></a>
                        </div>
                    </div>
                </body>
            </html>
        `))
	})

	http.HandleFunc("/increment-instagram-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		instagramClicks.Inc()
		w.Write([]byte("Contador do Instagram incrementado"))
	})

	http.HandleFunc("/increment-linkedin-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		linkedinClicks.Inc()
		w.Write([]byte("Contador do LinkedIn incrementado"))
	})

	http.HandleFunc("/increment-github-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		githubClicks.Inc()
		w.Write([]byte("Contador do GitHub incrementado"))
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
