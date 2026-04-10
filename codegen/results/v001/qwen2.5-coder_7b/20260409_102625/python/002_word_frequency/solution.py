import string

# Read the file and initialize a dictionary to store word counts
with open('input/text.txt', 'r') as file:
    text = file.read().lower()

# Strip punctuation and split into words
words = text.translate(str.maketrans('', '', string.punctuation)).split()

# Count the frequency of each word
word_count = {}
for word in words:
    if word in word_count:
        word_count[word] += 1
    else:
        word_count[word] = 1

# Sort by count descending, then alphabetically ascending
sorted_word_count = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))

# Output one line per unique word in the format: `word: count`
for word, count in sorted_word_count:
    print(f"{word}: {count}")