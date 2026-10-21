import re

with open('main.go', 'r', encoding='utf-8') as f:
    code = f.read()

# ۱. اصلاح پکیج‌های تکراری
code = re.sub(r'import \(.*?\)', '''import (
    "html/template"
    "log"
    "os"
    "path/filepath"
    "fmt"
    "time"
    "encoding/json"
)''', code, flags=re.DOTALL)

# ۲. حذف لیست هاردکد شده Reviews
code = re.sub(r'Reviews:\s*\[\]Review\{[\s\S]*?\n\t\t\},', '// Reviews loaded from JSON', code, flags=re.DOTALL)

# ۳. حذف لیست هاردکد شده QandAs
code = re.sub(r'QandAs:\s*\[\]QandA\{[\s\S]*?\n\t\t\},', '// QandAs loaded from JSON', code, flags=re.DOTALL)

# ۴. حذف کدهای تکراری خواندن JSON (اگه چند بار اضافه شدن)
code = re.sub(r'\n\s*// خواندن نظرات از فایل JSON[\s\S]*?(?=blogPosts :=)', '', code)
code = re.sub(r'\n\s*// خواندن پرسش و پاسخ از فایل JSON[\s\S]*?(?=blogPosts :=)', '', code)

# ۵. اضافه کردن یک بار کد خواندن JSON به صورت تمیز
json_code = '''
	// خواندن نظرات از فایل JSON
	reviewData, err := os.ReadFile("data_reviews.json")
	if err == nil {
		json.Unmarshal(reviewData, &data.Reviews)
	}

	// خواندن پرسش و پاسخ از فایل JSON
	qandaData, err := os.ReadFile("data_qanda.json")
	if err == nil {
		json.Unmarshal(qandaData, &data.QandAs)
	}

'''
code = code.replace('blogPosts := []BlogPost{', json_code + '\tblogPosts := []BlogPost{')

with open('main.go', 'w', encoding='utf-8') as f:
    f.write(code)

print("✅ فایل main.go با موفقیت تعمیر شد!")
