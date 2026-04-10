python
import string
import re
from collections import Counter

with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()

# Remove punctuation, keep only letters and spaces
translator = str.maketrans('', '', string.punctuation)
clean_text = text.translate(translator)

# Extract words (sequences of letters)
words = re.findall(r'[a-zA-Z]+', clean_text.lower())

# Count frequencies
counts = Counter(words)

# Sort by count descending, then by word ascending
sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")