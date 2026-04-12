import string

# Read the file and count word frequencies
word_counts = {}
with open('input/text.txt', 'r') as file:
    for line in file:
        words = line.lower().split()
        for word in words:
            word = word.strip(string.punctuation)
            if word:
                word_counts[word] = word_counts.get(word, 0) + 1

# Sort by count descending, then alphabetically ascending
sorted_word_counts = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

# Output the results
for word, count in sorted_word_counts:
    print(f"{word}: {count}")