using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {
        var inputFilePath = "input/text.txt";
        var outputFilePath = "output/word_frequency.txt";
        
        // Read the file
        string content = File.ReadAllLines(inputFilePath);
        
        // Initialize a dictionary to store word frequencies
        var wordFrequency = new Dictionary<string, int>();
        
        // Regular expression to match words (letters only)
        const string regex = @"[\w']+";
        
        // Split the content into words and count frequencies
        var words = Regex.Split(content, regex);
        foreach (var word in words)
        {
            word = word.ToLower(); // Convert to lowercase
            wordFrequency[word]++; // Increment count
        }
        
        // Create a list of (word, count) pairs
        var frequencyList = wordFrequency.ToList();
        
        // Sort by count descending, then alphabetically ascending
        frequencyList.Sort((a, b) =>
        {
            if (a.Value != b.Value)
                return b.Value.CompareTo(a.Value); // Descending order by count
            return a.Key.CompareTo(b.Key); // Ascending order by word
        });
        
        // Write sorted word frequencies to output file
        foreach (var entry in frequencyList)
        {
            Console.WriteLine($"{entry.Key}: {entry.Value}");
        }
    }
}