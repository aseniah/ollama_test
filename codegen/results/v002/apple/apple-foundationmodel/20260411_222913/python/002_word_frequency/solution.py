import string
from collections import Counter

def clean_text(text):
    # Convert to lowercase
    text = text.lower()
    # Remove punctuation using string.punctuation
    text = ''.join(char for char in text if char not in string.punctuation)
    return text

def count_word_frequencies(file_path):
    # Read the file
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()
    
    # Clean the text
    cleaned_text = clean_text(text)
    
    # Split text into words
    words = cleaned_text.split()
    
    # Count word frequencies
    word_count = Counter(words)
    
    # Sort words first by count descending, then alphabetically ascending
    sorted_words = sorted(word_count.items(), key=lambda item: (-item[1], item[0]))
    
    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

# Call the function with the specified file path
count_word_frequencies('input/text.txt')