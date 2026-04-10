import re
from collections import Counter

try:
    with open("input/text.txt", "r", encoding="utf-8") as file:
        text = file.read()
except FileNotFoundError:
    # If the file doesn't exist, we simply output nothing or handle gracefully depending on strictness.
    # Assuming the problem implies the file exists. If not found, print nothing.
    pass

# Normalize: convert to lowercase and extract only alphabetic sequences
# re.findall(r'[a-zA-Z]+', text.lower()) effectively strips punctuation and keeps letters
words = re.findall(r'[a-z]+', text.lower())

counter = Counter(words)

# Sort by count descending, then by word ascending
sorted_words = sorted(counter.items(), key=lambda item: (-item[1], item[0]))

for word, count in sorted_words:
    print(f"{word}: {count}")