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
	Title       string
	Name        string
	Slogan      string
	About       string
	LongAbout   string
	Phone       string
	PhoneRaw    string
	Whatsapp    string
	Address     string
	AddressRaw  string
	MapEmbed    string
	MapLink     string
	Instagram   string
	Experience  string
	OrdersCount string
	Rating      string
	ReviewsNum  string
	HeroImage   string
	AboutImage  string
	Services    []Service
	Gallery     []GalleryItem
	Reviews     []Review
	FAQs        []FAQ
	QandAs      []QandA
}

func main() {
	if err := os.MkdirAll("docs", 0755); err != nil {
		log.Fatal(err)
	}

	data := SiteData{
		Title:       "تعمیرات مبل شیراز",
		Name:        "تعمیرات مبل شیراز",
		Slogan:      "تعمیر تخصصی انواع مبل، صندلی و پرسی در کارگاه ما",
		About:       "با بیش از ۷ سال تجربه، مبل شما را در کارگاه تخصصی‌مان تعمیر می‌کنیم",
		LongAbout:   "کارگاه تعمیرات مبل قهرمانی با بیش از ۷ سال سابقه در زمینه تعمیر و بازسازی انواع مبل، صندلی و پرسی در شیراز فعالیت می‌کند. ما با بهره‌گیری از ابزارهای حرفه‌ای و مواد اولیه باکیفیت، مبل قدیمی شما را به روزهای اوجش برمی‌گردانیم. تخصص ما شامل تعمیر مبل استیل، چستر، راحتی، تعمیر صندلی اداری و غذاخوری، تعویض پارچه، روکش‌کشی، تعمیر فنر و اسکلت و بازسازی کامل مبل است. با ۷ سال تجربه و بیش از ۱۳۷ سفارش موفق، آماده خدمت‌رسانی به شما عزیزان هستیم. مبل خود را به کارگاه ما بیاورید و با کیفیتی بی‌نظیر و قیمتی منصفانه تحویل بگیرید.",
		Phone:       "09013643428",
		PhoneRaw:    "+989013643428",
		Whatsapp:    "989013643428",
		Address:     "شیراز، کارگاه تعمیرات مبل قهرمانی",
		AddressRaw:  "Shiraz",
		MapEmbed:    "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d111388.83368571!2d52.441!3d29.591!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0x0!2zU2hpcmF6!5e0!3m2!1sen!2s!4v1234567890",
		MapLink:     "https://www.google.com/maps/search/?api=1&query=Shiraz",
		Instagram:   "moblshz",
		Experience:  "۷",
		OrdersCount: "۱۳۷",
		Rating:      "۴.۹",
		ReviewsNum:  "۴۸۰",
		HeroImage:   "images/hero.jpg",
		AboutImage:  "images/about.jpg",
		Services: []Service{
			{Image: "images/service-1.jpg", Title: "تعمیر مبل استیل", Desc: "تعمیر و بازسازی کامل مبل‌های استیل در کارگاه", Price: "از ۵۰۰ هزار تومان"},
			{Image: "images/service-2.jpg", Title: "تعمیر صندلی", Desc: "تعمیر انواع صندلی اداری و غذاخوری", Price: "از ۲۰۰ هزار تومان"},
			{Image: "images/service-3.jpg", Title: "تعویض پارچه", Desc: "تعویض پارچه مبل با جدیدترین طرح‌ها", Price: "از ۸۰۰ هزار تومان"},
			{Image: "images/service-4.jpg", Title: "تعمیر فنر و اسکلت", Desc: "تعمیر اسکلت چوبی و فنرهای مبل", Price: "از ۴۰۰ هزار تومان"},
			{Image: "images/service-5.jpg", Title: "روکش مبل", Desc: "روکش‌کشی حرفه‌ای انواع مبل", Price: "از ۱ میلیون تومان"},
			{Image: "images/service-6.jpg", Title: "بازسازی کامل", Desc: "بازسازی صفر تا صد مبل قدیمی", Price: "از ۲ میلیون تومان"},
		},
		Gallery: []GalleryItem{
			{Image: "images/gallery-1.jpg", Title: "تعمیر مبل استیل"},
			{Image: "images/gallery-2.jpg", Title: "تعمیر صندلی"},
			{Image: "images/gallery-3.jpg", Title: "تعویض پارچه"},
			{Image: "images/gallery-4.jpg", Title: "تعمیر اسکلت"},
			{Image: "images/gallery-5.jpg", Title: "روکش‌کشی"},
			{Image: "images/gallery-6.jpg", Title: "بازسازی کامل"},
		},
		Reviews: []Review{
			{Name: "مریم رضایی", City: "شیراز", Date: "۲ روز پیش", Rating: 5, Service: "تعمیر مبل استیل", Initial: "م", Text: "مبل استیلمون رو بردیم کارگاهشون و واقعاً حرفه‌ای کار کردن. مبل که فکر می‌کردیم باید عوضش کنیم رو مثل روز اولش کردن. قیمت هم کاملاً منصفانه بود."},
			{Name: "علی محمدی", City: "شیراز", Date: "۱ هفته پیش", Rating: 5, Service: "بازسازی کامل", Initial: "ع", Text: "مبل چستر قدیمی مادرم رو کامل بازسازی کردن. از پارچه‌گیری تا اسکلت و فنر. نتیجه واقعاً عالی شد. حتماً به همه توصیه می‌کنم."},
			{Name: "زهرا کریمی", City: "شیراز", Date: "۲ هفته پیش", Rating: 4, Service: "تعویض پارچه", Initial: "ز", Text: "پارچه مبل رو عوض کردن و خیلی تمیز و حرفه‌ای کار کردن. کارگاهشون هم مرتب و منظمه. فقط یه کم دیرتر از قرار قبلی تحویل دادن."},
			{Name: "حسین احمدی", City: "شیراز", Date: "۳ هفته پیش", Rating: 5, Service: "تعمیر صندلی", Initial: "ح", Text: "چند تا صندلی اداری داشتم که فکر نمی‌کردم درست بشن. آقای قهرمانی واقعاً کارشون درجه یکه. صندلی‌ها مثل نو شدن."},
			{Name: "فاطمه نوری", City: "مرودشت", Date: "۱ ماه پیش", Rating: 4, Service: "روکش مبل", Initial: "ف", Text: "کارشون تمیز و باکیفیته. من از مرودشت اومدم شیراز و ارزشش رو داشت. نتیجه کار عالی بود."},
			{Name: "رضا مرادی", City: "شیراز", Date: "۱ ماه پیش", Rating: 5, Service: "تعمیر اسکلت", Initial: "ر", Text: "اسکلت مبل خیلی خراب بود و از چند جا پرسیدم گفتن باید مبل رو عوض کنی. ولی اینجا درستش کردن. عالی بود."},
			{Name: "سارا حسینی", City: "شیراز", Date: "۱ ماه پیش", Rating: 5, Service: "تعمیر مبل استیل", Initial: "س", Text: "خیلی وقت‌شناس و منصف هستن. قبل از شروع کار، قیمت رو دقیق گفتن و هیچ هزینه اضافه‌ای نگرفتن."},
			{Name: "محمد کریمی", City: "کازرون", Date: "۲ ماه پیش", Rating: 4, Service: "بازسازی کامل", Initial: "م", Text: "مبل قدیمی خونه رو کامل بازسازی کردن. کیفیت کار خیلی خوب بود. فقط یه کم بیشتر از چیزی که فکر می‌کردم طول کشید."},
			{Name: "نرگس رحیمی", City: "شیراز", Date: "۲ ماه پیش", Rating: 5, Service: "تعویض پارچه", Initial: "ن", Text: "از انتخاب پارچه تا تحویل نهایی، همه چیز حرفه‌ای بود. به همه دوستام معرفیشون کردم."},
		},
		FAQs: []FAQ{
			{Question: "هزینه تعمیر مبل چقدر است؟", Answer: "هزینه بستگی به نوع مبل و میزان کار دارد. برای اطلاع دقیق، کافیه عکس مبلتون رو واتساپ بفرستید تا در اسرع وقت قیمت دقیق بهتون اعلام بشه."},
			{Question: "آیا مبل رو از منزل ما می‌برید؟", Answer: "خیر، تعمیرات در کارگاه ما انجام می‌شود. شما مبل را به کارگاه ما می‌آورید و بعد از تعمیر تحویل می‌گیرید. برای هماهنگی و آدرس کارگاه، با ما تماس بگیرید."},
			{Question: "چقدر طول می‌کشد تا کار تمام شود؟", Answer: "بسته به نوع کار، معمولاً بین ۱ تا ۵ روز. برای کارهای فوری هم امکان‌پذیره."},
			{Question: "آیا روی کارتون ضمانت می‌دید؟", Answer: "بله، تمام تعمیرات ۶ ماه ضمانت دارند. اگر مشکلی پیش بیاید، رایگان رفع می‌شود."},
			{Question: "چه پارچه‌هایی برای مبل استفاده می‌کنید؟", Answer: "از بهترین پارچه‌های ایرانی و خارجی استفاده می‌کنیم. نمونه‌ها رو می‌تونید به صورت حضوری در کارگاه ببینید."},
		},
		QandAs: []QandA{
			{
				ID: 1, Name: "شاپور زارع", Date: "۱ روز پیش",
				Question: "سلام وقت بخیر من پایه ی زیر مبلام کج شده داره میشکنه باید دلبشه هزینه اش به چه صورت هست؟",
				Replies: []QReply{
					{Author: "تعمیرات مبل قهرمانی", Date: "۱ روز پیش", Text: "سلام. برای تعمیر پایه مبل، بستگی به نوع مبل و شدت شکستگی داره. اگه عکس مبل رو واتساپ بفرستید، دقیق راهنماییتون می‌کنم و قیمت رو اعلام می‌کنم."},
				},
			},
			{
				ID: 2, Name: "مریم مومنی", Date: "۹ روز پیش",
				Question: "سلام مبل تازه خریداری کردم که بعضی جاهایش زدگی و رنگ پریدگی داره امکان ترمیم وجود داره؟",
				Replies: []QReply{
					{Author: "تعمیرات مبل قهرمانی", Date: "۹ روز پیش", Text: "سلام. بله، ترمیم رنگ و زدگی مبل انجام می‌شه. کافیه عکس محل زدگی رو بفرستید تا بررسی کنم و هزینه رو بگم."},
				},
			},
			{
				ID: 3, Name: "عبدالعلی صابری", Date: "۱ ماه پیش",
				Question: "سلام سه تا مبل کلاسیک که تعویض پارچه و فوم و رنگ چوب کردن. چهارتا مبل یک نفره و یک سه نفره راحتی تمام پارچه که تعویض فوم و پارچه و فنر دارن. لطفا بفرمائید هزینه‌اش چقدره و امکان قسطی هست؟",
				Replies: []QReply{
					{Author: "تعمیرات مبل قهرمانی", Date: "۱ ماه پیش", Text: "سلام. برای این حجم کار، حتماً باید حضوری ببینیم. لطفاً عکس‌ها رو واتساپ بفرستید یا تشریف بیارید کارگاه تا هماهنگ کنیم. در مورد قسطی هم می‌تونیم صحبت کنیم."},
				},
			},
			{
				ID: 4, Name: "حسین احمدی", Date: "۵ روز پیش",
				Question: "برای مبل راحتی که فنرهاش خورده شده چقدر هزینه میگیرید؟ پارچه خودم دارم.",
				Replies: []QReply{
					{Author: "تعمیرات مبل قهرمانی", Date: "۵ روز پیش", Text: "سلام. تعویض فنر مبل راحتی بسته به تعداد فنرها و نوع مبل متفاوته. اگه پارچه خودتون باشه، هزینه کمتر می‌شه. عکس بفرستید تا دقیق بگم."},
				},
			},
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

	tmpl, err := template.New("index").Funcs(funcMap).Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join("docs", "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ docs/index.html ساخته شد")
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=5.0">
    <meta name="theme-color" content="#00bfa5">
    <title>{{.Title}}</title>
    <meta name="description" content="تعمیرات مبل در شیراز - تعمیر مبل استیل، چستر، راحتی، تعویض پارچه و روکش مبل در کارگاه تخصصی. با ۷ سال سابقه و ضمانت ۶ ماهه. تماس: {{.Phone}}">
    <meta name="keywords" content="تعمیرات مبل شیراز, تعمیر مبل, کارگاه تعمیر مبل, تعمیر مبل استیل, تعویض پارچه مبل, روکش مبل, تعمیر صندلی شیراز, بازسازی مبل شیراز, تعمیرات مبل قهرمانی">
    <meta name="author" content="تعمیرات مبل قهرمانی">
    <meta name="robots" content="index, follow">
    <link rel="canonical" href="https://moblshiraz.ir/">
    <meta property="og:title" content="{{.Title}}">
    <meta property="og:description" content="{{.About}}">
    <meta property="og:type" content="website">
    <meta property="og:url" content="https://moblshiraz.ir/">
    <meta property="og:image" content="https://moblshiraz.ir/images/hero.jpg">
    <meta property="og:locale" content="fa_IR">
    <link href="https://cdn.jsdelivr.net/gh/rastikerdar/vazirmatn@v33.003/Vazirmatn-font-face.css" rel="stylesheet">
    <script type="application/ld+json">
    {
      "@context": "https://schema.org",
      "@type": "LocalBusiness",
      "name": "تعمیرات مبل قهرمانی",
      "image": "https://moblshiraz.ir/images/hero.jpg",
      "telephone": "{{.PhoneRaw}}",
      "address": {
        "@type": "PostalAddress",
        "addressLocality": "شیراز",
        "addressRegion": "فارس",
        "addressCountry": "IR"
      },
      "areaServed": "شیراز و حومه",
      "description": "{{.About}}",
      "priceRange": "۲۰۰,۰۰۰ - ۵,۰۰۰,۰۰۰ تومان",
      "openingHours": "Sa-Th 09:00-20:00",
      "aggregateRating": {
        "@type": "AggregateRating",
        "ratingValue": "{{.Rating}}",
        "reviewCount": "{{.ReviewsNum}}"
      }
    }
    </script>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; font-family: 'Vazirmatn', Tahoma, sans-serif; -webkit-tap-highlight-color: transparent; }
        :root {
            --primary: #00bfa5;
            --primary-dark: #009688;
            --primary-light: #e0f2f1;
            --bg: #f8f9fa;
            --text: #212529;
            --text-sec: #6c757d;
            --card: #ffffff;
            --border: #e9ecef;
            --yellow: #ffd54f;
        }
        html { scroll-behavior: smooth; }
        body { background: var(--bg); color: var(--text); line-height: 1.8; font-size: 14px; }

        .top-bar { background: white; padding: 12px 16px; border-bottom: 1px solid var(--border); position: sticky; top: 0; z-index: 100; box-shadow: 0 2px 8px rgba(0,0,0,0.04); }
        .top-inner { display: flex; justify-content: space-between; align-items: center; max-width: 1100px; margin: auto; }
        .brand { display: flex; align-items: center; gap: 8px; font-weight: 800; font-size: 16px; color: var(--primary); }
        .brand-icon { width: 32px; height: 32px; border-radius: 10px; overflow: hidden; background: var(--primary); }
        .brand-icon img { width: 100%; height: 100%; object-fit: cover; }
        .top-actions { display: flex; gap: 8px; }
        .icon-btn { width: 38px; height: 38px; border-radius: 10px; display: flex; align-items: center; justify-content: center; text-decoration: none; font-size: 18px; transition: all 0.2s; }
        .icon-btn.phone { background: var(--primary-light); color: var(--primary); }
        .icon-btn.phone:active { background: var(--primary); color: white; }
        .icon-btn.whatsapp { background: #e7f9ef; color: #25D366; }
        .icon-btn.whatsapp:active { background: #25D366; color: white; }

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
        .quick-btn.call:active { background: var(--primary-dark); }
        .quick-btn.wa { background: #25D366; color: white; box-shadow: 0 8px 20px rgba(37,211,102,0.3); }
        .quick-btn.wa:active { background: #1da851; }

        .stats { background: white; padding: 20px 16px; border-top: 1px solid var(--border); }
        .stats-inner { max-width: 1100px; margin: auto; display: flex; justify-content: space-around; text-align: center; }
        .stat-item { flex: 1; }
        .stat-value { font-size: 18px; font-weight: 900; }
        .stat-label { font-size: 11px; color: var(--text-sec); margin-top: 4px; }
        .stat-item + .stat-item { border-right: 1px solid var(--border); }

        .tabs-wrap { position: sticky; top: 63px; z-index: 90; background: white; border-bottom: 1px solid var(--border); overflow-x: auto; -webkit-overflow-scrolling: touch; scrollbar-width: none; }
        .tabs-wrap::-webkit-scrollbar { display: none; }
        .tabs { display: flex; gap: 6px; padding: 10px 16px; max-width: 1100px; margin: auto; white-space: nowrap; }
        .tab { padding: 8px 16px; border-radius: 20px; font-size: 13px; font-weight: 600; background: transparent; color: var(--text-sec); border: 1.5px solid var(--border); cursor: pointer; white-space: nowrap; }
        .tab.active { background: var(--primary); color: white; border-color: var(--primary); }

        .section { padding: 24px 16px; max-width: 1100px; margin: auto; }
        .section-title { font-size: 18px; font-weight: 800; margin-bottom: 6px; display: flex; align-items: center; gap: 8px; }
        .section-title::before { content: ''; width: 4px; height: 20px; background: var(--primary); border-radius: 2px; }
        .section-sub { font-size: 12px; color: var(--text-sec); margin-bottom: 18px; }

        .services-list { display: flex; flex-direction: column; gap: 12px; }
        .service-card { background: white; border: 1px solid var(--border); border-radius: 16px; overflow: hidden; }
        .service-image { width: 100%; aspect-ratio: 16/9; overflow: hidden; background: #f0f0f0; }
        .service-image img { width: 100%; height: 100%; object-fit: cover; display: block; }
        .service-body { padding: 16px; }
        .service-name { font-size: 15px; font-weight: 700; margin-bottom: 4px; }
        .service-desc { font-size: 12px; color: var(--text-sec); line-height: 1.7; margin-bottom: 6px; }
        .service-price { font-size: 13px; font-weight: 700; color: var(--primary); }

        .gallery-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 12px; }
        .gallery-item { border-radius: 14px; overflow: hidden; background: white; border: 1px solid var(--border); transition: transform 0.3s ease, box-shadow 0.3s ease, border-color 0.3s ease; cursor: pointer; }
        .gallery-item:hover { transform: translateY(-5px); box-shadow: 0 12px 30px rgba(0,191,165,0.25); border-color: var(--primary); }
        .gallery-image { width: 100%; aspect-ratio: 1; overflow: hidden; background: #f0f0f0; }
        .gallery-image img { width: 100%; height: 100%; object-fit: cover; display: block; transition: transform 0.4s ease; }
        .gallery-item:hover .gallery-image img { transform: scale(1.1); }
        .gallery-title { font-size: 12px; color: var(--text-sec); font-weight: 600; padding: 10px; text-align: center; transition: color 0.3s; }
        .gallery-item:hover .gallery-title { color: var(--primary); font-weight: 700; }

        .about-box { background: white; border: 1px solid var(--border); border-radius: 16px; overflow: hidden; }
        .about-image { width: 100%; aspect-ratio: 16/9; overflow: hidden; background: #f0f0f0; }
        .about-image img { width: 100%; height: 100%; object-fit: cover; display: block; }
        .about-text { padding: 20px; font-size: 13px; line-height: 2; }

        .review-card { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 16px; margin-bottom: 12px; }
        .review-header { display: flex; align-items: center; gap: 10px; margin-bottom: 10px; }
        .review-avatar { width: 40px; height: 40px; border-radius: 50%; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-weight: 800; font-size: 16px; flex-shrink: 0; }
        .review-name { font-size: 13px; font-weight: 700; }
        .review-meta { font-size: 11px; color: var(--text-sec); }
        .review-text { font-size: 13px; line-height: 1.9; }
        .review-service { display: inline-block; margin-top: 10px; padding: 4px 10px; border-radius: 8px; background: var(--primary-light); color: var(--primary); font-size: 11px; font-weight: 600; }

        .review-stars { font-size: 16px; margin-bottom: 8px; letter-spacing: 2px; line-height: 1; }
        .review-stars .star { display: inline-block; }
        .review-stars .star.filled { color: #ffd54f; }
        .review-stars .star.empty { color: #d1d5db; }

        .review-hidden { display: none; }
        .review-hidden.show { display: block; }
        .show-more-btn { display: block; width: 100%; padding: 14px; background: white; border: 1.5px solid var(--primary); color: var(--primary); border-radius: 14px; font-size: 14px; font-weight: 700; cursor: pointer; font-family: inherit; margin-top: 8px; }
        .show-more-btn:active { background: var(--primary-light); }

        .qa-section { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 16px; }
        .qa-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px; }
        .qa-count { font-size: 14px; font-weight: 800; }
        .qa-count span { color: var(--primary); }
        .qa-ask-btn { background: var(--primary); color: white; padding: 10px 16px; border-radius: 10px; font-size: 13px; font-weight: 700; text-decoration: none; display: flex; align-items: center; gap: 6px; }
        .qa-ask-btn:active { background: var(--primary-dark); }

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
        .qa-show-more:active { background: var(--primary-light); }

        .faq-item { background: white; border: 1px solid var(--border); border-radius: 14px; margin-bottom: 10px; overflow: hidden; }
        .faq-question { padding: 16px; font-size: 13px; font-weight: 700; display: flex; justify-content: space-between; align-items: center; cursor: pointer; gap: 10px; }
        .faq-icon { width: 24px; height: 24px; flex-shrink: 0; border-radius: 50%; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-size: 14px; transition: transform 0.3s; }
        .faq-item.open .faq-icon { transform: rotate(180deg); }
        .faq-answer { padding: 0 16px; max-height: 0; overflow: hidden; font-size: 12.5px; color: var(--text-sec); line-height: 1.9; transition: all 0.3s ease; }
        .faq-item.open .faq-answer { padding: 0 16px 16px; max-height: 500px; }

        .contact-box { background: white; border: 1px solid var(--border); border-radius: 16px; padding: 20px; }
        .contact-item { display: flex; align-items: center; gap: 12px; padding: 14px 0; border-bottom: 1px solid var(--border); text-decoration: none; color: inherit; transition: background 0.2s; }
        .contact-item:last-child { border-bottom: none; }
        .contact-item:active { background: var(--bg); }
        .contact-icon { width: 40px; height: 40px; border-radius: 12px; background: var(--primary-light); color: var(--primary); display: flex; align-items: center; justify-content: center; font-size: 18px; flex-shrink: 0; }
        .contact-icon.wa { background: #e7f9ef; color: #25D366; }
        .contact-icon.map { background: #fef3c7; color: #d97706; }
        .contact-info { flex: 1; }
        .contact-label { font-size: 11px; color: var(--text-sec); }
        .contact-value { font-size: 14px; font-weight: 700; }

        .map-container { margin-top: 16px; border-radius: 16px; overflow: hidden; border: 1px solid var(--border); background: white; }
        .map-container iframe { width: 100%; height: 300px; border: 0; display: block; }

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
        .float-btn.call:active { background: var(--primary-dark); }
        .float-btn.wa { background: #25D366; color: white; }
        .float-btn.wa:active { background: #1da851; }

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
                <div class="brand-icon"><img src="{{.HeroImage}}" alt="{{.Name}}"></div>
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
        <p class="section-sub">نگاهی به کارهای اخیر تعمیرات مبل</p>
        <div class="gallery-grid">
            {{range .Gallery}}
            <div class="gallery-item">
                <div class="gallery-image"><img src="{{.Image}}" alt="{{.Title}} در شیراز" loading="lazy"></div>
                <div class="gallery-title">{{.Title}}</div>
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
                    <button class="qa-toggle" onclick="toggleQA(this)">
                        <span class="arrow">▾</span>
                        <span>مشاهده {{len $qa.Replies}} پاسخ</span>
                    </button>
                    <a href="https://wa.me/{{$.Whatsapp}}" class="qa-reply-btn" target="_blank">ثبت پاسخ ←</a>
                </div>
                <div class="qa-replies">
                    {{range $qa.Replies}}
                    <div class="qa-reply">
                        <div class="qa-reply-header">
                            <div class="qa-reply-avatar">{{slice .Author 0 3}}</div>
                            <div>
                                <div class="qa-reply-name">{{.Author}}</div>
                                <div class="qa-reply-date">{{.Date}}</div>
                            </div>
                        </div>
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
                <div>
                    <div class="review-name">{{$rev.Name}}</div>
                    <div class="review-meta">{{$rev.City}} · {{$rev.Date}}</div>
                </div>
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
            <div class="faq-question" onclick="toggleFaq(this)">
                <span>{{.Question}}</span>
                <span class="faq-icon">▾</span>
            </div>
            <div class="faq-answer">{{.Answer}}</div>
        </div>
        {{end}}
    </section>

    <section class="section" id="contact">
        <h2 class="section-title">تماس با ما</h2>
        <p class="section-sub">برای مشاوره و سفارش در تماس باشید</p>
        <div class="contact-box">
            <a href="tel:{{.PhoneRaw}}" class="contact-item">
                <div class="contact-icon">📞</div>
                <div class="contact-info">
                    <div class="contact-label">تماس تلفنی</div>
                    <div class="contact-value">{{.Phone}}</div>
                </div>
            </a>
            <a href="https://wa.me/{{.Whatsapp}}" class="contact-item" target="_blank">
                <div class="contact-icon wa">💬</div>
                <div class="contact-info">
                    <div class="contact-label">واتساپ</div>
                    <div class="contact-value">ارسال پیام در واتساپ</div>
                </div>
            </a>
            <a href="{{.MapLink}}" class="contact-item" target="_blank">
                <div class="contact-icon map">📍</div>
                <div class="contact-info">
                    <div class="contact-label">آدرس کارگاه</div>
                    <div class="contact-value">{{.Address}}</div>
                </div>
            </a>
            <a href="https://instagram.com/{{.Instagram}}" class="contact-item" target="_blank">
                <div class="contact-icon">📷</div>
                <div class="contact-info">
                    <div class="contact-label">اینستاگرام</div>
                    <div class="contact-value">@{{.Instagram}}</div>
                </div>
            </a>
        </div>
        <div class="trust-badges">
            <div class="trust-badge"><div class="trust-icon">✅</div><div class="trust-text">ضمانت ۶ ماهه</div></div>
            <div class="trust-badge"><div class="trust-icon">🏆</div><div class="trust-text">۷ سال سابقه</div></div>
            <div class="trust-badge"><div class="trust-icon">💰</div><div class="trust-text">قیمت منصفانه</div></div>
        </div>
        <div class="map-container">
            <iframe src="{{.MapEmbed}}" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade" title="موقعیت کارگاه تعمیرات مبل در شیراز"></iframe>
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

    <script>
        function toggleFaq(el) { el.parentElement.classList.toggle('open'); }
        function toggleQA(el) { el.closest('.qa-item').classList.toggle('open'); }
        function showMoreReviews(btn) {
            document.querySelectorAll('.review-hidden').forEach(el => el.classList.add('show'));
            btn.style.display = 'none';
        }
        document.querySelectorAll('.tab').forEach(tab => {
            tab.addEventListener('click', () => {
                document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
                tab.classList.add('active');
                const el = document.getElementById(tab.dataset.tab);
                if (el) el.scrollIntoView({ behavior: 'smooth', block: 'start' });
            });
        });
    </script>
</body>
</html>`
