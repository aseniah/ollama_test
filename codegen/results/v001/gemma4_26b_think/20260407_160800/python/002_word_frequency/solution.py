import re
from collections import Counter

def main():
    try:
        # Read the input file
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            content = f.read()

        # Convert all words to lowercase
        content = content.lower()

        # Use regex to find all sequences of letters. 
        # This effectively strips all punctuation and keeps only letters.
        # It also handles splitting words separated by punctuation (e.g., "hello,world" -> ["hello", "world"])
        words = re.findall(r'[a-z]+', content)

        # Count the frequency of each word
        counts = Counter(words)

        # Sort by count descending (-x[1]), then by word alphabetically ascending (x[0])
        sorted_counts = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

        # Output the results in the required format
        for word, count in sorted_counts:
            print(f"{word}: {count}")

    except FileNotFoundError:
        # If the file does not exist, the script exits silently per the "Do not output anything else" rule
        pass

if __name__ == "__main__":
    main()