import re
from collections import Counter

with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()

text = text.lower()
words = re.findall(r'[a-z]+', text)
counts = Counter(words)
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")