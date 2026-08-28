import re

def main():
    with open('input/text.txt', 'r') as f:
        text = f.read()
    
    # Convert to lowercase
    text = text.lower()
    
    # Keep only letters (strip punctuation)
    # We want to extract words consisting only of letters
    words = re.findall(r'[a-z]+', text)
    
    # Count frequencies
    freq = {}
    for word in words:
        freq[word] = freq.get(word, 0) + 1
    
    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == '__main__':
    main()