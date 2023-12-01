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

	metricsMtx sync.Mutex
)

func init() {
	prometheus.MustRegister(instagramClicks)
	prometheus.MustRegister(linkedinClicks)
	prometheus.MustRegister(githubClicks)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Write([]byte(`
            <!DOCTYPE html>
            <html lang="en">
                <head>
                    <meta charset="UTF-8">
                    <meta name="viewport" content="width=device-width, initial-scale=1.0">
                    <style>
                        body {
                            display: flex;
                            flex-direction: column;
                            align-items: center;
                            min-height: 100vh;
                            margin: 0;
                            font-family: Arial, sans-serif;
                            background-color: #D8923B; /* #BE7839 Cor de fundo da tela */
                        }

                        h1 {
                            text-align: center;
                            color: white; /* Alterado para a cor branca */
                        }

                        p {
                            text-align: center;
                            color: #333;
                        }

                        .image-container {
                            display: block;
                            margin: auto;
                            cursor: zoom-in;
                            transition: background-color 300ms;
                        }

                        .image-container img {
                            display: block;
                            margin: auto;
                            border-radius: 10px;
                        }

                        /* Navbar styles */
                        .navbar {
                            display: flex;
                            justify-content: space-between;
                            align-items: center;
                            padding: 10px;
                            background-color: white; /* Cor de fundo da barra de navegação */
                            width: 100%;
                        }

                        .navbar a {
                            margin-right: 5px;
                        }

                        /* Adicione o seguinte estilo para o botão Metrics */
                        .navbar a[href="/metrics"] {
                            text-decoration: none;
                            color: #333;
                            padding: 8px 16px;
                            border: 2px solid #000; /* Adiciona uma borda preta de 2px */
                            border-radius: 10px; /* Adiciona bordas arredondadas */
                            transition: background-color 300ms;
                        }

                        /* Conteúdo principal */
                        .content {
                            background-color: #BE7839; /* #D8923B Cor de fundo do conteúdo principal */
                            padding: 7px;
                            margin: 7px;
                            border-radius: 15px;
                            box-shadow: 0 0px 0px rgba(0, 0, 0, 0.1);
                        }

                        /* Alteração para a cor branca do texto */
                        .white-text p {
                            color: white;
                        }

                        /* Footer styles */
                        .footer {
                            display: flex;
                            justify-content: space-around;
                            align-items: center;
                            padding: 5px;
                            background-color: white; /* Cor de fundo do rodapé */
                            color: #333;
                            width: 100%;
                        }

                        .social-links img {
                            height: 30px;
                            width: 30px;
                            border-radius: 25%; /* Adiciona borda circular às imagens */
                            margin-right: 5px;
                            cursor: pointer;
                        }

                        /* Media query for smaller screens */
                        @media (max-width: 400px) {
                            .navbar {
                                flex-direction: column;
                                align-items: center;
                            }

                            .navbar a {
                                margin-right: 10px;
                                margin-bottom: 10px;
                            }

                            .social-links {
                                margin-top: 10px;
                            }
                        }
                    </style>
                </head>
                <body>
                    <div class="navbar">
                        <a href="/metrics" style="text-decoration: none; color: #333;">Metrics</a>
                        <!-- Adicione aqui o conteúdo da barra de navegação -->
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

	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		network := r.URL.Query().Get("network")
		metricsMtx.Lock()
		defer metricsMtx.Unlock()

		switch network {
		case "instagram":
			instagramClicks.Inc()
		case "linkedin":
			linkedinClicks.Inc()
		case "github":
			githubClicks.Inc()
		}
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

