package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly"
)

func main() {
    // Cria um coletor
    c := colly.NewCollector(
        // Define o User-Agent para evitar bloqueios
        colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
    )

    // Callback para quando um link <a> é encontrado
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        fmt.Println("Link encontrado:", link)
    })

    // Callback para quando uma requisição é feita com sucesso
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visitando:", r.URL.String())
    })

    // Callback para lidar com erros de requisição
    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Erro:", err)
    })

    // Inicia o coletor na URL alvo
    startURL := "https://example.com" // Substitua pela URL que deseja coletar
    err := c.Visit(startURL)
    if err != nil {
        log.Fatal("Erro ao visitar a URL:", err)
    }
}
