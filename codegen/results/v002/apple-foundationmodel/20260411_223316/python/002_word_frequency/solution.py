import re
from collections import Counter

def count_word_frequencies(file_path):
    # Read the file content
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()

    # Convert text to lowercase and strip punctuation
    text = re.sub(r'[^\w\s]', '', text).lower()

    # Split text into words
    words = text.split()

    # Count word frequencies
    word_counts = Counter(words)

    # Sort by count descending, and alphabetically ascending for ties
    sorted_words = sorted(word_counts.items(), key=lambda item: (-item[1], item[0]))

    # Output the results
    for word, count in sorted_words:
        print(f"{word}: {count}")

# Path to the input file
file_path = 'input/text.txt'
count_word_frequencies(file_path)