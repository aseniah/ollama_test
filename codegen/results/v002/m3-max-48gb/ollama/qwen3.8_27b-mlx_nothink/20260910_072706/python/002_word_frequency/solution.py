import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Extract only letters to form words (keep only letters)
# This removes punctuation and non-letter characters
words = re.findall(r'[a-z]+', text)

# Count frequency of each word
counter = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

# Output each word and its count
for word, count in sorted_words:
    print(f"{word}: {count}")