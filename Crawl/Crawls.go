package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"io/ioutil"
	"encoding/xml"
)

var urlChanne = make(chan string, 20000)// url channe distribute task
type Sconfig struct {
	XMLName xml.Name `xml:"note"`   //point to outer label
	SeedUrl string `xml:"url"`   // read url and save url into SeedUrl
}
var Seedurl string
func getSeedUrl(){
	file,err:=os.Open("Default.xml")
	if err!=nil{
			fmt.Printf("error:%v",err)
			return
		}
	defer file.Close()
	data, err :=ioutil.ReadAll(file)
	if err !=nil{
	fmt.Printf("error:%v",err)
		return
	}
	v :=Sconfig{}
	err=xml.Unmarshal(data,&v)
	if err !=nil{
		fmt.Printf("error:%v",err)
		return
	}
	Seedurl=v.SeedUrl
	fmt.Println(Seedurl)



}
func GetJokes(url string){
	doc, err := goquery.NewDocument(url)
	if err != nil{
		fmt.Println(err)
	}
	ul :=doc.Find("ul")
	a:=ul.Eq(6).Find("a")
	a.Each(func(i int,content *goquery.Selection){
		 temp,_:=content.Attr("href")
		 fmt.Println(temp)
	})

}

func main(){
	getSeedUrl()
	getUrl(Seedurl)
	//for i:=0;i< len(urlChanne);i++{
	//	go getUrl()
	//}
	for v:=range urlChanne{
		fmt.Println("sds"+v)
		getUrl(v)
	}
	//for i:=0;i<= len(urlChanne);i++{
	//	<-urlChanne
	//}
	defer   fmt.Println(2343)
	println(urlChanne)
	//	GetJokes(url)             getUrl
	///getUrlFromLink(url)
}

func getUrlFromLink(url string){
	fmt.Println("href")
	 doc1,e :=goquery.NewDocument(url)
	 if e!=nil{
	 	fmt.Println(e)
	 }
	 link :=doc1.Find("link")
	//href :=link.Eq(1).Find("link")
	 link.Each(func(f int, selection *goquery.Selection) {
	 	temp1,_:=selection.Attr("href")
	 	fmt.Println(temp1)

	 })
}
func getUrlFromsrc(url string){
	   fmt.Println("src")
	   doc2, err2:=goquery.NewDocument(url)
	   if err2!=nil{
	   	fmt.Println(err2)
	   }
	   doc2.Find("href")
	   
}
func getUrl(url string){
	fmt.Println(url)
//	fmt.Println(2343)
	doc3,err3 :=goquery.NewDocument(url)
	if err3!=nil{
	fmt.Println(err3)
	}

	doc3.Find("body a").Each(func(index int,item *goquery.Selection){
		LTag :=item
		Link,_:=LTag.Attr("href")
	//	linktext :=LTag.Text()
	//	fmt.Println("dfs"+Link)
		urlChanne <-Link
	})
	doc3.Find("body script").Each(func(index int,item *goquery.Selection){
		LTag :=item
		Link,_:=LTag.Attr("src")
		urlChanne <-Link
//		fmt.Println("shdjsa"+Link)
	})

}