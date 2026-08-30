using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var content = File.ReadAllText("input/text.txt");
var lower = content.ToLower();

var words = new List<string>();
var sb = new StringBuilder();
foreach (char c in lower) {
    if (char.IsLetter(c)) {
        sb.Append(c);
    } else {
        if (sb.Length > 0) {
            words.Add(sb.ToString());
            sb.Clear();
        }
    }
}
if (sb.Length > 0) words.Add(sb.ToString());

foreach (var g in words.GroupBy(w => w)
                      .OrderByDescending(g => g.Count())
                      .ThenBy(g => g.Key)) {
    Console.WriteLine($"{g.Key}: {g.Count()}");
}