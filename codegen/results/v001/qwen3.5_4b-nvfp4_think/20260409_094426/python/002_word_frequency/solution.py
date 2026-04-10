import re

# Read and process file
with open('input/text.txt', 'r') as f:
    text = f.read()

# Extract words (keep only letters)
words = re.findall(r'[a-zA-Z]+', text.lower())

# Count frequencies
word_counts = {}
for word in words:
    word_counts[word] = word_counts.get(word, 0) + 1

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")