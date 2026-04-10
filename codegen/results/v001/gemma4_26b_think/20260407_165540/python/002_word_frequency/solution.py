import re
from collections import Counter
import sys

def main():
    file_path = 'input/text.txt'
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read().lower()
        
        # Use regex to find all sequences of letters.
        # This effectively strips all punctuation, numbers, and whitespace,
        # treating any non-letter character as a delimiter.
        words = re.findall(r'[a-z]+', content)
        
        # Count the frequency of each word
        counts = Counter(words)
        
        # Sort by count descending (-x[1]), then by word alphabetically ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Output the results in the format 'word: count'
        for word, count in sorted_words:
            sys.stdout.write(f"{word}: {count}\n")
            
    except FileNotFoundError:
        # If the file does not exist, we exit silently as per the "Do not output anything else" rule
        pass
    except Exception:
        # Handle other potential errors silently to adhere to output constraints
        pass

if __name__ == "__main__":
    main()