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

                                    h1, p {
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
                                        max-width: 100%;
                                        max-height: 75vh;
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

                                    /* Conteúdo principal */
                                    .content {
                                        background-color: #BE7839; /* #D8923B Cor de fundo do conteúdo principal */
                                        padding: 15px;
                                        margin: 10px;
                                        border-radius: 10px;
                                        box-shadow: 0 4px 4px rgba(0, 0, 0, 0.1);
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
                                    <h1>Workshop DevOps teste !!!!</h1>
                                    <div class="image-container">
                                        <img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/flowramps.jpg" alt="Imagem">
                                    </div>
                                    <div>
                                        <p>Nome do Pod: ` + hostname + `</p>
                                    </div>
                                </div>
                                <div class="footer">
                                    <div class="social-links">
                                        <a href="https://www.instagram.com/flow.ramps/" title="Instagram" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/instagram.png"></a>
                                        <a href="https://www.linkedin.com/in/rafaelrampasso/" title="LinkedIn" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/linkedin.png"></a>
                                    </div>
                                </div>
                            </body>
                        </html>
                    `))
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

