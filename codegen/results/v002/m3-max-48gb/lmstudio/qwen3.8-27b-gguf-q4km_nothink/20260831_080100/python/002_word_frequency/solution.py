import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return
    
    # Convert to lowercase
    text = text.lower()
    
    # Extract only alphabetic words
    words = re.findall(r'[a-z]+', text)
    
    # Count frequencies
    counter = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(counter.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()