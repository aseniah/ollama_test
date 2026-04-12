import string
from collections import Counter

def process_text(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()
    
    # Convert to lowercase
    text = text.lower()
    
    # Remove punctuation
    translator = str.maketrans('', '', string.punctuation)
    text = text.translate(translator)
    
    # Split text into words
    words = text.split()
    
    # Count frequencies
    word_count = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))
    
    # Output the results
    for word, count in sorted_words:
        print(f"{word}: {count}")

# Call the function with the file path
process_text('input/text.txt')