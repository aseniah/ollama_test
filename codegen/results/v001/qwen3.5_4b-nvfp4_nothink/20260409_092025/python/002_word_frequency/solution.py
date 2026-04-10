python -c "
import string

words = []
with open('input/text.txt', 'r', encoding='utf-8') as f:
    text = f.read()
    text_lower = text.lower()
    # Split by whitespace and punctuation
    for word in text_lower.split():
        # Remove non-letter characters
        cleaned_word = ''.join(c for c in word if c.isalpha())
        if cleaned_word:
            words.append(cleaned_word)

counts = {}
for w in words:
    counts[w] = counts.get(w, 0) + 1

# Sort by count descending, then alphabetically ascending
sorted_words = sorted(counts.keys(), key=lambda w: (-counts[w], w))

for word in sorted_words:
    print(f\"{word}: {counts[word]}\")
"
done