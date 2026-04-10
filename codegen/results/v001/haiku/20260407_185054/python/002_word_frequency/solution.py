import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read().lower()

words = re.findall(r'[a-z]+', text)

word_counts = Counter(words)

sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")
