import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Keep only letters and whitespace, then split into words
# Replace all non-letter, non-whitespace characters with space
cleaned = re.sub(r'[^a-z\s]', ' ', text)

words = cleaned.split()

# Count frequencies
counter = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")