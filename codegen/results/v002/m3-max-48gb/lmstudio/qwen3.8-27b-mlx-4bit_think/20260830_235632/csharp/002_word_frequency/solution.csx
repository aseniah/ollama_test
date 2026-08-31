using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var text = File.ReadAllText("input/text.txt");
var words = text.Split(new[] { ' ', '\t', '\r', '\n', '\f', '\v' }, StringSplitOptions.RemoveEmptyEntries);

var counts = new Dictionary<string, int>();

foreach (var word in words)
{
    var cleaned = new string(word.ToLower().Where(c => char.IsLetter(c)).ToArray());
    if (string.IsNullOrEmpty(cleaned)) continue;
    if (counts.ContainsKey(cleaned))
        counts[cleaned]++;
    else
        counts[cleaned] = 1;
}

var sorted = counts.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key).ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}