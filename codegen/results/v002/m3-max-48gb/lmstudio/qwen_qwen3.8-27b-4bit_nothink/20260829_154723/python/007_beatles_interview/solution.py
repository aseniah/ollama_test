import csv, json, datetime

as_of = datetime.date(2025, 7, 1)

with open('input/input.csv') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

def parse_date(s):
    if not s or s.strip().lower() == 'null':
        return None
    return datetime.datetime.strptime(s.strip(), '%m/%d/%Y').date()

def age_at(as_of, birth):
    return as_of.year - birth.year - ((as_of.month, as_of.day) < (birth.month, birth.day))

def split_name(full_name):
    if not full_name or full_name.strip().lower() == 'null':
        return None, None
    parts = full_name.strip().split()
    if len(parts) == 1:
        return parts[0], ''
    return parts[0], ' '.join(parts[1:])

def rel(full_name, relationship):
    if not full_name or full_name.strip().lower() == 'null':
        return None
    first, last = split_name(full_name)
    return {'FirstName': first, 'LastName': last, 'Relationship': relationship}

out = []
for r in rows:
    full_name = r['Name'].strip()
    parts = full_name.split()
    first_name = parts[0]
    last_name = ' '.join(parts[1:]) if len(parts) > 1 else ''

    birth = parse_date(r['Birthday'])
    age = age_at(as_of, birth)
    birthday_str = birth.strftime('%Y-%m-%d')

    relatives = []
    for key, rel_type in [('Father','Father'),('Mother','Mother'),('Brother','Brother'),('Sister','Sister')]:
        val = r.get(key, '')
        rec = rel(val, rel_type)
        if rec:
            relatives.append(rec)

    out.append({
        'FirstName': first_name,
        'LastName': last_name,
        'Birthday': birthday_str,
        'Age': age,
        'Relatives': relatives
    })

print(json.dumps(out, indent=2))