import re
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Extract only sequences of letters (strips all punctuation)
words = re.findall(r'[a-zA-Z]+', text)

# Convert to lowercase
words = [w.lower() for w in words]

# Count frequencies
counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")