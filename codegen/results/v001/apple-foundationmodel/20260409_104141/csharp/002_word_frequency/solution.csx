using System;
using System.IO;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

// Function to clean the text by removing punctuation and converting to lowercase
string CleanText(string text) {
    // Remove punctuation using regular expressions
    var translator = new RegexTokenizer();
    translator.RemovePunctuation();
    return translator.Tokenize(text).ToString();
}

// Function to count word frequencies
Dictionary<string, int> CountWords(string text) {
    var wordCount = new Dictionary<string, int>();
    var words = text.Split(' ');
    
    foreach (var word in words) {
        // Convert word to lowercase and remove punctuation
        string cleanedWord = CleanText(word).ToLower();
        if (!string.IsNullOrWhiteSpace(cleanedWord)) {
            wordCount[cleanedWord]++;
        }
    }
    
    return wordCount;
}

// Main function
void Main() {
    if (Args.Length < 1) {
        Console.WriteLine("Please provide the input file path as an argument.");
        return;
    }

    string inputFilePath = Args[0];

    try {
        // Read the text from the input file
        string text = File.ReadAllText(inputFilePath);

        // Count the frequency of each word
        var wordCount = CountWords(text);

        // Sort the words by count descending, then alphabetically ascending
        var sortedWords = wordCount.OrderByDescending(x => x.Value).ThenBy(x => x.Key);

        // Output the results
        foreach (var entry in sortedWords) {
            Console.WriteLine($"{entry.Key}: {entry.Value}");
        }
    } catch (Exception ex) {
        Console.WriteLine($"An error occurred: {ex.Message}");
    }
}