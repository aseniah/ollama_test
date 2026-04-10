import string
from collections import Counter

with open("input/text.txt", "r") as f:
    text = f.read()

words = text.lower().split()
words = [w.translate(str.maketrans("", "", string.punctuation)) for w in words]
words = [w for w in words if w]

counts = Counter(words)

for word, count in sorted(counts.items(), key=lambda x: (-x[1], x[0])):
    print(f"{word}: {count}")
