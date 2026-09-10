import re
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

# Split into tokens by whitespace, then keep only letters (strip punctuation)
words = re.findall(r"[a-zA-Z]+", text.lower())

counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")