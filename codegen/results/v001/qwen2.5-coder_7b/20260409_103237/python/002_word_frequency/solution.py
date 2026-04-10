import re
from collections import Counter

# Open and read file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase, remove punctuation, split into words
words = re.findall(r'\b\w+\b', text.lower())

# Count word frequencies using a counter object
word_counts = Counter(words)

# Sort by count descending and alphabetically ascending for ties
sorted_counts = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Print each word with its count in the desired format
for word, count in sorted_counts:
    print(f'{word}: {count}')