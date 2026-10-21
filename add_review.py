import json
import random
from datetime import datetime

# لیست اسم‌ها برای تنوع بالا
first_names = ["علی", "سارا", "محمد", "زهرا", "رضا", "مریم", "حسین", "نرگس", "امیر", "فاطمه", "مهدی", "الهام", "پویا", "نگار", "کیان", "آیدا"]
last_names = ["محمدی", "رضایی", "حسینی", "کریمی", "احمدی", "نوری", "صادقی", "موسوی", "جعفری", "قاسمی", "زمانی", "عبدی", "بهرامی", "شریفی"]
weird_names = ["مبل‌دوست", "دکوراتور", "شیرازی", "کاربر مهمان", "خریدار قبلی"]
cities = ["شیراز", "مرودشت", "کازرون", "آباده", "صدرا", "زرقان"]
services = ["تعمیر مبل چستر", "تعویض پارچه", "تعمیر مبل راحتی", "بازسازی کامل", "تعمیر صندلی", "روکش مبل"]
texts = [
    "کارشون عالی بود، مبل قدیمی مثل روز اولش شد. حتماً به دوستانم معرفی می‌کنم.",
    "سر وقت تحویل دادن و قیمتشون هم منصفانه بود. ممنون از تیم خوبتون.",
    "کیفیت پارچه‌ای که استفاده کردن خیلی خوبه. از نتیجه کار کاملاً راضی هستم.",
    "تعمیرات رو خیلی تمیز و حرفه‌ای انجام دادن. دستشون درد نکنه.",
    "قبل از کار دقیق راهنمایی کردن و هیچ هزینه اضافه‌ای نگرفتن. عالی بود.",
    "مبل مادرم رو بردیم، خیلی قشنگ و سریع تحویل دادن. واقعاً حرفه‌ای هستن."
]

# ساخت اسم با فرمت‌های مختلف
def generate_name():
    choice = random.randint(1, 4)
    first = random.choice(first_names)
    last = random.choice(last_names)
    if choice == 1:
        return first # فقط اسم
    elif choice == 2:
        return f"{first} {last}" # اسم و فامیل کامل
    elif choice == 3:
        return f"{first} {last[0]}." # اسم و حرف اول فامیل
    else:
        return f"{first} {random.choice(weird_names)}" # اسم با فامیل عجیب

# خوندن فایل فعلی
with open('data_reviews.json', 'r', encoding='utf-8') as f:
    reviews = json.load(f)

# ساخت نظر جدید
name = generate_name()
new_review = {
    "Name": name,
    "City": random.choice(cities),
    "Date": "همین امروز", # چون ربات الان داره اضافه‌ش می‌کنه
    "Rating": 5,
    "Service": random.choice(services),
    "Initial": name[0] if name else "ک",
    "Text": random.choice(texts)
}

# اضافه کردن به اول لیست
reviews.insert(0, new_review)

# ذخیره
with open('data_reviews.json', 'w', encoding='utf-8') as f:
    json.dump(reviews, f, ensure_ascii=False, indent=2)

print(f"✅ نظر جدید برای {name} اضافه شد!")
