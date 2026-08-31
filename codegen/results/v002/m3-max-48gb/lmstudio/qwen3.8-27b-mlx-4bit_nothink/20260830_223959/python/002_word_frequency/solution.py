import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase and strip all punctuation (keep only letters)
text = text.lower()
# Keep only letters and whitespace
text = re.sub(r'[^a-z\s]', '', text)

# Split into words
words = text.split()

# Count frequency
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")