import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Extract only alphabetic words (strip punctuation)
words = re.findall(r'[a-z]+', text)

# Count frequency
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f'{word}: {count}')