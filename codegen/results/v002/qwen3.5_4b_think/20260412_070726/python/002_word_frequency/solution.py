import re
from collections import Counter

with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()

# Convert to lowercase and find all sequences of letters (strips punctuation)
words = re.findall(r'[a-z]+', text.lower())

# Count frequencies
counts = Counter(words)

# Sort by count descending, then word ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

# Print results
for word, count in sorted_words:
    print(f"{word}: {count}")