import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            content = f.read()
    except FileNotFoundError:
        content = ""
    
    # Convert to lowercase
    content = content.lower()
    
    # Keep only letters and spaces? We need words (letters only). Strip punctuation, keep only letters.
    # A word is a sequence of letters. Punctuation attached becomes separate? 
    # "Strip all punctuation (keep only letters)" means after stripping punctuation, words are formed from letters.
    # Common approach: extract sequences of letters.
    words = re.findall(r'[a-z]+', content)
    
    if not words:
        return
    
    counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()