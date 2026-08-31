import re

def main():
    with open('input/text.txt', 'r') as f:
        text = f.read()
    
    # Convert to lowercase
    text = text.lower()
    
    # Extract only letters (and whitespace) to preserve word boundaries
    # We want to keep letters and spaces, remove everything else
    text = re.sub(r'[^a-z\s]', '', text)
    
    # Split into words
    words = text.split()
    
    # Count frequencies
    freq = {}
    for word in words:
        if word:  # non-empty
            freq[word] = freq.get(word, 0) + 1
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(freq.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == '__main__':
    main()