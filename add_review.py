import json
import random
from datetime import datetime

# ================= دیتای اسم‌ها =================
fa_first = ["علی", "سارا", "محمد", "زهرا", "رضا", "مریم", "حسین", "نرگس", "امیر", "فاطمه", "مهدی", "الهام", "پویا", "نگار", "کیان", "آیدا", "شایان", "پریسا", "بابک", "مینا", "آرش", "لیلا", "سینا", "رها", "کامران", "شقایق", "فرهاد", "سمیرا"]
en_first = ["Ali", "Sara", "Reza", "Maryam", "Hossein", "Amir", "Fatemeh", "Mohammad", "Zahra", "Narges", "Mehdi", "Elham", "Pouya", "Negar", "Kian", "Aida", "Shayan", "Parisa", "BabaK", "Mina", "Arash", "Leila", "Sina", "Raha", "Kamran", "Shaghayegh", "Farhad", "Samira"]
fa_last = ["محمدی", "رضایی", "حسینی", "کریمی", "احمدی", "نوری", "صادقی", "موسوی", "جعفری", "قاسمی", "زمانی", "عبدی", "بهرامی", "شریفی", "اکبری", "امینی", "رحیمی", "علوی", "فرهادی", "کاظمی", "مرادی", "هاشمی", "یوسفی", "طاهری"]
en_last = ["Mohammadi", "Rezaei", "Hosseini", "Karimi", "Ahmadi", "Nouri", "Sadeghi", "Mousavi", "Jafari", "Ghasemi", "Zamani", "Abdi", "Bahrami", "Sharifi", "Akbari", "Amini", "Rahimi", "Alavi", "Farhadi", "Kazemi", "Moradi", "Hashemi", "Yousefi", "Taheri"]

cities = ["شیراز", "مرودشت", "کازرون", "آباده", "صدرا", "زرقان", "اقلید", "فسا", "داراب"]
services = ["تعمیر مبل چستر", "تعویض پارچه", "تعمیر مبل راحتی", "بازسازی کامل", "تعمیر صندلی", "روکش مبل", "تعمیر سرویس خواب"]

# متن‌های طولانی‌تر (۱ تا ۳ خط)
review_texts = [
    "کارشون واقعاً عالی بود. مبل قدیمی که فکر می‌کردیم باید بندازیمش دور رو بردیم کارگاهشون و مثل روز اولش تحویل گرفتیم. کیفیت پارچه و دوختش فوق‌العاده بود. حتماً به دوستانم معرفی می‌کنم.",
    "سر وقت تحویل دادن و قیمتشون هم کاملاً منصفانه و منصفانه‌تر از بقیه جاهایی بود که پرسیدم. از مشاوره قبل از شروع کار تا تحویل نهایی، همه چیز حرفه‌ای بود. ممنون از تیم خوبتون.",
    "کیفیت پارچه‌ای که استفاده کردن خیلی خوبه و رنگش هم دقیقاً همونی بود که انتخاب کرده بودیم. از نتیجه کار کاملاً راضی هستیم. فقط یه کم بیشتر از چیزی که فکر می‌کردیم طول کشید که اونم به خاطر دقت بالاشون بود.",
    "تعمیرات رو خیلی تمیز و حرفه‌ای انجام دادن. مبل راحتی‌مون فنرهاش صدا می‌داد و اسفنجش فرورفته بود، ولی الان مثل نو شده. دستشون درد نکنه، واقعاً کارشون درسته.",
    "قبل از کار دقیق راهنمایی کردن و هیچ هزینه اضافه‌ای نگرفتن. حتی بعد از تحویل هم پیگیر شدن که همه چیز خوبه یا نه. اینجور برخوردی خیلی کم پیش میاد، ممنون.",
    "مبل مادرم رو که خیلی برامون ارزش داشت رو بردیم و خیلی قشنگ و سریع بازسازی کردن. اسکلت و پارچه‌ش رو کامل عوض کردن و نتیجه واقعاً چشم‌نواز شد. حتماً دوباره همکاری می‌کنیم.",
    "من از شهرستان اومدم شیراز و خیلی نگران بودم که مبل رو چطور تحویل بدم. ولی راهنمایی‌های تلفنی و واتساپیشون عالی بود. وقتی مبل رو تحویل گرفتم، از کیفیتش شوکه شدم. پیشنهاد می‌کنم حتماً بهشون اعتماد کنید."
]

qanda_questions = [
    "سلام، هزینه تعمیر مبل استیل ۹ نفره چقدر میشه؟",
    "برای مبل چستر که پارچه‌ش پاره شده، پارچه خودم رو دارم. هزینه دستمزد چقدره؟",
    "مبل راحتی ما فنرهاش صدا میده، آیا در محل هم تعمیر میکنید؟",
    "تعمیر اسکلت مبل چقدر زمان میبره؟",
    "آیا برای تعمیر صندلی اداری هم قطعات یدکی دارید؟",
    "سلام، مبل چرمی ما ترک خورده، قابل ترمیمه یا باید پارچه‌ش کنیم؟",
    "هزینه تعویض اسفنج مبل راحتی چقدره؟",
    "برای تعمیر مبل، باید حتماً مبل رو بیاریم کارگاه یا میشه در محل انجام داد؟"
]
qanda_answers = [
    "سلام. برای اعلام دقیق قیمت، لطفاً عکس مبل رو به واتساپ ما بفرستید تا کارشناسان ما بررسی کنند و بهترین قیمت رو اعلام کنند.",
    "سلام. اگه پارچه رو خودتون تهیه کنید، هزینه بسیار کمتر میشه. لطفاً عکس مبل و متراژ پارچه‌تون رو واتساپ بفرستید.",
    "سلام. بله، بسیاری از تعمیرات جزئی در محل انجام میشه. برای هماهنگی و مشاوره رایگان با ما تماس بگیرید.",
    "سلام. تعمیر اسکلت معمولاً بین ۳ تا ۵ روز کاری زمان میبره. بسته به نوع مبل متغیره.",
    "سلام. بله، قطعات یدکی انواع صندلی‌های اداری موجوده. عکس صندلی رو بفرستید تا راهنمایی کنیم.",
    "سلام. ترک خوردگی چرم در بسیاری از موارد قابل ترمیم و رنگ‌آمیزی هست. عکس مبل رو بفرستید تا دقیق بگم.",
    "سلام. هزینه تعویض اسفنج بسته به نوع فوم و ابعاد مبل متفاوته. لطفاً عکس مبل رو واتساپ کنید.",
    "سلام. برای کارهای اساسی بهتره مبل به کارگاه آورده بشه، اما برای برخی تعمیرات جزئی امکان پذیرش در محل وجود داره. با ما تماس بگیرید."
]

# ================= منطق تولید اسم =================
def get_all_used_names():
    used = set()
    for file in ['data_reviews.json', 'data_qanda.json']:
        try:
            with open(file, 'r', encoding='utf-8') as f:
                data = json.load(f)
                for item in data:
                    if 'Name' in item:
                        used.add(item['Name'].strip())
        except:
            pass
    return used

def generate_unique_name(used_names):
    # تلاش برای ساخت اسم یکتا (حداکثر ۵۰ بار تلاش میکنه تا ارور نده)
    for _ in range(50):
        first_fa = random.choice(fa_first)
        first_en = random.choice(en_first)
        last_fa = random.choice(fa_last)
        last_en = random.choice(en_last)
        
        # ترکیب‌های مختلف فارسی و انگلیسی
        choice = random.randint(1, 6)
        if choice == 1: name = first_fa
        elif choice == 2: name = f"{first_fa} {last_fa}"
        elif choice == 3: name = f"{first_en} {last_en}"
        elif choice == 4: name = f"{first_fa} {last_en}"
        elif choice == 5: name = f"{first_en} {last_fa}"
        else: name = f"{first_fa} {last_fa[0]}."
        
        if name not in used_names:
            return name
    return f"کاربر {random.randint(1000, 9999)}" # فالبک

# ================= پردازش =================
now = datetime.now()
today_str = now.strftime('%Y-%m-%d')
used_names = get_all_used_names()

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
    if (now - last_date).days < random.randint(2, 4):
        should_add = False

if should_add:
    name = generate_unique_name(used_names)
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
    q_name = generate_unique_name(used_names)
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

print("✅ پردازش با موفقیت انجام شد! (بررسی یکتایی اسم و نظرات طولانی اعمال شد)")
