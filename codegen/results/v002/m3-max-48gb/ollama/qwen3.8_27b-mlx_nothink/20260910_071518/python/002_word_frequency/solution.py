import re
from collections import defaultdict

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Extract words: keep only letters, split by non-letter sequences
words = re.findall(r'[a-z]+', text)

# Count frequency
freq = defaultdict(int)
for word in words:
    freq[word] += 1

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")