import re
import sys
from collections import Counter

def solve():
    """
    Reads the file input/text.txt, counts the frequency of each word 
    after cleaning, and prints the results sorted by count (descending) 
    and then alphabetically (ascending).
    """
    file_path = 'input/text.txt'
    
    try:
        # Attempt to read the input file
        with open(file_path, '                ' + 'r') as file:
            content = file.read()
        
        # Convert all text to lowercase as per rules
        content = content.lower()
        
        # Use regular expressions to find all sequences of letters.
        # This effectively strips all punctuation and keeps only letters,
        # treating any non-letter character as a delimiter.
        words = re.findall(r'[a-z]+', content)
        
        # Count the frequency of each word
        word_counts = Counter(words)
        
        # Sorting requirements:
        # 1. Sort by count descending (-x[1])
        # 2. For ties, sort by word alphabetically ascending (x[0])
        sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Output the results in the format 'word: count'
        for word, count in sorted_words:
            sys.stdout.write(f"{word}: {count}\n")
            
    except FileNotFoundError:
        # If the file is not found, the program exits without outputting anything
        pass
    except Exception:
        # Ensure no other errors or logs are printed to stdout
        pass

if __name__ == "__main__":
    solve()