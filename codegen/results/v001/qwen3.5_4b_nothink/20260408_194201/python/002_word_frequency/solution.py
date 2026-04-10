python3 <<'EOF'
import re
from collections import Counter

with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()

# Extract words: sequences of letters (a-z, A-Z)
# Convert to lowercase and remove non-letters
words = re.findall(r'[a-zA-Z]+', text.lower())

if not words:
    return

counts = Counter(words)

# Sort by count descending, then by word ascending
sorted_words = sorted(counts.keys(), key=lambda w: (-counts[w], w))

for word in sorted_words:
    print(f"{word}: {counts[word]}")
EOF