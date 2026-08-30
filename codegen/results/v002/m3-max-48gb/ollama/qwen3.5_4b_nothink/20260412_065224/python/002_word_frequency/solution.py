import re
from collections import Counter

# Read the input file
with open('input/text.txt', 'r') as f:
    text = f.read().lower()

# Extract words (only letters)
words = re.findall(r'[a-z]+', text)

# Count frequencies
word_counts = Counter(words)

# Sort: primary by count descending, secondary by word alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")