import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Extract words: keep only letters, so split by non-letter characters
# First, remove all non-letter characters, but we need to handle word boundaries
# A better approach: split by anything that is not a letter, then filter out empty strings
words = re.split(r'[^a-z]+', text)
words = [word for word in words if word]  # Remove empty strings

# Count frequencies
counter = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")