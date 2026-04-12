import re
from collections import Counter

# Read the file and convert it to lowercase
with open('input/text.txt', 'r') as file:
    text = file.read().lower()

# Remove punctuation and split into words
words = re.findall(r'\b\w+\b', text)

# Count the frequency of each word
word_count = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_word_count = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))

# Output one line per unique word in the format: `word: count`
for word, count in sorted_word_count:
    print(f'{word}: {count}')