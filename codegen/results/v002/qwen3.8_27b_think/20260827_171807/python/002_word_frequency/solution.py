import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Split into words
words = text.split()

# Strip punctuation - keep only letters
cleaned = []
for word in words:
    cleaned_word = re.sub(r'[^a-z]', '', word)
    if cleaned_word:
        cleaned.append(cleaned_word)

# Count frequencies
counts = Counter(cleaned)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")