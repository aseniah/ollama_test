from collections import Counter
import re

# Read the input file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase and extract words (sequences of letters only)
text_lower = text.lower()
words = re.findall(r'[a-z]+', text_lower)

# Count word frequencies
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Print results in the required format
for word, count in sorted_words:
    print(f"{word}: {count}")