import re
from collections import Counter

def count_word_frequency():
    try:
        # Read the input file
        with open('input/text.txt', 'r') as file:
            text = file.read()
        
        # Convert to lowercase
        text = text.lower()
        
        # Strip all punctuation (keep only letters)
        # re.findall(r'[a-z]+', text) extracts all contiguous sequences of lowercase letters
        words = re.findall(r'[a-z]+', text)
        
        # Count frequency of each word
        counts = Counter(words)
        
        # Sort by count descending; for ties, sort alphabetically ascending
        # The key (-item[1], item[0]) achieves this:
        # -item[1] sorts the count in descending order
        # item[0] sorts the word in ascending order
        sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))
        
        # Output in the format "word: count"
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        # If file is missing, do nothing per the "Do not output anything else" rule,
        # though typically this would be handled by the environment.
        pass

if __name__ == "__main__":
    count_word_frequency()