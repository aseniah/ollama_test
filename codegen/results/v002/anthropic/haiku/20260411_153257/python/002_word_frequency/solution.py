import sys
import re
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

words = re.findall(r'[a-zA-Z]+', text.lower())
counts = Counter(words)

for word, count in sorted(counts.items(), key=lambda x: (-x[1], x[0])):
    print(f"{word}: {count}")
