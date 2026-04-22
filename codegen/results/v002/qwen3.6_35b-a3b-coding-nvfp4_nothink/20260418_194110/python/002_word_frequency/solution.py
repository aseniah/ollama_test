import re
from collections import Counter

def main():
    with open('input/text.txt', 'r') as f:
        text = f.read()
    
    # Convert to lowercase
    text = text.lower()
    
    # Replace all non-letter characters with spaces
    # This effectively strips punctuation and splits on any non-letter
    cleaned_text = re.sub(r'[^a-z]', ' ', text)
    
    # Split into words
    words = cleaned_text.split()
    
    # Filter out empty strings (if any)
    words = [w for w in words if w]
    
    # Count frequencies
    counter = Counter(words)
    
    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == '__main__':
    main()