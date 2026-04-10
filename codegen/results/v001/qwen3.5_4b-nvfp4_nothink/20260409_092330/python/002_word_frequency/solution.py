import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Remove punctuation and keep only letters, then extract words
words = re.findall(r'[a-z]+', text.lower())

# Count frequencies
counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")