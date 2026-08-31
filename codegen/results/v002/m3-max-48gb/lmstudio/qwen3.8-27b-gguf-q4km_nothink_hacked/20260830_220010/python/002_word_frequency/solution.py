import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Keep only letters (strip all punctuation)
# We want to keep only alphabetic characters, so we'll extract words consisting of letters
# Split by non-letter characters
words = re.findall(r'[a-z]+', text)

# Count frequencies
counter = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")