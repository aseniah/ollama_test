import string
from collections import Counter

def count_word_frequencies(file_path):
    # Read the file
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()
    
    # Convert text to lowercase and remove punctuation
    text = text.lower()
    text = text.translate(str.maketrans('', '', string.punctuation))
    
    # Split the text into words
    words = text.split()
    
    # Count the frequency of each word
    word_counts = Counter(words)
    
    # Sort words by count descending, and alphabetically ascending for ties
    sorted_words = sorted(word_counts.items(), key=lambda item: (-item[1], item[0]))
    
    # Output the results
    for word, count in sorted_words:
        print(f"{word}: {count}")

# Assuming the file is located in the input directory
count_word_frequencies('input/text.txt')