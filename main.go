package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Service struct {
	Image string
	Title string
	Desc  string
	Price string
}

type GalleryItem struct {
	Image string
	Title string
	Link  string
}

type Review struct {
	Name    string
	City    string
	Date    string
	Rating  int
	Text    string
	Service string
	Initial string
}

type FAQ struct {
	Question string
	Answer   string
}

type QReply struct {
	Author string
	Date   string
	Text   string
}

type QandA struct {
	ID       int
	Name     string
	Date     string
	Question string
	Replies  []QReply
}

type SiteData struct {
	Title             string
	Name              string
	Slogan            string
	About             string
	LongAbout         string
	Phone             string
	PhoneRaw          string
	Whatsapp          string
	Address           string
	Instagram         string
	Experience        string
	OrdersCount       string
	Rating            string
	ReviewsNum        string
	HeroImage         string
	AboutImage        string
	Services          []Service
	Gallery           []GalleryItem
	RestorationImages []GalleryItem
	ParcheImages      []GalleryItem
	SandaliImages     []GalleryItem
	Reviews           []Review
	FAQs              []FAQ
	QandAs            []QandA
}

type RestorationData struct {
	Name      string
	Phone     string
	PhoneRaw  string
	Whatsapp  string
	Instagram string
	Images    []GalleryItem
}

func main() {
	if err := os.MkdirAll("docs", 0755); err != nil {
		log.Fatal(err)
	}

	data := SiteData{
		Title:       "تعمیرات مبل شیراز",
		Name:        "مبل شیراز",
		Slogan:      "تعمیر تخصصی انواع مبل، صندلی و پرسی در کارگاه ما",
		About:       "با بیش از ۷ سال تجربه، مبل شما را در کارگاه تخصصی‌مان تعمیر می‌کنیم",
		LongAbout:   "کارگاه تعمیرات مبل قهرمانی با بیش از ۷ سال سابقه در زمینه تعمیر و بازسازی انواع مبل، صندلی و پرسی در شیراز فعالیت می‌کند. ما با بهره‌گیری از ابزارهای حرفه‌ای و مواد اولیه باکیفیت، مبل قدیمی شما را به روزهای اوجش برمی‌گردانیم.",
		Phone:       "09013643428",
		PhoneRaw:    "+989013643428",
		Whatsapp:    "989013643428",
		Address:     "شیراز، سلطان‌آباد، خیابان مسکن مهر، اولین دوربرگردان سمت چپ، کارگاه تعمیرات مبل قهرمانی",
		Instagram:   "moblshiraz.ir",
		Experience:  "۷",
		OrdersCount: "۱۳۷",
		Rating:      "4.9",
		ReviewsNum:  "480",
		HeroImage:   "images/hero.jpg",
		AboutImage:  "images/about.jpg",
		Services: []Service{
			{Image: "images/service-3.jpg", Title: "تعمیر مبل چستر", Desc: "تعمیر و بازسازی کامل مبل‌های استیل در کارگاه", Price: "از ۵۰۰ هزار تومان"},
			{Image: "images/service-2.jpg", Title: "تعمیر صندلی", Desc: "تعمیر انواع صندلی اداری و غذاخوری", Price: "از ۲۰۰ هزار تومان"},
			{Image: "images/service-1.jpg", Title: "تعویض پارچه", Desc: "تعویض پارچه مبل با جدیدترین طرح‌ها", Price: "از ۸۰۰ هزار تومان"},
			{Image: "images/service-4.jpg", Title: "تعمیر مبل راحتی", Desc: "تعمیر اسکلت چوبی و فنرهای مبل", Price: "از ۴۰۰ هزار تومان"},
			{Image: "images/service-5.jpg", Title: "تعمیرات سرویس خواب", Desc: "روکش‌کشی حرفه‌ای انواع مبل", Price: "از ۱ میلیون تومان"},
			{Image: "images/service-6.jpg", Title: "بازسازی کامل", Desc: "بازسازی صفر تا صد مبل قدیمی", Price: "از ۲ میلیون تومان"},
		},
		Gallery: []GalleryItem{
			{Image: "images/gallery-1.jpg", Title: "تعمیر مبل چستر"},
			{Image: "images/gallery-4.jpg", Title: "تعمیر صندلی", Link: "gallery-sandali.html"},
			{Image: "images/gallery-3.jpg", Title: "تعویض پارچه", Link: "gallery-parche.html"},
			{Image: "images/gallery-2.jpg", Title: "تعمیر مبل راحتی"},
                        {Image: "images/gallery-5.jpg", Title: "تعمیرات سرویس خواب"},
			{Image: "images/gallery-6.jpg", Title: "بازسازی کامل", Link: "gallery-restoration.html"},
		},
		RestorationImages: []GalleryItem{
			{Image: "images/before-after-1.jpg", Title: "بازسازی کامل مبل راحتی"},
		},
		ParcheImages: []GalleryItem{
			{Image: "images/before-after-2.jpg", Title: "تعویض پارچه مبل"},
  		        {Image: "images/before-after-5.jpg", Title: "تعویض پارچه مبل ال"},
			{Image: "images/before-after-6.jpg", Title: "تعویض پارچه مبل استیل"},
		},
		SandaliImages: []GalleryItem{
			{Image: "images/before-after-3.jpg", Title: "تعویض پارچه صندلی ناهارخوری"},
			{Image: "images/before-after-4.jpg", Title: "تعمیر و تعویض پارچه صندلی"},
		},
		Reviews: []Review{
			{Name: "مریم رضایی", City: "شیراز", Date: "۲ روز پیش", Rating: 5, Service: "تعمیر مبل استیل", Initial: "م", Text: "مبل استیلمون رو بردیم کارگاهشون و واقعاً حرفه‌ای کار کردن. مبل که فکر می‌کردیم باید عوضش کنیم رو مثل روز اولش کردن. قیمت هم کاملاً منصفانه بود."},
			{Name: "علی محمدی", City: "شیراز", Date: "۱ هفته پیش", Rating: 5, Service: "بازسازی کامل", Initial: "ع", Text: "مبل چستر قدیمی مادرم رو کامل بازسازی کردن. از پارچه‌گیری تا اسکلت و فنر. نتیجه واقعاً عالی شد. حتماً به همه توصیه می‌کنم."},
			{Name: "زهرا کریمی", City: "شیراز", Date: "۲ هفته پیش", Rating: 4, Service: "تعویض پارچه", Initial: "ز", Text: "پارچه مبل رو عوض کردن و خیلی تمیز و حرفه‌ای کار کردن. کارگاهشون هم مرتب و منظمه."},
			{Name: "حسین احمدی", City: "شیراز", Date: "۳ هفته پیش", Rating: 5, Service: "تعمیر صندلی", Initial: "ح", Text: "چند تا صندلی اداری داشتم که فکر نمی‌کردم درست بشن. آقای قهرمانی واقعاً کارشون درجه یکه. صندلی‌ها مثل نو شدن."},
			{Name: "فاطمه نوری", City: "مرودشت", Date: "۱ ماه پیش", Rating: 4, Service: "روکش مبل", Initial: "ف", Text: "کارشون تمیز و باکیفیته. من از مرودشت اومدم شیراز و ارزشش رو داشت. نتیجه کار عالی بود."},
			{Name: "رضا مرادی", City: "شیراز", Date: "۱ ماه پیش", Rating: 5, Service: "تعمیر اسکلت", Initial: "ر", Text: "اسکلت مبل خیلی خراب بود و از چند جا پرسیدم گفتن باید مبل رو عوض کنی. ولی اینجا درستش کردن. عالی بود."},
			{Name: "سارا حسینی", City: "شیراز", Date: "۱ ماه پیش", Rating: 5, Service: "تعمیر مبل استیل", Initial: "س", Text: "خیلی وقت‌شناس و منصف هستن. قبل از شروع کار، قیمت رو دقیق گفتن و هیچ هزینه اضافه‌ای نگرفتن."},
			{Name: "محمد کریمی", City: "کازرون", Date: "۲ ماه پیش", Rating: 4, Service: "بازسازی کامل", Initial: "م", Text: "مبل قدیمی خونه رو کامل بازسازی کردن. کیفیت کار خیلی خوب بود. فقط یه کم بیشتر از چیزی که فکر می‌کردم طول کشید."},
			{Name: "نرگس رحیمی", City: "شیراز", Date: "۲ ماه پیش", Rating: 5, Service: "تعویض پارچه", Initial: "ن", Text: "از انتخاب پارچه تا تحویل نهایی، همه چیز حرفه‌ای بود. به همه دوستام معرفیشون کردم."},
		},
		FAQs: []FAQ{
			{Question: "هزینه تعمیر مبل چقدر است؟", Answer: "هزینه بستگی به نوع مبل و میزان کار دارد. برای اطلاع دقیق، کافیه عکس مبلتون رو واتساپ بفرستید تا در اسرع وقت قیمت دقیق بهتون اعلام بشه."},
			{Question: "آیا مبل رو از منزل ما می‌برید؟", Answer: "خیر، تعمیرات در کارگاه ما انجام می‌شود. شما مبل را به کارگاه ما می‌آورید و بعد از تعمیر تحویل می‌گیرید."},
			{Question: "چقدر طول می‌کشد تا کار تمام شود؟", Answer: "بسته به نوع کار، معمولاً بین ۱ تا ۵ روز. برای کارهای فوری هم امکان‌پذیره."},
			{Question: "آیا روی کارتون ضمانت می‌دید؟", Answer: "بله، تمام تعمیرات ۶ ماه ضمانت دارند. اگر مشکلی پیش بیاید، رایگان رفع می‌شود."},
			{Question: "چه پارچه‌هایی برای مبل استفاده می‌کنید؟", Answer: "از بهترین پارچه‌های ایرانی و خارجی استفاده می‌کنیم. نمونه‌ها رو می‌تونید به صورت حضوری در کارگاه ببینید."},
		},
		QandAs: []QandA{
			{ID: 1, Name: "شاپور زارع", Date: "۱ روز پیش", Question: "سلام وقت بخیر من پایه ی زیر مبلام کج شده داره میشکنه باید دلبشه هزینه اش به چه صورت هست؟", Replies: []QReply{{Author: "تعمیرات مبل قهرمانی", Date: "۱ روز پیش", Text: "سلام. برای تعمیر پایه مبل، بستگی به نوع مبل و شدت شکستگی داره. اگه عکس مبل رو واتساپ بفرستید، دقیق راهنماییتون می‌کنم و قیمت رو اعلام می‌کنم."}}},
			{ID: 2, Name: "مریم مومنی", Date: "۹ روز پیش", Question: "سلام مبل تازه خریداری کردم که بعضی جاهایش زدگی و رنگ پریدگی داره امکان ترمیم وجود داره؟", Replies: []QReply{{Author: "تعمیرات مبل قهرمانی", Date: "۹ روز پیش", Text: "سلام. بله، ترمیم رنگ و زدگی مبل انجام می‌شه. کافیه عکس محل زدگی رو بفرستید تا بررسی کنم و هزینه رو بگم."}}},
			{ID: 3, Name: "حسین احمدی", Date: "۵ روز پیش", Question: "برای مبل راحتی که فنرهاش خورده شده چقدر هزینه میگیرید؟ پارچه خودم دارم.", Replies: []QReply{{Author: "تعمیرات مبل قهرمانی", Date: "۵ روز پیش", Text: "سلام. تعویض فنر مبل راحتی بسته به تعداد فنرها و نوع مبل متفاوته. اگه پارچه خودتون باشه، هزینه کمتر می‌شه. عکس بفرستید تا دقیق بگم."}}},
		},
	}

	funcMap := template.FuncMap{
		"loop": func(n int) []int {
			result := make([]int, n)
			for i := 0; i < n; i++ {
				result[i] = i
			}
			return result
		},
		"stars": func(rating int) template.HTML {
			html := ""
			for i := 0; i < 5; i++ {
				if i < rating {
					html += `<span class="star filled">★</span>`
				} else {
					html += `<span class="star empty">★</span>`
				}
			}
			return template.HTML(html)
		},
	}

	// ساخت index.html
	tmpl, err := template.New("index").Funcs(funcMap).Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath.Join("docs", "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(f, data); err != nil {
		log.Fatal(err)
	}
	f.Close()

	// ساخت gallery-restoration.html
	restData := RestorationData{
		Name:      data.Name,
		Phone:     data.Phone,
		PhoneRaw:  data.PhoneRaw,
		Whatsapp:  data.Whatsapp,
		Instagram: data.Instagram,
		Images:    data.RestorationImages,
	}
	tmpl2, err := template.New("restoration").Parse(restorationTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f2, err := os.Create(filepath.Join("docs", "gallery-restoration.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl2.Execute(f2, restData); err != nil {
		log.Fatal(err)
	}
	f2.Close()

	// ساخت gallery-parche.html
	parcheData := RestorationData{
		Name:      data.Name,
		Phone:     data.Phone,
		PhoneRaw:  data.PhoneRaw,
		Whatsapp:  data.Whatsapp,
		Instagram: data.Instagram,
		Images:    data.ParcheImages,
	}
	tmpl3, err := template.New("parche").Parse(parcheTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f3, err := os.Create(filepath.Join("docs", "gallery-parche.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl3.Execute(f3, parcheData); err != nil {
		log.Fatal(err)
	}
	f3.Close()

	// ساخت gallery-sandali.html
	sandaliData := RestorationData{
		Name:      data.Name,
		Phone:     data.Phone,
		PhoneRaw:  data.PhoneRaw,
		Whatsapp:  data.Whatsapp,
		Instagram: data.Instagram,
		Images:    data.SandaliImages,
	}
	tmpl4, err := template.New("sandali").Parse(sandaliTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f4, err := os.Create(filepath.Join("docs", "gallery-sandali.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl4.Execute(f4, sandaliData); err != nil {
		log.Fatal(err)
	}
	f4.Close()
	log.Println("✅ docs/index.html و docs/gallery-restoration.html ساخته شدند")
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=5.0">
    <meta name="theme-color" content="#00bfa5">
    <title>{{.Title}}</title>
    <link rel="icon" type="image/png" href="images/logo.png">
    <meta name="description" content="تعمیرات مبل در شیراز - تعمیر مبل استیل، چستر، راحتی، تعویض پارچه و روکش مبل در کارگاه تخصصی. با ۷ سال سابقه و ضمانت ۶ ماهه. تماس: {{.Phone}}">
    <meta name="keywords" content="تعمیرات مبل شیراز, تعمیر مبل, کارگاه تعمیر مبل, تعمیر مبل استیل, تعویض پارچه مبل, روکش مبل, تعمیر صندلی شیراز, بازسازی مبل شیراز">
    <meta name="author" content="تعمیرات مبل قهرمانی">
    <meta name="robots" content="index, follow">
    <link rel="canonical" href="https://moblshiraz.ir/">
    <meta property="og:title" content="{{.Title}}">
    <meta property="og:description" content="{{.About}}">
    <meta property="og:type" content="website">
    <meta property="og:url" content="https://moblshiraz.ir/">
    <meta property="og:locale" content="fa_IR">
    <link href="https://cdn.jsdelivr.net/gh/rastikerdar/vazirmatn@v33.003/Vazirmatn-font-face.css" rel="stylesheet">
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "LocalBusiness",
      "name": "تعمیرات مبل قهرمانی",
      "telephone": "{{.PhoneRaw}}",
      "address": {"@type": "PostalAddress", "addressLocality": "شیراز", "addressRegion": "فارس", "addressCountry": "IR"},
      "areaServed": "شیراز و حومه",
      "description": "{{.About}}",
      "priceRange": "200000 - 5000000",
      "openingHours": "Sa-Th 09:00-20:00",
      "aggregateRating": {"@type": "AggregateRating", "ratingValue": "{{.Rating}}", "reviewCount": "{{.ReviewsNum}}"}
    }
    </script>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; font-family: 'Vazirmatn', Tahoma, sans-serif; -webkit-tap-highlight-color: transparent; }
        :root { --primary: #00bfa5; --primary-dark: #009688; --primary-light: #e0f2f1; --bg: #f8f9fa; --text: #212529; --text-sec: #6c757d; --border: #e9ecef; --yellow: #ffd54f; }
        html { scroll-behavior: smooth; }
        body { background: var(--bg); color: var(--text); line-height: 1.8; font-size: 14px; }

        .top-bar { background: white; padding: 12px 16px; border-bottom: 1px solid var(--border); position: sticky; top: 0; z-index: 100; box-shadow: 0 2px 8px rgba(0,0,0,0.04); }
        .top-inner { display: flex; justify-content: space-between; align-items: center; max-width: 1100px; margin: auto; width: 100%; }
        .brand { display: flex; align-items: center; gap: 8px; font-weight: 800; font-size: 16px; color: var(--primary); }
        .brand-icon { width: 32px; height: 32px; border-radius: 10px; overflow: hidden; background: var(--primary); }
        .brand-icon img { width: 100%; height: 100%; object-fit: cover; }
        .top-actions { display: flex; gap: 8px; }
        .icon-btn { width: 38px; height: 38px; border-radius: 10px; display: flex; align-items: center; justify-content: center; text-decoration: none; font-size: 18px; transition: all 0.2s; }
        .icon-btn.phone { background: var(--primary-light); color: var(--primary); }
        .icon-btn.whatsapp { background: #e7f9ef; color: #25D366; }

        .hero { background: white; padding: 20px 16px 30px; }
        .hero-inner { max-width: 1100px; margin: auto; }
        .hero-title { font-size: 22px; font-weight: 900; line-height: 1.5; margin-bottom: 8px; }
        .hero-sub { font-size: 13px; color: var(--text-sec); margin-bottom: 20px; }
        .hero-image { width: 100%; aspect-ratio: 16/10; border-radius: 16px; overflow: hidden; margin-bottom: 20px; background: #f0f0f0; }
        .hero-image img { width: 100%; height: 100%; object-fit: cover; display: block; }
        .price-box { background: var(--primary-light); border-radius: 12px; padding: 14px 16px; margin-bottom: 16px; display: flex; justify-content: space-between; align-items: center; }
        .price-label { font-size: 12px; color: var(--text-sec); }
        .price-value { font-size: 15px; font-weight: 800; color: var(--primary); }

        .quick-actions { display: flex; gap: 10px; margin-bottom: 16px; }
        .quick-btn { flex: 1; padding: 14px; border-radius: 12px; text-decoration: none; text-align: center; font-size: 14px; font-weight: 700; display: flex; align-items: center; justify-content: center; gap: 6px; transition: all 0.2s; }
        .quick-btn.call { background: var(--primary); color: white; box-shadow: 0 8px 20px rgba(0,191,165,0.3); }
        .quick-btn.wa { background: #25D366; color: white; box-shadow: 0 8px 20px rgba(37,211,102,0.3); }

        .stats { background: white; padding: 20px 16px; border-top: 1px solid var(--border); }
        .stats-inner { max-width: 1100px; margin: auto; display: flex; justify-content: space-around; text-align: center; }
        .stat-item { flex: 1; }
        .stat-value { font-size: 18px; font-weight: 900; }
        .stat-label { font-size: 11px; color: var(--text-sec); margin-top: 4px; }
        .stat-item + .stat-item { border-right: 1px solid var(--border); }

        .tabs-wrap { position: sticky; top: 60px; z-index: 90; background: white; border-bottom: 1px solid var(--border); overflow-x: auto; scrollbar-width: none; -webkit-overflow-scrolling: touch; }
        .tabs-wrap::-webkit-scrollbar { display: none; }
        .tabs { display: flex; gap: 6px; padding: 10px 16px; max-width: 1100px; margin: auto; white-space: nowrap; }
        .tab { padding: 8px 16px; border-radius: 20px; font-size: 13px; font-weight: 600; background: transparent; color: var(--text-sec); border: 1.5px solid var(--border); cursor: pointer; white-space: nowrap; }
        .tab.active { background: var(--primary); color: white; border-color: var(--primary); }
       
        .tabs-spacer { height: 57px; display: block; }
        .section { padding: 24px 16px; max-width: 1100px; margin: auto; }
        .section-title { font-size: 18px; font-weight: 800; margin-bottom: 6px; display: flex; align-items: center; gap: 8px; }
        .section-title::before { content: ''; width: 4px; height: 20px; background: var(--primary); border-radius: 2px; }
        .section-sub { font-size: 12px; color: var(--text-sec); margin-bottom: 18px; }

        .services-list { display: flex; flex-direction: column; gap: 12px; }
        .service-card { background: white; border: 1px solid var(--border); border-radius: 16px; overflow: hidden; }
        .about-box { background: white; border: 1px solid var(--border); border-radius: 16px; overflow: hidden; }
        .about-image { width: 100%; aspect-ratio: 16/9; overflow: hidden; background: #f0f0f0; }
        .about-image img { width: 100%; height: 100%; object-fit: cover; display: block; }
        .about-text { padding: 20px; font-size: 13px; line-height: 2; }
        .service-image { width: 100%; aspect-ratio: 16/9; overflow: hidden; background: #f0f0f0; }
        .service-image img { width: 100%; height: 100%; object-fit: cover; display: block; }
        .service-body { padding: 16px; }
        .service-name { font-size: 15px; font-weight: 700; margin-bottom: 4px; }
        .service-desc { font-size: 12px; color: var(--text-sec); line-height: 1.7; margin-bottom: 6px; }
        .service-price { font-size: 13px; font-weight: 700; color: var(--primary); }

        .gallery-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 12px; }
        .gallery-item { border-radius: 14px; overflow: hidden; background: white; border: 1px solid var(--border); transition: transform 0.3s ease, box-shadow 0.3s ease, border-color 0.3s ease; cursor: pointer; position: relative; }
        .gallery-item:hover { transform: translateY(-5px); box-shadow: 0 12px 30px rgba(0,191,165,0.25); border-color: var(--primary); }
        .gallery-item::after { content: '🔍'; position: absolute; top: 10px; left: 10px; background: rgba(0,191,165,0.9); color: white; width: 34px; height: 34px; border-radius: 50%; display: flex; align-items: center; justify-content: center; font-size: 15px; opacity: 0; transition: opacity 0.3s; }
        .gallery-item:hover::after { opacity: 1; }
        .gallery-item.linked::after { content: '📂'; background: rgba(255,152,0,0.95); }
        .gallery-image { width: 100%; aspect-ratio: 1; overflow: hidden; background: #f0f0f0; }
        .gallery-image img { width: 100%; height: 100%; object-fit: cover; display: block; transition: transform 0.4s ease; }
        .gallery-item:hover .gallery-image img { transform: scale(1.1); }
        .gallery-title { font-size: 12px; color: var(--text-sec); font-weight: 600; padding: 10px; text-align: center; transition: color 0.3s; }
        .gallery-item:hover .gallery-title { color: var(--primary); font-weight: 700; }

        .lightbox { display: none; position: fixed; inset: 0; background: rgba(0,0,0,0.95); z-index: 1000; align-items: center; justify-content: center; padding: 20px; }
        .lightbox.active { display: flex; }
        .lightbox-inner { position: relative; max-width: 90vw; max-height: 90vh; display: flex; flex-direction: column; align-items: center; }
        .lightbox img { max-width: 100%; max-height: 80vh; object-fit: contain; border-radius: 12px; box-shadow: 0 20px 60px rgba(0,0,0,0.6); }
        .lightbox-caption { color: white; font-size: 15px; font-weight: 700; margin-top: 16px; text-align: center; }
        .lightbox-counter { color: rgba(255,255,255,0.7); font-size: 13px; margin-top: 6px; }
        .lightbox-close { position: absolute; top: -10px; left: -10px; width: 44px; height: 44px; border-radius: 50%; background: rgba(255,255,255,0.15); color: white; border: none; font-size: 22px; cursor: pointer; display: flex; align-items: center; justify-content: center; backdrop-filter: blur(10px); }
        .lightbox-nav { position: absolute; top: 50%; transform: translateY(-50%); width: 50px; height: 50px; border-radius: 50%; background: rgba(255,255,255,0.15); color: white; border: none; font-size: 26px; cursor: pointer; display: flex; align-items: center; justify-content: center; backdrop-filter: blur(10px); }
        .lightbox-prev { right: -60px; }
        .lightbox-next { left: -60px; }
        @media (max-width: 768px) { .lightbox-prev { right: 10px; } .lightbox-next { left: 10px; } .lightbox-close { top: 10px; left: 10px; } }

        .review-card { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 16px; margin-bottom: 12px; }
        .review-header { display: flex; align-items: center; gap: 10px; margin-bottom: 10px; }
        .review-avatar { width: 40px; height: 40px; border-radius: 50%; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-weight: 800; font-size: 16px; flex-shrink: 0; }
        .review-name { font-size: 13px; font-weight: 700; }
        .review-meta { font-size: 11px; color: var(--text-sec); }
        .review-text { font-size: 13px; line-height: 1.9; }
        .review-service { display: inline-block; margin-top: 10px; padding: 4px 10px; border-radius: 8px; background: var(--primary-light); color: var(--primary); font-size: 11px; font-weight: 600; }
        .review-stars { font-size: 16px; margin-bottom: 8px; letter-spacing: 2px; line-height: 1; }
        .review-stars .star.filled { color: #ffd54f; }
        .review-stars .star.empty { color: #d1d5db; }

        .review-hidden { display: none; }
        .review-hidden.show { display: block; }
        .show-more-btn { display: block; width: 100%; padding: 14px; background: white; border: 1.5px solid var(--primary); color: var(--primary); border-radius: 14px; font-size: 14px; font-weight: 700; cursor: pointer; font-family: inherit; margin-top: 8px; }

        .qa-section { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 16px; }
        .qa-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px; }
        .qa-count { font-size: 14px; font-weight: 800; }
        .qa-count span { color: var(--primary); }
        .qa-ask-btn { background: var(--primary); color: white; padding: 10px 16px; border-radius: 10px; font-size: 13px; font-weight: 700; text-decoration: none; display: flex; align-items: center; gap: 6px; }
        .qa-item { border-bottom: 1px solid var(--border); padding: 16px 0; }
        .qa-item:last-child { border-bottom: none; padding-bottom: 0; }
        .qa-item:first-child { padding-top: 0; }
        .qa-meta { font-size: 11px; color: var(--text-sec); margin-bottom: 6px; }
        .qa-meta .qa-author { font-weight: 700; color: var(--text); }
        .qa-question { font-size: 13px; font-weight: 600; line-height: 1.9; margin-bottom: 10px; }
        .qa-actions { display: flex; align-items: center; gap: 16px; font-size: 12px; }
        .qa-toggle { background: none; border: none; color: var(--primary); font-size: 12px; font-weight: 700; cursor: pointer; display: flex; align-items: center; gap: 4px; font-family: inherit; padding: 0; }
        .qa-toggle .arrow { transition: transform 0.3s; display: inline-block; font-size: 10px; }
        .qa-item.open .qa-toggle .arrow { transform: rotate(180deg); }
        .qa-reply-btn { background: none; border: none; color: var(--primary); font-size: 12px; font-weight: 700; cursor: pointer; font-family: inherit; padding: 0; text-decoration: none; }
        .qa-replies { max-height: 0; overflow: hidden; transition: max-height 0.4s ease; }
        .qa-item.open .qa-replies { max-height: 1000px; margin-top: 12px; }
        .qa-reply { background: var(--bg); border-radius: 12px; padding: 14px; margin-top: 10px; border-right: 3px solid var(--primary); }
        .qa-reply-header { display: flex; align-items: center; gap: 8px; margin-bottom: 8px; }
        .qa-reply-avatar { width: 32px; height: 32px; border-radius: 50%; background: var(--primary); color: white; display: flex; align-items: center; justify-content: center; font-size: 11px; font-weight: 800; }
        .qa-reply-name { font-size: 12px; font-weight: 700; }
        .qa-reply-date { font-size: 10px; color: var(--text-sec); }
        .qa-reply-text { font-size: 12.5px; line-height: 1.9; color: var(--text); }
        .qa-show-more { display: block; width: 100%; margin-top: 16px; padding: 12px; background: transparent; border: 1.5px solid var(--primary); color: var(--primary); border-radius: 12px; font-size: 13px; font-weight: 700; cursor: pointer; font-family: inherit; }

        .faq-item { background: white; border: 1px solid var(--border); border-radius: 14px; margin-bottom: 10px; overflow: hidden; }
        .faq-question { padding: 16px; font-size: 13px; font-weight: 700; display: flex; justify-content: space-between; align-items: center; cursor: pointer; gap: 10px; }
        .faq-icon { width: 24px; height: 24px; flex-shrink: 0; border-radius: 50%; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-size: 14px; transition: transform 0.3s; }
        .faq-item.open .faq-icon { transform: rotate(180deg); }
        .faq-answer { padding: 0 16px; max-height: 0; overflow: hidden; font-size: 12.5px; color: var(--text-sec); line-height: 1.9; transition: all 0.3s ease; }
        .faq-item.open .faq-answer { padding: 0 16px 16px; max-height: 500px; }

        .contact-box { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 20px; }
        .contact-item { display: flex; align-items: center; gap: 12px; padding: 14px 0; border-bottom: 1px solid var(--border); text-decoration: none; color: inherit; }
        .contact-item:last-child { border-bottom: none; }
        .contact-icon { width: 40px; height: 40px; border-radius: 12px; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-size: 18px; flex-shrink: 0; }
        .contact-icon.wa { background: #e7f9ef; color: #25D366; }
        .contact-icon.map { background: #fef3c7; color: #d97706; }
        .contact-info { flex: 1; }
        .contact-label { font-size: 11px; color: var(--text-sec); }
        .contact-value { font-size: 14px; font-weight: 700; }

        .neshan-map-container { margin-top: 20px; border-radius: 16px; overflow: hidden; border: 1px solid var(--border); background: white; }
        .neshan-map-container #neshan-map { width: 100%; height: 300px; z-index: 1; }
        .map-link { display: block; padding: 14px; text-align: center; background: var(--primary-light); color: var(--primary); text-decoration: none; font-size: 14px; font-weight: 700; }
        .map-link:active { background: var(--primary); color: white; }
        .trust-badges { display: flex; gap: 12px; margin-top: 16px; }
        .trust-badge { flex: 1; background: white; border: 1px solid var(--border); border-radius: 12px; padding: 14px 10px; text-align: center; }
        .trust-icon { font-size: 28px; margin-bottom: 6px; }
        .trust-text { font-size: 10px; color: var(--text-sec); line-height: 1.5; }

        .footer { background: white; border-top: 1px solid var(--border); padding: 30px 16px 100px; text-align: center; margin-top: 30px; }
        .footer-brand { font-size: 16px; font-weight: 800; color: var(--primary); margin-bottom: 8px; }
        .footer-text { font-size: 12px; color: var(--text-sec); margin-bottom: 16px; }
        .footer-social { display: flex; justify-content: center; gap: 12px; margin-bottom: 20px; }
        .social-btn { width: 40px; height: 40px; border-radius: 12px; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-size: 18px; text-decoration: none; }
        .footer-copy { font-size: 11px; color: #adb5bd; padding-top: 16px; border-top: 1px solid var(--border); }

        .float-actions { position: fixed; bottom: 20px; left: 20px; right: 20px; display: flex; gap: 10px; z-index: 99; }
        .float-btn { flex: 1; padding: 16px; border-radius: 16px; text-align: center; font-size: 15px; font-weight: 700; text-decoration: none; display: flex; align-items: center; justify-content: center; gap: 8px; box-shadow: 0 10px 30px rgba(0,0,0,0.2); }
        .float-btn.call { background: var(--primary); color: white; }
        .float-btn.wa { background: #25D366; color: white; }

        @media (min-width: 768px) {
            body { font-size: 15px; }
            .hero-title { font-size: 32px; }
            .hero { padding: 40px 20px 50px; }
            .hero-image { aspect-ratio: 21/9; }
            .section { padding: 40px 20px; }
            .section-title { font-size: 22px; }
            .services-list { display: grid; grid-template-columns: repeat(2, 1fr); gap: 16px; }
            .gallery-grid { grid-template-columns: repeat(3, 1fr); }
            .float-actions { left: auto; right: 30px; bottom: 30px; width: auto; }
            .float-btn { padding: 14px 24px; border-radius: 50px; font-size: 14px; }
            .stats-inner { max-width: 700px; }
            .stat-value { font-size: 22px; }
        }
    </style>
</head>
<body>

    <header class="top-bar">
        <div class="top-inner">
            <div class="brand">
                <div class="brand-icon"><img src="images/logo.png" alt="{{.Name}}"></div>
                <span>{{.Name}}</span>
            </div>
            <div class="top-actions">
                <a href="tel:{{.PhoneRaw}}" class="icon-btn phone" aria-label="تماس">📞</a>
                <a href="https://wa.me/{{.Whatsapp}}" class="icon-btn whatsapp" aria-label="واتساپ" target="_blank">💬</a>
            </div>
        </div>
    </header>

    <section class="hero">
        <div class="hero-inner">
            <h1 class="hero-title">{{.Slogan}}</h1>
            <p class="hero-sub">{{.About}}</p>
            <div class="hero-image"><img src="{{.HeroImage}}" alt="تعمیرات مبل در شیراز"></div>
            <div class="price-box">
                <span class="price-label">هزینه خدمات</span>
                <span class="price-value">از ۲۰۰ هزار تومان</span>
            </div>
            <div class="quick-actions">
                <a href="tel:{{.PhoneRaw}}" class="quick-btn call">📞 تماس فوری</a>
                <a href="https://wa.me/{{.Whatsapp}}" class="quick-btn wa" target="_blank">💬 واتساپ</a>
            </div>
        </div>
    </section>

    <section class="stats">
        <div class="stats-inner">
            <div class="stat-item"><div class="stat-value">{{.OrdersCount}}</div><div class="stat-label">سفارش موفق</div></div>
            <div class="stat-item"><div class="stat-value">{{.Rating}}</div><div class="stat-label">امتیاز ({{.ReviewsNum}} نظر)</div></div>
            <div class="stat-item"><div class="stat-value">{{.Experience}} سال</div><div class="stat-label">سابقه کار</div></div>
        </div>
    </section>

    <div class="tabs-wrap">
        <div class="tabs">
            <button class="tab active" data-tab="about">درباره ما</button>
            <button class="tab" data-tab="services">خدمات</button>
            <button class="tab" data-tab="gallery">نمونه کارها</button>
            <button class="tab" data-tab="qanda">پرسش و پاسخ</button>
            <button class="tab" data-tab="reviews">نظرات</button>
            <button class="tab" data-tab="faq">سوالات متداول</button>
            <button class="tab" data-tab="contact">تماس</button>
        </div>
    </div>

    <div class="tabs-spacer"></div>

    <section class="section" id="about">
        <h2 class="section-title">درباره ما</h2>
        <p class="section-sub">با بیش از ۷ سال تجربه در تعمیرات مبل در شیراز</p>
        <div class="about-box">
            <div class="about-image"><img src="{{.AboutImage}}" alt="کارگاه تعمیرات مبل قهرمانی در شیراز"></div>
            <div class="about-text"><p>{{.LongAbout}}</p></div>
        </div>
    </section>

    <section class="section" id="services">
        <h2 class="section-title">خدمات ما</h2>
        <p class="section-sub">با کیفیت‌ترین خدمات تعمیرات مبل در شیراز</p>
        <div class="services-list">
            {{range .Services}}
            <div class="service-card">
                <div class="service-image"><img src="{{.Image}}" alt="{{.Title}} در شیراز" loading="lazy"></div>
                <div class="service-body">
                    <div class="service-name">{{.Title}}</div>
                    <div class="service-desc">{{.Desc}}</div>
                    <div class="service-price">{{.Price}}</div>
                </div>
            </div>
            {{end}}
        </div>
    </section>

    <section class="section" id="gallery">
        <h2 class="section-title">نمونه کارها</h2>
        <p class="section-sub">روی هر عکس کلیک کنید</p>
        <div class="gallery-grid">
            {{range $index, $item := .Gallery}}
            <div class="gallery-item{{if $item.Link}} linked{{end}}" data-link="{{$item.Link}}" data-index="{{$index}}" onclick="handleGalleryClick(this)">
                <div class="gallery-image"><img src="{{$item.Image}}" alt="{{$item.Title}} در شیراز" loading="lazy"></div>
                <div class="gallery-title">{{$item.Title}}</div>
            </div>
            {{end}}
        </div>
    </section>

    <section class="section" id="qanda">
        <h2 class="section-title">پرسش و پاسخ</h2>
        <p class="section-sub">سوالات مشتریان و پاسخ‌های ما</p>
        <div class="qa-section">
            <div class="qa-header">
                <div class="qa-count"><span>{{len .QandAs}}</span> پرسش</div>
                <a href="https://wa.me/{{.Whatsapp}}" class="qa-ask-btn" target="_blank">+ ثبت پرسش</a>
            </div>
            {{range $index, $qa := .QandAs}}
            <div class="qa-item{{if eq $index 0}} open{{end}}">
                <div class="qa-meta"><span class="qa-author">{{$qa.Name}}</span> · {{$qa.Date}}</div>
                <div class="qa-question">{{$qa.Question}}</div>
                <div class="qa-actions">
                    <button class="qa-toggle" onclick="toggleQA(this)"><span class="arrow">▾</span><span>مشاهده {{len $qa.Replies}} پاسخ</span></button>
                    <a href="https://wa.me/{{$.Whatsapp}}" class="qa-reply-btn" target="_blank">ثبت پاسخ ←</a>
                </div>
                <div class="qa-replies">
                    {{range $qa.Replies}}
                    <div class="qa-reply">
                        <div class="qa-reply-header"><div class="qa-reply-avatar">{{slice .Author 0 3}}</div><div><div class="qa-reply-name">{{.Author}}</div><div class="qa-reply-date">{{.Date}}</div></div></div>
                        <div class="qa-reply-text">{{.Text}}</div>
                    </div>
                    {{end}}
                </div>
            </div>
            {{end}}
            <button class="qa-show-more" onclick="alert('به زودی...')">+ نمایش بیشتر</button>
        </div>
    </section>

    <section class="section" id="reviews">
        <h2 class="section-title">نظرات مشتریان</h2>
        <p class="section-sub">افتخار ما، رضایت شماست</p>
        {{range $index, $rev := .Reviews}}
        <div class="review-card{{if ge $index 3}} review-hidden{{end}}">
            <div class="review-header">
                <div class="review-avatar">{{$rev.Initial}}</div>
                <div><div class="review-name">{{$rev.Name}}</div><div class="review-meta">{{$rev.City}} · {{$rev.Date}}</div></div>
            </div>
            <div class="review-stars">{{stars $rev.Rating}}</div>
            <p class="review-text">{{$rev.Text}}</p>
            <span class="review-service">{{$rev.Service}}</span>
        </div>
        {{end}}
        {{if gt (len .Reviews) 3}}
        <button class="show-more-btn" onclick="showMoreReviews(this)">+ نمایش نظرات بیشتر ({{len .Reviews}} نظر)</button>
        {{end}}
    </section>

    <section class="section" id="faq">
        <h2 class="section-title">سوالات متداول</h2>
        <p class="section-sub">پاسخ سوالات پرتکرار شما</p>
        {{range .FAQs}}
        <div class="faq-item">
            <div class="faq-question" onclick="toggleFaq(this)"><span>{{.Question}}</span><span class="faq-icon">▾</span></div>
            <div class="faq-answer">{{.Answer}}</div>
        </div>
        {{end}}
    </section>

    <section class="section" id="contact">
        <h2 class="section-title">تماس با ما</h2>
        <p class="section-sub">برای مشاوره و سفارش در تماس باشید</p>
        <div class="contact-box">
            <a href="tel:{{.PhoneRaw}}" class="contact-item"><div class="contact-icon">📞</div><div class="contact-info"><div class="contact-label">تماس تلفنی</div><div class="contact-value">{{.Phone}}</div></div></a>
            <a href="https://wa.me/{{.Whatsapp}}" class="contact-item" target="_blank"><div class="contact-icon wa">💬</div><div class="contact-info"><div class="contact-label">واتساپ</div><div class="contact-value">ارسال پیام در واتساپ</div></div></a>
            <a href="https://instagram.com/{{.Instagram}}" class="contact-item" target="_blank"><div class="contact-icon">📷</div><div class="contact-info"><div class="contact-label">اینستاگرام</div><div class="contact-value">@{{.Instagram}}</div></div></a>
            <div class="contact-item"><div class="contact-icon map">📍</div><div class="contact-info"><div class="contact-label">آدرس کارگاه</div><div class="contact-value">{{.Address}}</div></div></div>
        </div>
        <a href="https://nshn.ir/be_bgLzhPFQAyo" target="_blank" style="display:flex; align-items:center; justify-content:center; gap:10px; background:linear-gradient(135deg, #00bfa5, #009688); color:white; text-align:center; padding:18px; border-radius:14px; font-weight:bold; text-decoration:none; margin-top:20px; box-shadow:0 8px 20px rgba(0,191,165,0.3); font-family:tahoma;">
    📍 مشاهده آدرس کارگاه روی نقشه نشان
</a>
            <div class="trust-badges">
            <div class="trust-badge"><div class="trust-icon">✅</div><div class="trust-text">ضمانت ۶ ماهه</div></div>
            <div class="trust-badge"><div class="trust-icon">🏆</div><div class="trust-text">۷ سال سابقه</div></div>
            <div class="trust-badge"><div class="trust-icon">💰</div><div class="trust-text">قیمت منصفانه</div></div>
        </div>
    </section>

    <footer class="footer">
        <div class="footer-brand">{{.Name}}</div>
        <p class="footer-text">{{.Slogan}}</p>
        <div class="footer-social">
            <a href="tel:{{.PhoneRaw}}" class="social-btn" aria-label="تماس">📞</a>
            <a href="https://wa.me/{{.Whatsapp}}" class="social-btn" target="_blank" aria-label="واتساپ">💬</a>
            <a href="https://instagram.com/{{.Instagram}}" class="social-btn" target="_blank" aria-label="اینستاگرام">📷</a>
        </div>
        <p class="footer-copy">© ۱۴۰۳ {{.Name}} — تعمیرات مبل در شیراز</p>
    </footer>

    <div class="float-actions">
        <a href="tel:{{.PhoneRaw}}" class="float-btn call">📞 تماس</a>
        <a href="https://wa.me/{{.Whatsapp}}" class="float-btn wa" target="_blank">💬 واتساپ</a>
    </div>

    <div class="lightbox" id="lightbox" onclick="if(event.target===this)closeLightbox()">
        <div class="lightbox-inner">
            <button class="lightbox-close" onclick="closeLightbox()" aria-label="بستن">✕</button>
            <button class="lightbox-nav lightbox-prev" onclick="prevImage()" aria-label="قبلی">›</button>
            <img id="lightbox-img" src="" alt="">
            <button class="lightbox-nav lightbox-next" onclick="nextImage()" aria-label="بعدی">‹</button>
            <div class="lightbox-caption" id="lightbox-caption"></div>
            <div class="lightbox-counter" id="lightbox-counter"></div>
        </div>
    </div>

    <script>
        function toggleFaq(el) { el.parentElement.classList.toggle('open'); }
        function toggleQA(el) { el.closest('.qa-item').classList.toggle('open'); }
        function showMoreReviews(btn) { document.querySelectorAll('.review-hidden').forEach(el => el.classList.add('show')); btn.style.display = 'none'; }
        document.querySelectorAll('.tab').forEach(tab => {
            tab.addEventListener('click', () => {
                document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
                tab.classList.add('active');
                const el = document.getElementById(tab.dataset.tab);
                if (el) el.scrollIntoView({ behavior: 'smooth', block: 'start' });
            });
        });

        const galleryImages = [{{range .Gallery}}{ src: '{{.Image}}', title: '{{.Title}}' },{{end}}];
        let currentIndex = 0;

        function handleGalleryClick(el) {
            const link = el.dataset.link;
            if (link && link !== '') { window.location.href = link; }
            else { openLightbox(parseInt(el.dataset.index)); }
        }

        function openLightbox(index) { currentIndex = index; updateLightbox(); document.getElementById('lightbox').classList.add('active'); document.body.style.overflow = 'hidden'; }
        function closeLightbox() { document.getElementById('lightbox').classList.remove('active'); document.body.style.overflow = ''; }
        function nextImage() { currentIndex = (currentIndex + 1) % galleryImages.length; updateLightbox(); }
        function prevImage() { currentIndex = (currentIndex - 1 + galleryImages.length) % galleryImages.length; updateLightbox(); }
        function updateLightbox() {
            const item = galleryImages[currentIndex];
            document.getElementById('lightbox-img').src = item.src;
            document.getElementById('lightbox-caption').textContent = item.title;
            document.getElementById('lightbox-counter').textContent = (currentIndex + 1) + ' از ' + galleryImages.length;
        }
        document.addEventListener('keydown', (e) => {
            if (!document.getElementById('lightbox').classList.contains('active')) return;
            if (e.key === 'Escape') closeLightbox();
            if (e.key === 'ArrowLeft') nextImage();
            if (e.key === 'ArrowRight') prevImage();
        });
    </script>
</body>
</html>`

const sandaliTemplate = `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="theme-color" content="#00bfa5">
    <title>گالری تعمیر صندلی | {{.Name}}</title>
    <meta name="description" content="نمونه کارهای تعمیر صندلی در کارگاه تعمیرات مبل قهرمانی شیراز">
    <link href="https://cdn.jsdelivr.net/gh/rastikerdar/vazirmatn@v33.003/Vazirmatn-font-face.css" rel="stylesheet">
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; font-family: 'Vazirmatn', Tahoma, sans-serif; -webkit-tap-highlight-color: transparent; }
        :root { --primary: #00bfa5; --primary-light: #e0f2f1; --bg: #f8f9fa; --text: #212529; --text-sec: #6c757d; --border: #e9ecef; }
        body { background: var(--bg); color: var(--text); line-height: 1.8; font-size: 14px; }

        .top-bar { background: white; padding: 12px 16px; border-bottom: 1px solid var(--border); position: sticky; top: 0; z-index: 100; box-shadow: 0 2px 8px rgba(0,0,0,0.04); }
        .top-inner { display: flex; justify-content: space-between; align-items: center; max-width: 1100px; margin: auto; }
        .back-btn { display: flex; align-items: center; gap: 6px; text-decoration: none; color: var(--primary); font-weight: 700; font-size: 14px; padding: 8px 14px; border-radius: 10px; background: var(--primary-light); }
        .brand { font-weight: 800; font-size: 15px; color: var(--text); }

        .header { background: white; padding: 30px 16px; text-align: center; border-bottom: 1px solid var(--border); }
        .header h1 { font-size: 24px; font-weight: 900; margin-bottom: 8px; }
        .header p { font-size: 13px; color: var(--text-sec); }

        .gallery-container { max-width: 1100px; margin: auto; padding: 24px 16px 100px; }
        .gallery-grid { display: grid; grid-template-columns: repeat(1, 1fr); gap: 20px; }
        .gallery-item { border-radius: 16px; overflow: hidden; background: white; border: 1px solid var(--border); transition: transform 0.3s ease, box-shadow 0.3s ease; cursor: pointer; box-shadow: 0 4px 20px rgba(0,0,0,0.04); }
        .gallery-item:hover { transform: translateY(-4px); box-shadow: 0 15px 35px rgba(0,191,165,0.2); }
        .gallery-image { width: 100%; overflow: hidden; background: #f0f0f0; }
        .gallery-image img { width: 100%; height: auto; display: block; transition: transform 0.4s ease; }
        .gallery-item:hover .gallery-image img { transform: scale(1.03); }
        .gallery-title { font-size: 14px; color: var(--text); font-weight: 700; padding: 14px; text-align: center; }

        .empty-msg { text-align: center; padding: 60px 20px; color: var(--text-sec); background: white; border-radius: 16px; }

        .lightbox { display: none; position: fixed; inset: 0; background: rgba(0,0,0,0.95); z-index: 1000; align-items: center; justify-content: center; padding: 20px; }
        .lightbox.active { display: flex; }
        .lightbox-inner { position: relative; max-width: 90vw; max-height: 90vh; display: flex; flex-direction: column; align-items: center; }
        .lightbox img { max-width: 100%; max-height: 85vh; object-fit: contain; border-radius: 12px; box-shadow: 0 20px 60px rgba(0,0,0,0.6); }
        .lightbox-caption { color: white; font-size: 15px; font-weight: 700; margin-top: 16px; }
        .lightbox-close { position: absolute; top: -10px; left: -10px; width: 44px; height: 44px; border-radius: 50%; background: rgba(255,255,255,0.15); color: white; border: none; font-size: 22px; cursor: pointer; display: flex; align-items: center; justify-content: center; backdrop-filter: blur(10px); }

        @media (min-width: 768px) {
            .gallery-grid { grid-template-columns: repeat(2, 1fr); gap: 24px; }
            .header h1 { font-size: 32px; }
        }
    </style>
</head>
<body>

    <header class="top-bar">
        <div class="top-inner">
            <a href="index.html" class="back-btn">← بازگشت به سایت</a>
            <div class="brand">🛋️ {{.Name}}</div>
        </div>
    </header>

    <div class="header">
        <h1>گالری تعمیر صندلی</h1>
        <p>نمونه‌کارهای تعمیر صندلی در کارگاه ما (قبل و بعد)</p>
    </div>

    <div class="gallery-container">
        {{if .Images}}
        <div class="gallery-grid">
            {{range $index, $item := .Images}}
            <div class="gallery-item" onclick="openLightbox({{$index}})">
                <div class="gallery-image"><img src="{{$item.Image}}" alt="{{$item.Title}}" loading="lazy"></div>
                <div class="gallery-title">{{$item.Title}}</div>
            </div>
            {{end}}
        </div>
        {{else}}
        <div class="empty-msg">📭 هنوز عکسی برای این گالری اضافه نشده است.</div>
        {{end}}
    </div>

    <div class="lightbox" id="lightbox" onclick="if(event.target===this)closeLightbox()">
        <div class="lightbox-inner">
            <button class="lightbox-close" onclick="closeLightbox()">✕</button>
            <img id="lightbox-img" src="" alt="">
            <div class="lightbox-caption" id="lightbox-caption"></div>
        </div>
    </div>

    <script>
        const images = [{{range .Images}}{ src: '{{.Image}}', title: '{{.Title}}' },{{end}}];
        function openLightbox(index) {
            document.getElementById('lightbox-img').src = images[index].src;
            document.getElementById('lightbox-caption').textContent = images[index].title;
            document.getElementById('lightbox').classList.add('active');
            document.body.style.overflow = 'hidden';
        }
        function closeLightbox() {
            document.getElementById('lightbox').classList.remove('active');
            document.body.style.overflow = '';
        }
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') closeLightbox();
        });
    </script>
</body>
</html>`

const parcheTemplate = `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="theme-color" content="#00bfa5">
    <title>گالری تعویض پارچه مبل | {{.Name}}</title>
    <meta name="description" content="نمونه کارهای تعویض پارچه مبل در کارگاه تعمیرات مبل قهرمانی شیراز">
    <link href="https://cdn.jsdelivr.net/gh/rastikerdar/vazirmatn@v33.003/Vazirmatn-font-face.css" rel="stylesheet">
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; font-family: 'Vazirmatn', Tahoma, sans-serif; -webkit-tap-highlight-color: transparent; }
        :root { --primary: #00bfa5; --primary-light: #e0f2f1; --bg: #f8f9fa; --text: #212529; --text-sec: #6c757d; --border: #e9ecef; }
        body { background: var(--bg); color: var(--text); line-height: 1.8; font-size: 14px; }

        .top-bar { background: white; padding: 12px 16px; border-bottom: 1px solid var(--border); position: sticky; top: 0; z-index: 100; box-shadow: 0 2px 8px rgba(0,0,0,0.04); }
        .top-inner { display: flex; justify-content: space-between; align-items: center; max-width: 1100px; margin: auto; }
        .back-btn { display: flex; align-items: center; gap: 6px; text-decoration: none; color: var(--primary); font-weight: 700; font-size: 14px; padding: 8px 14px; border-radius: 10px; background: var(--primary-light); }
        .brand { font-weight: 800; font-size: 15px; color: var(--text); }

        .header { background: white; padding: 30px 16px; text-align: center; border-bottom: 1px solid var(--border); }
        .header h1 { font-size: 24px; font-weight: 900; margin-bottom: 8px; }
        .header p { font-size: 13px; color: var(--text-sec); }

        .gallery-container { max-width: 1100px; margin: auto; padding: 24px 16px 100px; }
        .gallery-grid { display: grid; grid-template-columns: repeat(1, 1fr); gap: 20px; }
        .gallery-item { border-radius: 16px; overflow: hidden; background: white; border: 1px solid var(--border); transition: transform 0.3s ease, box-shadow 0.3s ease; cursor: pointer; box-shadow: 0 4px 20px rgba(0,0,0,0.04); }
        .gallery-item:hover { transform: translateY(-4px); box-shadow: 0 15px 35px rgba(0,191,165,0.2); }
        .gallery-image { width: 100%; overflow: hidden; background: #f0f0f0; }
        .gallery-image img { width: 100%; height: auto; display: block; transition: transform 0.4s ease; }
        .gallery-item:hover .gallery-image img { transform: scale(1.03); }
        .gallery-title { font-size: 14px; color: var(--text); font-weight: 700; padding: 14px; text-align: center; }

        .empty-msg { text-align: center; padding: 60px 20px; color: var(--text-sec); background: white; border-radius: 16px; }

        .lightbox { display: none; position: fixed; inset: 0; background: rgba(0,0,0,0.95); z-index: 1000; align-items: center; justify-content: center; padding: 20px; }
        .lightbox.active { display: flex; }
        .lightbox-inner { position: relative; max-width: 90vw; max-height: 90vh; display: flex; flex-direction: column; align-items: center; }
        .lightbox img { max-width: 100%; max-height: 85vh; object-fit: contain; border-radius: 12px; box-shadow: 0 20px 60px rgba(0,0,0,0.6); }
        .lightbox-caption { color: white; font-size: 15px; font-weight: 700; margin-top: 16px; }
        .lightbox-close { position: absolute; top: -10px; left: -10px; width: 44px; height: 44px; border-radius: 50%; background: rgba(255,255,255,0.15); color: white; border: none; font-size: 22px; cursor: pointer; display: flex; align-items: center; justify-content: center; backdrop-filter: blur(10px); }

        @media (min-width: 768px) {
            .gallery-grid { grid-template-columns: repeat(2, 1fr); gap: 24px; }
            .header h1 { font-size: 32px; }
        }
    </style>
</head>
<body>

    <header class="top-bar">
        <div class="top-inner">
            <a href="index.html" class="back-btn">← بازگشت به سایت</a>
            <div class="brand">🛋️ {{.Name}}</div>
        </div>
    </header>

    <div class="header">
        <h1>گالری تعویض پارچه مبل</h1>
        <p>نمونه‌کارهای تعویض پارچه مبل در کارگاه ما (قبل و بعد)</p>
    </div>

    <div class="gallery-container">
        {{if .Images}}
        <div class="gallery-grid">
            {{range $index, $item := .Images}}
            <div class="gallery-item" onclick="openLightbox({{$index}})">
                <div class="gallery-image"><img src="{{$item.Image}}" alt="{{$item.Title}}" loading="lazy"></div>
                <div class="gallery-title">{{$item.Title}}</div>
            </div>
            {{end}}
        </div>
        {{else}}
        <div class="empty-msg">📭 هنوز عکسی برای این گالری اضافه نشده است.</div>
        {{end}}
    </div>

    <div class="lightbox" id="lightbox" onclick="if(event.target===this)closeLightbox()">
        <div class="lightbox-inner">
            <button class="lightbox-close" onclick="closeLightbox()">✕</button>
            <img id="lightbox-img" src="" alt="">
            <div class="lightbox-caption" id="lightbox-caption"></div>
        </div>
    </div>

    <script>
        const images = [{{range .Images}}{ src: '{{.Image}}', title: '{{.Title}}' },{{end}}];
        function openLightbox(index) {
            document.getElementById('lightbox-img').src = images[index].src;
            document.getElementById('lightbox-caption').textContent = images[index].title;
            document.getElementById('lightbox').classList.add('active');
            document.body.style.overflow = 'hidden';
        }
        function closeLightbox() {
            document.getElementById('lightbox').classList.remove('active');
            document.body.style.overflow = '';
        }
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') closeLightbox();
        });
    </script>
</body>
</html>`

const restorationTemplate = `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="theme-color" content="#00bfa5">
    <title>گالری بازسازی کامل مبل | {{.Name}}</title>
    <meta name="description" content="نمونه کارهای بازسازی کامل مبل در کارگاه تعمیرات مبل قهرمانی شیراز">
    <link href="https://cdn.jsdelivr.net/gh/rastikerdar/vazirmatn@v33.003/Vazirmatn-font-face.css" rel="stylesheet">
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; font-family: 'Vazirmatn', Tahoma, sans-serif; -webkit-tap-highlight-color: transparent; }
        :root { --primary: #00bfa5; --primary-light: #e0f2f1; --bg: #f8f9fa; --text: #212529; --text-sec: #6c757d; --border: #e9ecef; }
        body { background: var(--bg); color: var(--text); line-height: 1.8; font-size: 14px; }

        .top-bar { background: white; padding: 12px 16px; border-bottom: 1px solid var(--border); position: sticky; top: 0; z-index: 100; box-shadow: 0 2px 8px rgba(0,0,0,0.04); }
        .top-inner { display: flex; justify-content: space-between; align-items: center; max-width: 1100px; margin: auto; }
        .back-btn { display: flex; align-items: center; gap: 6px; text-decoration: none; color: var(--primary); font-weight: 700; font-size: 14px; padding: 8px 14px; border-radius: 10px; background: var(--primary-light); }
        .brand { font-weight: 800; font-size: 15px; color: var(--text); }

        .header { background: white; padding: 30px 16px; text-align: center; border-bottom: 1px solid var(--border); }
        .header h1 { font-size: 24px; font-weight: 900; margin-bottom: 8px; }
        .header p { font-size: 13px; color: var(--text-sec); }

        .gallery-container { max-width: 1100px; margin: auto; padding: 24px 16px 100px; }
        .gallery-grid { display: grid; grid-template-columns: repeat(1, 1fr); gap: 20px; }
        .gallery-item { border-radius: 16px; overflow: hidden; background: white; border: 1px solid var(--border); transition: transform 0.3s ease, box-shadow 0.3s ease; cursor: pointer; box-shadow: 0 4px 20px rgba(0,0,0,0.04); }
        .gallery-item:hover { transform: translateY(-4px); box-shadow: 0 15px 35px rgba(0,191,165,0.2); }
        .gallery-image { width: 100%; overflow: hidden; background: #f0f0f0; }
        .gallery-image img { width: 100%; height: auto; display: block; transition: transform 0.4s ease; }
        .gallery-item:hover .gallery-image img { transform: scale(1.03); }
        .gallery-title { font-size: 14px; color: var(--text); font-weight: 700; padding: 14px; text-align: center; }

        .empty-msg { text-align: center; padding: 60px 20px; color: var(--text-sec); background: white; border-radius: 16px; }

        .lightbox { display: none; position: fixed; inset: 0; background: rgba(0,0,0,0.95); z-index: 1000; align-items: center; justify-content: center; padding: 20px; }
        .lightbox.active { display: flex; }
        .lightbox-inner { position: relative; max-width: 90vw; max-height: 90vh; display: flex; flex-direction: column; align-items: center; }
        .lightbox img { max-width: 100%; max-height: 85vh; object-fit: contain; border-radius: 12px; box-shadow: 0 20px 60px rgba(0,0,0,0.6); }
        .lightbox-caption { color: white; font-size: 15px; font-weight: 700; margin-top: 16px; }
        .lightbox-close { position: absolute; top: -10px; left: -10px; width: 44px; height: 44px; border-radius: 50%; background: rgba(255,255,255,0.15); color: white; border: none; font-size: 22px; cursor: pointer; display: flex; align-items: center; justify-content: center; backdrop-filter: blur(10px); }

        @media (min-width: 768px) {
            .gallery-grid { grid-template-columns: repeat(2, 1fr); gap: 24px; }
            .header h1 { font-size: 32px; }
        }
    </style>
</head>
<body>

    <header class="top-bar">
        <div class="top-inner">
            <a href="index.html" class="back-btn">← بازگشت به سایت</a>
            <div class="brand">🛋️ {{.Name}}</div>
        </div>
    </header>

    <div class="header">
        <h1>گالری بازسازی کامل مبل</h1>
        <p>نمونه‌کارهای بازسازی کامل مبل در کارگاه ما (قبل و بعد)</p>
    </div>

    <div class="gallery-container">
        {{if .Images}}
        <div class="gallery-grid">
            {{range $index, $item := .Images}}
            <div class="gallery-item" onclick="openLightbox({{$index}})">
                <div class="gallery-image"><img src="{{$item.Image}}" alt="{{$item.Title}}" loading="lazy"></div>
                <div class="gallery-title">{{$item.Title}}</div>
            </div>
            {{end}}
        </div>
        {{else}}
        <div class="empty-msg">📭 هنوز عکسی برای این گالری اضافه نشده است.</div>
        {{end}}
    </div>

    <div class="lightbox" id="lightbox" onclick="if(event.target===this)closeLightbox()">
        <div class="lightbox-inner">
            <button class="lightbox-close" onclick="closeLightbox()">✕</button>
            <img id="lightbox-img" src="" alt="">
            <div class="lightbox-caption" id="lightbox-caption"></div>
        </div>
    </div>

    <script>
        const images = [{{range .Images}}{ src: '{{.Image}}', title: '{{.Title}}' },{{end}}];
        function openLightbox(index) {
            document.getElementById('lightbox-img').src = images[index].src;
            document.getElementById('lightbox-caption').textContent = images[index].title;
            document.getElementById('lightbox').classList.add('active');
            document.body.style.overflow = 'hidden';
        }
        function closeLightbox() {
            document.getElementById('lightbox').classList.remove('active');
            document.body.style.overflow = '';
        }
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') closeLightbox();
        });
    </script>
</body>
</html>`
