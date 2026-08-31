import re

with open("input/text.txt", "r") as f:
    text = f.read()

# Convert to lowercase, split into words, strip non-letter characters
words = re.findall(r"[a-z]+", text.lower())

# Count frequencies
freq = {}
for word in words:
    freq[word] = freq.get(word, 0) + 1

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")