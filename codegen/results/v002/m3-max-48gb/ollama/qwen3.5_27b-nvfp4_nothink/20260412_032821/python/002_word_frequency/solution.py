import re
from collections import Counter

def main():
    file_path = 'input/text.txt'
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()
    
    # Strip punctuation: keep only letters and whitespace
    text = re.sub(r'[^a-z\s]', ' ', text)
    
    # Split into words (split by whitespace)
    words = text.split()
    
    # Count frequency
    word_counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(word_counts.items(), key=lambda item: (-item[1], item[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()