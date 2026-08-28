import re
from collections import Counter

def main():
    try:
        # Read the input file
        with open('input/text.txt', 'r', encoding='utf-8') as file:
            text = file.read()

        # Convert to lowercase
        text = text.lower()

        # Use regex to find all sequences of letters (stripping all punctuation/digits)
        words = re.findall(r'[a-z]+', text)

        # Count the frequency of each word
        counts = Counter(words)

        # Sort the results:
        # 1. By count descending (-x[1])
        # 2. By word ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

        # Output each word and its count
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()