import string
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

words = [word.strip(string.punctuation).lower() for word in text.split()]
words = [word for word in words if word]

counts = Counter(words)

for word, count in sorted(counts.items(), key=lambda x: (-x[1], x[0])):
    print(f"{word}: {count}")
