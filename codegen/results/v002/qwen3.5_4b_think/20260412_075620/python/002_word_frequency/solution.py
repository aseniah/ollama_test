#!/usr/bin/env python3
import string
import re

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Tokenize the text
# Extract words (only alphabetic characters)
words = re.findall(r'[a-zA-Z]+', text.lower())

# Count word frequencies
word_count = {}
for word in words:
    word_count[word] = word_count.get(word, 0) + 1

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")