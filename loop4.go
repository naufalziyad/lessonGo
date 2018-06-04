package main

// author : Naufal Ziyad Luthfiansyah //

import ("fmt"
"encoding/xml")

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func main() {
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
   for _, Location :=range s.Locations {
      fmt.Println("\n%s",Location)
   }
}