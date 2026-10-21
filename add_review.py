import json
import random
from datetime import datetime

first_names = ["علی", "سارا", "محمد", "زهرا", "رضا", "مریم", "حسین", "نرگس", "امیر", "فاطمه", "مهدی", "الهام", "پویا", "نگار", "کیان", "آیدا"]
last_names = ["محمدی", "رضایی", "حسینی", "کریمی", "احمدی", "نوری", "صادقی", "موسوی", "جعفری", "قاسمی", "زمانی", "عبدی", "بهرامی", "شریفی"]
weird_names = ["مبل‌دوست", "دکوراتور", "شیرازی", "کاربر مهمان", "خریدار قبلی"]
cities = ["شیراز", "مرودشت", "کازرون", "آباده", "صدرا", "زرقان"]
services = ["تعمیر مبل چستر", "تعویض پارچه", "تعمیر مبل راحتی", "بازسازی کامل", "تعمیر صندلی", "روکش مبل"]
review_texts = [
    "کارشون عالی بود، مبل قدیمی مثل روز اولش شد. حتماً به دوستانم معرفی می‌کنم.",
    "سر وقت تحویل دادن و قیمتشون هم منصفانه بود. ممنون از تیم خوبتون.",
    "کیفیت پارچه‌ای که استفاده کردن خیلی خوبه. از نتیجه کار کاملاً راضی هستم.",
    "تعمیرات رو خیلی تمیز و حرفه‌ای انجام دادن. دستشون درد نکنه.",
    "قبل از کار دقیق راهنمایی کردن و هیچ هزینه اضافه‌ای نگرفتن. عالی بود.",
    "مبل مادرم رو بردیم، خیلی قشنگ و سریع تحویل دادن. واقعاً حرفه‌ای هستن."
]
qanda_questions = [
    "سلام، هزینه تعمیر مبل استیل ۹ نفره چقدر میشه؟",
    "برای مبل چستر که پارچه‌ش پاره شده، پارچه خودم رو دارم. هزینه دستمزد چقدره؟",
    "مبل راحتی ما فنرهاش صدا میده، آیا در محل هم تعمیر میکنید؟",
    "تعمیر اسکلت مبل چقدر زمان میبره؟",
    "آیا برای تعمیر صندلی اداری هم قطعات یدکی دارید؟",
    "سلام، مبل چرمی ما ترک خورده، قابل ترمیمه یا باید پارچه‌ش کنیم؟",
    "هزینه تعویض اسفنج مبل راحتی چقدره؟"
]
qanda_answers = [
    "سلام. برای اعلام دقیق قیمت، لطفاً عکس مبل رو به واتساپ ما بفرستید تا کارشناسان ما بررسی کنند و بهترین قیمت رو اعلام کنند.",
    "سلام. اگه پارچه رو خودتون تهیه کنید، هزینه بسیار کمتر میشه. لطفاً عکس مبل و متراژ پارچه‌تون رو واتساپ بفرستید.",
    "سلام. بله، بسیاری از تعمیرات جزئی در محل انجام میشه. برای هماهنگی و مشاوره رایگان با ما تماس بگیرید.",
    "سلام. تعمیر اسکلت معمولاً بین ۳ تا ۵ روز کاری زمان میبره. بسته به نوع مبل متغیره.",
    "سلام. بله، قطعات یدکی انواع صندلی‌های اداری موجوده. عکس صندلی رو بفرستید تا راهنمایی کنیم.",
    "سلام. ترک خوردگی چرم در بسیاری از موارد قابل ترمیم و رنگ‌آمیزی هست. عکس مبل رو بفرستید تا دقیق بگم.",
    "سلام. هزینه تعویض اسفنج بسته به نوع فوم و ابعاد مبل متفاوته. لطفاً عکس مبل رو واتساپ کنید."
]

def generate_name():
    choice = random.randint(1, 4)
    first = random.choice(first_names)
    last = random.choice(last_names)
    if choice == 1: return first
    elif choice == 2: return f"{first} {last}"
    elif choice == 3: return f"{first} {last[0]}."
    else: return f"{first} {random.choice(weird_names)}"

now = datetime.now()
today_str = now.strftime('%Y-%m-%d')

# ---- Reviews ----
try:
    with open('data_reviews.json', 'r', encoding='utf-8') as f:
        reviews = json.load(f)
except:
    reviews = []

should_add = True
if reviews:
    last_date_str = reviews[0].get('CreatedAt', '2000-01-01')
    last_date = datetime.strptime(last_date_str, '%Y-%m-%d')
    # انتخاب تصادفی بین ۲ تا ۴ روز
    if (now - last_date).days < random.randint(2, 4):
        should_add = False

if should_add:
    name = generate_name()
    new_review = {
        "Name": name,
        "City": random.choice(cities),
        "CreatedAt": today_str,
        "Rating": 5,
        "Service": random.choice(services),
        "Initial": name[0] if name else "ک",
        "Text": random.choice(review_texts)
    }
    reviews.insert(0, new_review)

with open('data_reviews.json', 'w', encoding='utf-8') as f:
    json.dump(reviews, f, ensure_ascii=False, indent=2)

# ---- QandA ----
try:
    with open('data_qanda.json', 'r', encoding='utf-8') as f:
        qandas = json.load(f)
except:
    qandas = []

should_add_qa = True
if qandas:
    last_date_str = qandas[0].get('CreatedAt', '2000-01-01')
    last_date = datetime.strptime(last_date_str, '%Y-%m-%d')
    if (now - last_date).days < random.randint(2, 4):
        should_add_qa = False

if should_add_qa:
    q_name = generate_name()
    new_qanda = {
        "ID": len(qandas) + 1,
        "Name": q_name,
        "CreatedAt": today_str,
        "Question": random.choice(qanda_questions),
        "Replies": [
            {
                "Author": "مبل شیراز",
                "CreatedAt": today_str,
                "Text": random.choice(qanda_answers)
            }
        ]
    }
    qandas.insert(0, new_qanda)

with open('data_qanda.json', 'w', encoding='utf-8') as f:
    json.dump(qandas, f, ensure_ascii=False, indent=2)

print("✅ پردازش با موفقیت انجام شد!")
