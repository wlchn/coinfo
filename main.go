package main

import (
	"coinfo/source"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func newHTTPClient(rawProxyURL string) *http.Client {
	transport := http.DefaultTransport.(*http.Transport)
	// if rawProxyURL != "" {
	// 	proxyURL, err := url.Parse(rawProxyURL)
	// 	if err != nil {
	// 		logrus.Warnf("Failed to parse proxy URL: %s, error: %v, using system proxy\n", rawProxyURL, err)
	// 	} else {
	// 		transportWithProxy := *transport
	// 		transportWithProxy.Proxy = http.ProxyURL(proxyURL)
	// 		transport = &transportWithProxy
	// 		logrus.Debugf("Using proxy %s", rawProxyURL)
	// 	}
	// }

	logrus.Debugf("HTTP request timeout is set to %d", viper.GetInt("timeout"))
	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(viper.GetInt("timeout")) * time.Second,
	}
}

func main() {
	router := gin.Default()

	router.GET("/", CoinHandler)

	router.Run()
}

// CoinHandler get coin info form cmc or exchanges
func CoinHandler(c *gin.Context) {
	coin, exist := c.GetQuery("coin")
	if !exist {
		coin = "the key is not exist!"
	}

	httpClient := newHTTPClient(viper.GetString("proxy"))
	var client = source.CreateExchangeClient("CoinMarketCap", httpClient)
	if client == nil {
		logrus.Warnf("Unknown exchange %s, skipping", "CoinMarketCap")
	}
	symbolPrice, _ := client.GetSymbolPrice(coin)

	c.JSON(200, symbolPrice)
}
