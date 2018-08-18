package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"io/ioutil"
	"encoding/xml"
	"regexp"
	"strings"
)

var urlChanne = make(chan string,200000)// url channe distribute task
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
	getUrl(Seedurl,0)

	     num :=0
	for{
	for v:=range urlChanne{
	    num=num+1
	    fmt.Println("num="+string(num))
		reg :=regexp.MustCompile("javascript")
		if v!="" && reg.MatchString(v)==false{
			getUrl(v,-2)
		}
	}}
	fmt.Println(len(urlChanne))

}

func getUrlFromLink(url string){
	fmt.Println("href")
	 doc1,e :=goquery.NewDocument(url)
	 if e!=nil{
	 	fmt.Println(e)
	 }
	 link :=doc1.Find("link")
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
func getUrl(url string,sum int){

	fmt.Println("url"+url)

 reg :=regexp.MustCompile("https")
	reg2 :=regexp.MustCompile("http")
	if reg.MatchString(url)==false && reg2.MatchString(url)==false{
	url="https:"+url
	}else if reg.MatchString(url)==false && reg2.MatchString(url)==true{
	url=strings.Replace(url,"http","https",1)
	}
	doc3,err3 :=goquery.NewDocument(url)
	if err3!=nil{
	fmt.Println(err3)
	}  else{
	doc3.Find("body a").Each(func(index int,item *goquery.Selection){
		LTag :=item
		Link,_:=LTag.Attr("href")

		if Link!=""{

		if sum!=0{

			urlChanne <-Link
			fmt.Println(Link)

		}else{	urlChanne <-Link
			fmt.Println("dfs"+Link)
		}
		}
	})
	doc3.Find("body script").Each(func(index int,item *goquery.Selection){
		LTag :=item
		Link,_:=LTag.Attr("src")
		if Link!=""{

		if sum!=0{

		urlChanne <-Link
			fmt.Println(Link)

		}else{
			urlChanne <-Link
			fmt.Println("shdjsa"+Link)
		}
		}
	})

	}

}
