import re
from collections import Counter

def solve():
    """
    Reads 'input/text.txt', counts the frequency of each word,
    and prints the results according to the specified rules.
    """
    file_path = 'input/text.txt'
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            text = f.read()
            
        # 1. Convert all words to lowercase
        text = text.lower()
        
        # 2. Strip all punctuation (keep only letters)
        # Using regex to find all sequences of characters from a to z.
        # This effectively treats any non-alphabetic character as a delimiter.
        words = re.findall(r'[a-z]+', text)
        
        # 3. Count the frequency of each word
        counts = Counter(words)
        
        # 4. Sort by count descending; for ties, sort alphabetically ascending
        # -x[1] handles descending order for the count
        # x[0] handles ascending order for the word
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # 5. Output in the format: word: count
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        # If the file is not found, we do not output anything as per the instructions.
        pass
    except Exception:
        # Standard practice to prevent unexpected output in automated environments.
        pass

if __name__ == "__main__":
    solve()