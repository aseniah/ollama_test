import string
from collections import Counter

with open('input/text.txt', 'r') as f:
    text = f.read()

# Replace punctuation with space and split into words
text_no_punct = text.translate(str.maketrans('', '', string.punctuation))
words = text_no_punct.split()

# Filter only alphabetic characters (letters) and convert to lowercase
filtered_words = [''.join(filter(str.isalpha, word.lower())) for word in words]

word_counts = Counter(filtered_words)

# Sort by count descending; for ties, sort alphabetically ascending
sorted_items = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_items:
    print(f"{word}: {count}")