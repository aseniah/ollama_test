import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()

# Extract words (alphanumeric sequences) and convert to lowercase
words = re.findall(r'[a-zA-Z]+', text.lower())

# Count word frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.keys(), key=lambda w: (-word_counts[w], w))

# Output results
for word in sorted_words:
    print(f"{word}: {word_counts[word]}")