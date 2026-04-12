import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Remove punctuation - keep only letters and spaces
text = re.sub(r'[^a-z\s]', ' ', text)

# Split into words (multiple spaces will be handled by split() with no arguments)
words = text.split()

# Count word frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output in required format
for word, count in sorted_words:
    print(f"{word}: {count}")