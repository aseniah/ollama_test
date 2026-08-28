import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Find all words (sequences of letters only)
words = re.findall(r'[a-z]+', text)

# Count frequency
counter = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")