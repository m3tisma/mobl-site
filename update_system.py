import json
import re
from datetime import datetime, timedelta
import os

def parse_date(date_str):
    if "امروز" in date_str: return 0
    match = re.search(r'([۰-۹]+)\s*(روز|هفته|ماه)', date_str)
    if not match: return 0
    trans = str.maketrans('۰۱۲۳۴۵۶۷۸۹', '0123456789')
    num = int(match.group(1).translate(trans))
    unit = match.group(2)
    if unit == 'روز': return num
    elif unit == 'هفته': return num * 7
    elif unit == 'ماه': return num * 30
    return 0

def migrate(file):
    if not os.path.exists(file): return
    with open(file, 'r', encoding='utf-8') as f:
        data = json.load(f)
    for item in data:
        d = item.get('Date', 'همین امروز')
        days = parse_date(d)
        created = (datetime.now() - timedelta(days=days)).strftime('%Y-%m-%d')
        item['CreatedAt'] = created
        if 'Date' in item: del item['Date']
        
        if 'Replies' in item:
            for reply in item['Replies']:
                rd = reply.get('Date', 'همین امروز')
                rdays = parse_date(rd)
                rcreated = (datetime.now() - timedelta(days=rdays)).strftime('%Y-%m-%d')
                reply['CreatedAt'] = rcreated
                if 'Date' in reply: del reply['Date']
    with open(file, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)

migrate('data_reviews.json')
migrate('data_qanda.json')

with open('main.go', 'r', encoding='utf-8') as f:
    go_code = f.read()

if '"time"' not in go_code:
    go_code = go_code.replace('"html/template"', '"html/template"\n\t"time"')

go_code = re.sub(r'(type Review struct \{[^}]*?)Date\s+string', r'\1CreatedAt string', go_code, flags=re.DOTALL)
go_code = re.sub(r'(type QandA struct \{[^}]*?)Date\s+string', r'\1CreatedAt string', go_code, flags=re.DOTALL)
go_code = re.sub(r'(type QReply struct \{[^}]*?)Date\s+string', r'\1CreatedAt string', go_code, flags=re.DOTALL)

go_code = go_code.replace('{{$qa.Date}}', '{{relativeDate $qa.CreatedAt}}')
go_code = go_code.replace('{{$rev.Date}}', '{{relativeDate $rev.CreatedAt}}')
go_code = go_code.replace('qa-reply-date">{{.Date}}', 'qa-reply-date">{{relativeDate .CreatedAt}}')

if '"relativeDate":' not in go_code:
    func_code = '''
        "relativeDate": func(dateStr string) string {
            t, err := time.Parse("2006-01-02", dateStr)
            if err != nil {
                return dateStr
            }
            now := time.Now()
            now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
            t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
            days := int(now.Sub(t).Hours() / 24)
            if days <= 0 { return "همین امروز" }
            if days == 1 { return "۱ روز پیش" }
            if days == 2 { return "۲ روز پیش" }
            if days == 3 { return "۳ روز پیش" }
            if days == 4 { return "۴ روز پیش" }
            if days == 5 { return "۵ روز پیش" }
            if days == 6 { return "۶ روز پیش" }
            if days >= 7 && days < 14 { return "۱ هفته پیش" }
            if days >= 14 && days < 21 { return "۲ هفته پیش" }
            if days >= 21 && days < 30 { return "۳ هفته پیش" }
            months := days / 30
            if months == 1 { return "۱ ماه پیش" }
            if months == 2 { return "۲ ماه پیش" }
            if months == 3 { return "۳ ماه پیش" }
            if months == 4 { return "۴ ماه پیش" }
            if months == 5 { return "۵ ماه پیش" }
            if months == 6 { return "۶ ماه پیش" }
            return fmt.Sprintf("%d ماه پیش", months)
        },
    '''
    go_code = go_code.replace('funcMap := template.FuncMap{', 'funcMap := template.FuncMap{\n' + func_code)

with open('main.go', 'w', encoding='utf-8') as f:
    f.write(go_code)
print("✅ کدها با موفقیت آپدیت شدند!")
