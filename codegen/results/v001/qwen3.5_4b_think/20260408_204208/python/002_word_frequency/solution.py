from collections import Counter
import re

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read().lower()

# Extract words (only alphabetic characters)
words = re.findall(r'[a-z]+', text)

# Count frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")