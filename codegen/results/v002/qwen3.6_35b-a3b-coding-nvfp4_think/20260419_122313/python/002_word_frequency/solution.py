import re
from collections import Counter

def solve():
    # Read the content of the file
    try:
        with open('input/text.txt', 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert all text to lowercase
    text = text.lower()
    
    # Split the text into words based on whitespace
    words = text.split()
    
    # Process each word to keep only letters
    cleaned_words = []
    for word in words:
        # Remove any character that is not a lowercase letter
        # This handles punctuation stripping
        clean_word = re.sub(r'[^a-z]', '', word)
        if clean_word:
            cleaned_words.append(clean_word)
            
    # Count the frequency of each word
    counts = Counter(cleaned_words)
    
    # Sort the words based on the rules:
    # 1. Count descending
    # 2. Word alphabetically ascending
    # The key lambda function returns a tuple: (-count, word)
    # Python sorts tuples element by element.
    # -count ensures descending order for counts.
    # word ensures ascending alphabetical order for ties.
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output the results in the specified format
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()