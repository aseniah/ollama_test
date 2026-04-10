import re
import string
from collections import Counter

# Read the file
try:
    with open('input/text.txt', 'r', encoding='utf-8') as f:
        text = f.read()
except FileNotFoundError:
    text = ''

# Convert to lowercase
text = text.lower()

# Extract words (keep only letters)
words = re.findall(r'[a-z]+', text)

# Count word frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output in the required format
for word, count in sorted_words:
    print(f"{word}: {count}")
```