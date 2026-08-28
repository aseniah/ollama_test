import collections
import string

# Read the file
with open('input/text.txt', 'r') as file:
    content = file.read()

# Convert to lowercase
content = content.lower()

# Remove punctuation
translator = str.maketrans('', '', string.punctuation)
content = content.translate(translator)

# Split into words and filter out empty strings
words = [word for word in content.split() if word]

# Count frequency
word_count = collections.Counter(words)

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))

# Output results
for word, count in sorted_words:
    print(f"{word}: {count}")