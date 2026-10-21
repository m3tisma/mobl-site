import json
import random
import datetime
import os
import sys

# لیست اسم‌ها و متن‌ها
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

def generate_name():
    choice = random.randint(1, 4)
    first = random.choice(first_names)
    last = random.choice(last_names)
    if choice == 1: return first 
    elif choice == 2: return f"{first} {last}" 
    elif choice == 3: return f"{first} {last[0]}." 
    else: return f"{first} {random.choice(weird_names)}" 

# ==========================================
# منطق زمان‌بندی تصادفی (بین 3 تا 7 روز)
# ==========================================
state_file = 'last_run.json'
today = datetime.date.today()

if os.path.exists(state_file):
    with open(state_file, 'r', encoding='utf-8') as f:
        state = json.load(f)
    last_date = datetime.date.fromisoformat(state['last_date'])
    target_days = state['target_days']
    
    days_passed = (today - last_date).days
    if days_passed < target_days:
        print(f"⏳ هنوز زوده! {days_passed} روز گذشته، باید {target_days} روز بگذره.")
        sys.exit(0) # خروج بدون خطا (ربات نظر نمیده)

# ==========================================
# اگر زمانش رسیده باشه، نظر جدید می‌سازیم
# ==========================================
with open('data_reviews.json', 'r', encoding='utf-8') as f:
    reviews = json.load(f)

name = generate_name()
new_review = {
    "Name": name,
    "City": random.choice(cities),
    "Date": "همین امروز",
    "Rating": 5,
    "Service": random.choice(services),
    "Initial": name[0] if name else "ک",
    "Text": random.choice(texts)
}

reviews.insert(0, new_review)

with open('data_reviews.json', 'w', encoding='utf-8') as f:
    json.dump(reviews, f, ensure_ascii=False, indent=2)

# انتخاب تصادفی زمان اجرای بعدی (بین 3 تا 7 روز)
next_days = random.randint(3, 7)
with open(state_file, 'w', encoding='utf-8') as f:
    json.dump({'last_date': today.isoformat(), 'target_days': next_days}, f, ensure_ascii=False, indent=2)

print(f"✅ نظر جدید برای {name} اضافه شد! اجرای بعدی بعد از {next_days} روز.")
