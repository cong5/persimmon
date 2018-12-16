package service

import (
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
)

type SitemapService struct{}

func (this *SitemapService) BuildSiteMap(domain string, platform string) (string, error) {
	xml := ""
	posts, pErr := postService.GetSlugList(50000, 1)
	if pErr != nil {
		return "", pErr
	}

	if platform == "google" {
		optArr := map[string]string{
			"domain": domain,
		}
		xml = this.makeGoogleSiteMap(optArr, posts)
	} else {
		baiduOptArr := map[string]string{
			"priority":   "0.8",
			"changefreq": "weekly",
			"domain":     domain,
		}
		xml = this.makeBaiduSiteMap(baiduOptArr, posts)
	}

	return xml, nil
}

func (this *SitemapService) makeBaiduSiteMap(optArr map[string]string, posts []models.Posts) string {

	begin := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" 
xmlns:mobile="http://www.baidu.com/schemas/sitemap-mobile/1/">`
	end := `</urlset>`

	items := ""
	lastBuildDate := ""
	for key, val := range posts {
		pubDate := utils.Date("2006-01-02", val.CreatedAt)
		if key == 0 {
			lastBuildDate = pubDate
		}
		items += `<url>
			<loc>http://` + optArr["domain"] + "/post/" + val.Slug + `</loc>
			<mobile:mobile type="pc,mobile"/>
			<lastmod>` + pubDate + `</lastmod>
		  </url>`
	}

	siteInfo := `<url>
    <loc>http://` + optArr["domain"] + `</loc>
    <lastmod>` + lastBuildDate + `</lastmod>
    <changefreq>` + optArr["changefreq"] + `</changefreq>
    <priority>` + optArr["priority"] + `</priority>
  </url>`

	html := begin + siteInfo + items + end
	return html
}

func (this *SitemapService) makeGoogleSiteMap(optArr map[string]string, posts []models.Posts) string {

	begin := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	end := `</urlset>`

	items := ""
	for _, val := range posts {
		items += `<url>
			<loc>http://` + optArr["domain"] + "/post/" + val.Slug + `</loc>
		  </url>`
	}

	html := begin + items + end
	return html
}
