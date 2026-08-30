import re
from collections import Counter

def main():
    # Read the file
    with open('input/text.txt', 'r') as f:
        text = f.read()
    
    # Extract words (keep only letters) and convert to lowercase
    words = re.findall(r'[a-z]+', text.lower())
    
    # Count word frequencies
    word_counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(word_counts.keys(), key=lambda w: (-word_counts[w], w))
    
    # Output results
    for word in sorted_words:
        print(f"{word}: {word_counts[word]}")

if __name__ == "__main__":
    main()