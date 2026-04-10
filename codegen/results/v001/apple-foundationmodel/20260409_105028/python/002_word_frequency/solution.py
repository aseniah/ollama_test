import string
from collections import Counter

def read_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()
    return text

def preprocess_text(text):
    # Convert to lowercase
    text = text.lower()
    # Remove punctuation (keep only letters)
    text = text.translate(str.maketrans('', '', string.punctuation))
    return text

def count_word_frequencies(text):
    # Split text into words
    words = text.split()
    # Count frequencies
    word_count = Counter(words)
    return word_count

def main():
    file_path = 'input/text.txt'
    text = read_file(file_path)
    preprocessed_text = preprocess_text(text)
    word_count = count_word_frequencies(preprocessed_text)

    # Sort by count descending, then alphabetically ascending
    sorted_word_count = sorted(word_count.items(), key=lambda item: (-item[1], item[0]))

    for word, count in sorted_word_count:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()