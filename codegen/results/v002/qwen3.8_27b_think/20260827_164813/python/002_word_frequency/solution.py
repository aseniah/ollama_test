import re

with open('input/text.txt', 'r') as f:
    text = f.read()

words = text.split()
freq = {}

for word in words:
    # Keep only letters, convert to lowercase
    cleaned = ''.join(c for c in word if c.isalpha()).lower()
    if cleaned:
        freq[cleaned] = freq.get(cleaned, 0) + 1

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")