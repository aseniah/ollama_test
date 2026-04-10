import string
from collections import Counter

def clean_text(text):
    # Convert to lowercase
    text = text.lower()
    # Remove punctuation
    translator = str.maketrans('', '', string.punctuation)
    text = text.translate(translator)
    return text

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as file:
            text = file.read()
        
        cleaned_text = clean_text(text)
        words = cleaned_text.split()
        
        # Count frequency of each word
        word_count = Counter(words)
        
        # Sort by count descending, then alphabetically ascending
        sorted_words = sorted(word_count.items(), key=lambda x: (-x[1], x[0]))
        
        # Output the result
        for word, count in sorted_words:
            print(f"{word}: {count}")
    
    except FileNotFoundError:
        print("The file 'input/text.txt' was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()