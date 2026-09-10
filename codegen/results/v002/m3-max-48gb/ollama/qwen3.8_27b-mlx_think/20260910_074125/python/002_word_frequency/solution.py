import re
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

# Split into words, convert to lowercase, keep only letters
words = []
for token in text.split():
    # Keep only letters
    word = re.sub(r'[^a-z]', '', token.lower())
    if word:
        words.append(word)

counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")