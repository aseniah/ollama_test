import re
from collections import Counter

def main():
    with open('input/text.txt', 'r') as f:
        text = f.read()
        
    # Convert to lowercase and extract only sequences of letters
    words = re.findall(r'[a-z]+', text.lower())
    
    # Count frequencies
    counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()