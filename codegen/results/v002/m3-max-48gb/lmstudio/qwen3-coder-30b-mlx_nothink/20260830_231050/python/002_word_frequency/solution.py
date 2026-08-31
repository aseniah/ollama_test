import re
from collections import Counter

# Read the file
with open('input/text.txt', 'r') as file:
    content = file.read()

# Convert to lowercase and extract words (only letters)
words = re.findall(r'[a-zA-Z]+', content.lower())

# Count frequency of each word
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")