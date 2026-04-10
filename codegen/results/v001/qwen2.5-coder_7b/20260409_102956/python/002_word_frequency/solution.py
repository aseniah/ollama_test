import string

# Function to clean the text and return a list of words
def clean_text(text):
    # Convert to lowercase
    text = text.lower()
    
    # Remove punctuation using str.translate
    translator = str.maketrans('', '', string.punctuation)
    text = text.translate(translator)
    
    # Split into words
    words = text.split()
    
    return words

# Read the file and clean the text
with open('input/text.txt', 'r') as file:
    text = file.read()
words = clean_text(text)

# Count the frequency of each word
word_count = {}
for word in words:
    if word in word_count:
        word_count[word] += 1
    else:
        word_count[word] = 1

# Sort by count descending, then alphabetically ascending
sorted_word_count = sorted(word_count.items(), key=lambda item: (-item[1], item[0]))

# Output the result
for word, count in sorted_word_count:
    print(f"{word}: {count}")