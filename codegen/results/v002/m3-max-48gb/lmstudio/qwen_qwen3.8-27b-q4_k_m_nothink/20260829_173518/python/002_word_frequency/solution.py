import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Extract only alphabetic words (keep only letters)
# Split by non-alphabetic characters
words = re.findall(r'[a-z]+', text)

# Count frequency
counter = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")