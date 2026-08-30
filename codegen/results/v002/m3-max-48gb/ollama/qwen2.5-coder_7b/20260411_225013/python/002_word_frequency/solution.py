import re
from collections import Counter

# Read the file and count word frequencies
with open('input/text.txt', 'r') as file:
    text = file.read().lower()
    words = re.findall(r'\b\w+\b', text)
    word_count = Counter(words)

# Sort by frequency (descending) and then alphabetically (ascending for ties)
sorted_word_count = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))

# Output the result
for word, count in sorted_word_count:
    print(f"{word}: {count}")