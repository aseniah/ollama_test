import re
from collections import Counter

# Read input file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Extract words: lowercase + alphabetic characters only
words = [word for word in text.lower()]
words = [re.sub(r'[^a-z]', '', char) for char in words]

# Group into actual words (by matching sequences between non-letters and spaces/newlines etc.)
import re
words = re.findall(r'[a-z]+', text.lower())

# Count word frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Print results in required format
for word, count in sorted_words:
    print(f'{word}: {count}')