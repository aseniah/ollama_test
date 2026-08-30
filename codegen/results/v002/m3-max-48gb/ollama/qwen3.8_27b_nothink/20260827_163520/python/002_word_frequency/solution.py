import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return
    
    # Convert to lowercase
    text = text.lower()
    
    # Keep only letters and spaces (strip punctuation)
    # Remove all non-letter characters except spaces
    cleaned = re.sub(r'[^a-z\s]', '', text)
    
    # Split into words
    words = cleaned.split()
    
    # Count frequencies
    word_counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == '__main__':
    main()