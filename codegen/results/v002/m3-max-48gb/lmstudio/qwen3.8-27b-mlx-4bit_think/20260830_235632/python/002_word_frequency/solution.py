import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

words = []
for word in text.split():
    cleaned = ''.join(c.lower() for c in word if c.isalpha())
    if cleaned:
        words.append(cleaned)

counts = Counter(words)

sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")