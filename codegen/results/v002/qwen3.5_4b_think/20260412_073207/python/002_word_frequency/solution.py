import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase and extract only letters
text = text.lower()
# Match sequences of only letters (word)
words = re.findall(r'[a-z]+', text)

# Count word frequencies
counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

# Output the results
for word, count in sorted_words:
    print(f"{word}: {count}")