package crawl

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/types"
	"github.com/lestrrat-go/libxml2/xpath"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CrawlBook struct {
}

func NewCrawlBook() *CrawlBook {
	return &CrawlBook{}
}
func (ch *CrawlBook) fetchBiQu(url string) (types.Document, error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "fcip=111; ASP.NET_SessionId=3yo4zvar3cmtwnrghdatj1ia; __gads=ID=211622b936e75ba1-221485f031c50057:T=1607951574:RT=1607951574:S=ALNI_MZJTPFVongaYLiaLTuSKwfHBr1Q2Q; __atuvc=1%7C51; __atuvs=5fd764e2947d7b29000; _ga=GA1.2.556581904.1607951575; _gid=GA1.2.830984973.1607951587; _gat=1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		logrus.Warn("Http status code:", resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := libxml2.ParseHTMLReader(resp.Body)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	return doc, nil

}
func (ch *CrawlBook) fetch(url string) (types.Document, error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "fcip=111; ASP.NET_SessionId=3yo4zvar3cmtwnrghdatj1ia; __gads=ID=211622b936e75ba1-221485f031c50057:T=1607951574:RT=1607951574:S=ALNI_MZJTPFVongaYLiaLTuSKwfHBr1Q2Q; __atuvc=1%7C51; __atuvs=5fd764e2947d7b29000; _ga=GA1.2.556581904.1607951575; _gid=GA1.2.830984973.1607951587; _gat=1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		logrus.Warn("Http status code:", resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := libxml2.ParseHTMLReader(resp.Body)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	return doc, nil

}
func (ch *CrawlBook) BiQuParseUrls(url string) (string, string, error) {
	doc, err := ch.fetch(url)
	if err != nil {
		return "", "", err
	}
	defer doc.Free()
	content := ""
	name := ""
	nodes := xpath.NodeList(doc.Find(`//*[@id="content"]`))
	for _, node := range nodes {
		mcontent, err := node.Find("./text()")
		if err != nil {
			return "", "", err
		}
		for _, m := range mcontent.NodeList() {
			content += m.TextContent()
		}
	}
	nodesName := xpath.NodeList(doc.Find(`//*[@id="wrapper"]/div/div/div/h1`))
	for _, node := range nodesName {
		mname, err := node.Find("./text()")
		if err != nil {
			return "", "", err
		}
		name = mname.NodeList()[0].TextContent()
	}

	return content, name, nil
}

func (ch *CrawlBook) UUParseUrls(url string,title string) (string, string, error) {
	doc, err := ch.fetch(url)
	if err != nil {
		return "", "", err
	}
	defer doc.Free()
	content := ""
	name := ""
	nodes := xpath.NodeList(doc.Find(`//div[@class="postTitle"]/a`))
	for _, node := range nodes {
		mcontent, err := node.Find("./text()")
		if err != nil {
			return "", "", err
		}
		for i, m := range mcontent.NodeList() {
			if m.TextContent()==title{
				nodesName := xpath.NodeList(doc.Find(`//div[@class="postContent`))
				mcontent,err:=nodesName[i+1].Find("//text()")
				if err != nil {
					return "", "", err
				}
				name = mcontent.NodeList()[0].TextContent()
			}
			content += m.TextContent()
		}
	}
	return content, name, nil
}

