import collections
import re

def main():
    try:
        # Read the input file
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
        
        # Convert to lowercase
        text = text.lower()
        
        # Use regex to find all words consisting only of letters
        # This effectively strips all punctuation and splits the text into words
        words = re.findall(r'[a-z]+', text)
        
        # Count the frequency of each word
        counts = collections.Counter(words)
        
        # Sort the results:
        # 1. Primary key: count descending (-item[1])
        # 2. Secondary key: word ascending (item[0])
        sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))
        
        # Output in the format word: count
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()