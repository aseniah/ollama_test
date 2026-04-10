import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

text = text.lower()
text = re.sub(r'[^a-z\s]', '', text)
words = text.split()

counts = Counter(words)
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")