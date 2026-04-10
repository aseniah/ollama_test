python
from collections import Counter
import string

with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read().lower()

# Remove punctuation: keep only letters and spaces
cleaned_text = ''.join(ch for ch in text if ch.isalpha() or ch.isspace())

# Tokenize and count
words = cleaned_text.split()
count = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(count.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")
</think>