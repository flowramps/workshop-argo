# Metrics Prometheus



Comando para executar carga de stres através do K6
```
k6 run nome-do-script.js
```

você pode usar o seguinte PromQL para consultar o count de cada rede social:
```
social_media_clicks_instagram
social_media_clicks_linkedin
social_media_clicks_github

```

Para calcular a soma total de cliques em todas as redes sociais, você pode usar a função sum:
```
sum(social_media_clicks_instagram) + sum(social_media_clicks_linkedin) + sum(social_media_clicks_github)

```

Você também pode usar a função sum diretamente sobre todas as métricas:
```
sum(social_media_clicks_instagram + social_media_clicks_linkedin + social_media_clicks_github)

```


Para realizar uma consulta no Prometheus Query Language (PromQL) para o contador page_views_total, você pode usar a função sum() para agregar os valores do contador ao longo do tempo. A consulta ficaria assim:
~      
```
sum(page_views_total)

```



Esta consulta retorna a soma total de todas as visualizações de página ao longo do tempo. Se você quiser uma média, pode usar a função rate() para calcular a taxa média de incremento por segundo:
```
rate(page_views_total[5m])

```
Neste exemplo, rate(page_views_total[5m]) calcula a taxa média de incremento nos últimos 5 minutos.

Lembre-se de ajustar o intervalo [5m] conforme necessário com base na granularidade dos seus dados e na resolução desejada para a consulta.










