import re

# Read the file content
with open('input/text.txt', 'r') as f:
    text = f.read()

# Tokenize: split by whitespace and filter out non-letter characters, then lower case
words = []
for token in text.split():
    # Keep only letters (a-z) after converting to lowercase
    cleaned_token = re.sub(r'[^a-z]', '', token.lower())
    if cleaned_token:  # Only add if not empty string
        words.append(cleaned_token)

# Count frequencies
from collections import Counter
word_counts = Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.keys(), key=lambda w: (-word_counts[w], w))

# Output the results
output_lines = [f"{word}: {word_counts[word]}" for word in sorted_words]
print('\n'.join(output_lines))