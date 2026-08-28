import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase and extract words (keeping only letters)
words = re.findall(r'[a-zA-Z]+', text.lower())

# Count frequency
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")