import sys
import string
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

words = text.lower().split()
cleaned = [w.strip(string.punctuation) for w in words]
cleaned = [w for w in cleaned if w.isalpha()]

counts = Counter(cleaned)
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")
