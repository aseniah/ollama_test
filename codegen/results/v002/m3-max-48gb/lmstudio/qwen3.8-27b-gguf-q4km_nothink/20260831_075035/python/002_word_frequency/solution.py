import re

with open('input/text.txt', 'r') as f:
    text = f.read()

# Convert to lowercase
text = text.lower()

# Strip all punctuation (keep only letters) - also remove digits and whitespace
# We need to extract words consisting only of letters
words = re.findall(r'[a-z]+', text)

# Count frequency
freq = {}
for word in words:
    if word:  # make sure it's not empty
        freq[word] = freq.get(word, 0) + 1

# Sort by count descending, then alphabetically ascending for ties
sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")