var inputFile = "input/input.csv";
var outputFile = "input/expected_format.json"; // For reference, we will output to stdout instead.

var lines = File.ReadAllLines(inputFile);
var headerLine = lines[0].Split(',');

var people = new List<Dictionary<string, string>>();

foreach (var line in lines.Skip(1)) {
    var values = line.Split(',');
    var firstName = values[0].Trim();
    var lastName = values[1].Substring(values[1].IndexOf('-') + 1).Trim().Replace("/", "-"); 
    // The Birthday is DD/MM/YYYY in input, but output requires YYYY-MM-DD format.
    // We need to parse "10/9/1940" -> "1940-10-09".
    var birthdayStr = values[2];
    var day = int.Parse(birthdayStr.Substring(0, 2));
    var month = int.Parse(birthdayStr.Substring(3, 1));
    var year = int.Parse(birthdayStr.Substring(6));
    
    // Construct the input birthday string as YYYY-MM-DD for reference in calculation logic
    string inputBirthdayFormatted = $"{year}-{month:D2}-{day:D2}";
    
    // Parse father, mother, brother, sister. If null or empty, skip them or handle gracefully.
    var fatherName = values[3]?.Trim();
    var motherName = values[4]?.Trim();
    var brotherName = values[5]?.Trim();
    var sisterName = values[6]?.Trim();
    
    // We need to parse the Birthday as YYYY-MM-DD for calculation, but the input is DD/MM/YYYY.
    // Let's re-parse carefully based on input format "DD/M/YYYY".
    string birthdayInput = values[2]; 
    int dayB = int.Parse(birthdayInput.Substring(0, 2));
    int monthB = int.Parse(birthdayInput.Substring(3, 1));
    int yearB = int.Parse(birthdayInput.Substring(6));
    
    // Calculate age as of July 1, 2025 (2025-07-01).
    int currentYear = 2025;
    int currentMonth = 7;
    int age = currentYear - yearB;
    
    // Adjust for birthday month vs current month/day.
    if (currentMonth < monthB || (currentMonth == monthB && day < dayB)) {
        age--;
    }
    
    var relatives = new List<Dictionary<string, string>>();
    
    if (!string.IsNullOrEmpty(fatherName)) {
        // Check surname match? The expected output shows "Alfred Lennon" where input is just "Alfred Lennon".
        // Wait, the CSV has full names.
        // Let's assume the family members are strings in the CSV columns.
        var relativesData = new List<(string F, string L, string R)>();
        
        if (!string.IsNullOrEmpty(fatherName)) {
            // The CSV format is: Father (Full Name?), Mother, Brother, Sister.
            // Wait, looking at input.csv: "Alfred Lennon", "Julia Stanley". These are full names.
            // However, the expected output JSON structure separates FirstName and LastName.
            // Let's parse the Full Name from the CSV column if it contains a space.
            // Input format seems to be "FirstName LastName" for relatives too? 
            // No, example: "Alfred Lennon". Yes, likely first + last.
            
            var parts = fatherName.Split(' ');
            if (parts.Length >= 2) {
                relativesData.Add((new Dictionary<string, string>() {{"FirstName", parts[0].Trim()}, {"LastName", parts[1].Trim()}, {"Relationship", "Father"}}));
            }
        }
        
        if (!string.IsNullOrEmpty(motherName)) {
             var parts = motherName.Split(' ');
             if (parts.Length >= 2) {
                relativesData.Add((new Dictionary<string, string>() {{"FirstName", parts[0].Trim()}, {"LastName", parts[1].Trim()}, {"Relationship", "Mother"}}));
            }
        }
        
        // Note: The input CSV columns might just be last names? 
        // Let's look at row 2 (James): "Jim McCartney" (Father), "Mary McCartney" (Mother).
        // Row 3 (Ringo): "Richard Starkey" (Father), "Elsie Gleave" (Mother).
        // It seems the column contains Full Name.
        var lastNames = new List<string>();
        
        if (string.IsNullOrEmpty(fatherName)) { fatherName = ""; }
        if (string.IsNullOrEmpty(motherName)) { motherName = ""; }
        if (string.IsNullOrEmpty(brotherName)) { brotherName = ""; }
        if (string.IsNullOrEmpty(sisterName)) { sisterName = ""; }

        // Parse the names into first and last.
        var parseName = (name) => {
            var parts = name.Split(' ');
            return parts.Length >= 2 ? new Tuple<string,string>(parts[0].Trim(), parts[1].Trim()) : new Tuple<string,string>(name, "");
        };

        // Let's construct the output structure.
        // Wait, I need to be careful. The prompt says "Calculate ages as of July 1, 2025".
        // The expected JSON shows Age calculated correctly (e.g., John Lennon born 1940-10-09, age 40 on 2025-07-01).
        // Let's re-check logic: Born Oct 9, 1940. As of July 1, 2025. 
        // Has he had his birthday this year? No (Oct > Jul). So age = 2025 - 1940 - 1 = 39.
        // Wait, the expected output says Age: 40 for John Lennon!
        // Why 40? 
        // Maybe the calculation is simply CurrentYear - BornYear? 
        // 2025 - 1940 = 85. No.
        // Wait, let's look at dates again.
        // John: 10/9/1940 (Oct 9). Age on July 1, 2025? 
        // If he was born in 1940, and it's currently 2025. 
        // Birthday is Oct 9. Current date July 1. 
        // So he has NOT had his birthday yet in 2025. Age = 2024 + (some logic)?
        // Actually, let's count years: 2025 - 1940 = 85? No, expected is 40.
        // Ah, maybe I misread the input date format or the expected output year?
        // Expected Output: "Age": 40. 
        // Born 1940. Current 2025. 
        // (2025 - 1940) = 85. Why 40?
        // Is it possible the input date is Year/Month/Day? 
        // Input: 10/9/1940 -> Oct 9, 1940. 
        // Maybe the current year isn't 2025? No, prompt says "Calculate ages as of July 1, 2025".
        
        // Wait, look at James McCartney: Born 6/18/1942 (Jun 18). Expected Age: 83.
        // 2025 - 1942 = 83. 
        // So for James, he IS his birthday or past it? Jun 18 is before Jul 1. So yes, age = 2025 - 1942 = 83.
        
        // Look at John Lennon: Born 10/9/1940 (Oct 9). Expected Age: 40? 
        // Wait, looking at the expected JSON text provided by user... 
        // "{ \"Birthday\": \"1940-10-09\", \"Age\": 40 }"
        // Is it possible the year in expected output is wrong? Or my math?
        // 2025 - 1940 = 85. If age is 40, then the calculation would be 2025 - 1940 + something negative? No.
        // Wait, could the input year be different? 
        // 10/9/1940. 
        // Maybe the target date is not 2025? 
        // Prompt says: "Calculate ages as of July 1, 2025".
        
        // Is it possible I am misinterpreting the "Age" column in expected format? 
        // Let's check Ringo Starr. Born 7/7/1940 (Jul 7). Expected Age: 84.
        // Current date: Jul 1, 2025. 
        // Born Jul 7. Has NOT had birthday yet this year (Jul 1 < Jul 7). 
        // So age should be (2024) - 1940 + adjustment?
        // Years passed: 2025 - 1940 = 85. Minus 1 because birthday not passed. Age = 84.
        
        // Now back to John Lennon. Born Oct 9, 1940. 
        // Current date Jul 1, 2025. 
        // Years passed: 85. Birthday (Oct) is in future relative to Jul 1. 
        // So age should be 84? 
        // But expected output says 40.
        
        // WAIT! 
        // Maybe the input date format is MM/Day/Year? No, "1940" at end suggests Year.
        // Is it possible the prompt implies a different base year?
        // Or maybe... wait, did I miss something in the input CSV headers?
        // Name,Birthday,Died,...
        
        // Let's look at George Harrison. Born 2/25/1943 (Feb 25). Died 11/29/2001.
        // Age on July 1, 2025? 2025 - 1943 = 82. 
        // Feb 25 < Jul 1. Birthday passed. Age = 82.
        // But expected output says: "Age": 58.
        
        // This is extremely strange. All ages are roughly 1/3 or half of what they should be mathematically (85 vs 40, 83 vs 83, 84 vs 84, 82 vs 58).
        // Actually: 
        // John: Expected 40. Calc 84 (if birthday passed) or 83 (if not). 40 is way off.
        // James: Expected 83. Calc 83. (Correct math!).
        // Ringo: Expected 84. Calc 84 (if birthday passed? No, Jul 7 > Jul 1, so not passed. 2025-1940=85-1=84). Correct!
        // George: Expected 58. Calc 82. 
        // Why is George different from James/Ringo?
        
        // Hypothesis: Is the input year actually Year/Month/Day?
        // "1940" -> 19, 40? No.
        
        // Let's re-read the prompt carefully. "Calculate ages as of July 1, 2025".
        // Maybe I am misreading the numbers in the expected output? 
        // The prompt says: "expected_format.json ... shows the expected output — use it to infer the transformation rules".
        // If the provided example has a bug (John Lennon age 40 instead of ~83), I must follow the data OR find the logic.
        // Is there a relation between Name and Age? 
        // John -> 40. James -> 83. Ringo -> 84. George -> 58.
        // Maybe the year in the CSV is not the birth year?
        // No, "1940" is clearly birth year for Beatles members.
        
        // WAIT! 
        // Is it possible the prompt has a typo in the "expected output" example provided by the user, and I should ignore the anomaly in John Lennon and George Harrison if they don't follow the logic?
        // Or maybe... 
        // Let's check the date format again. 
        // Input: 10/9/1940. 
        // If MM/DD/YYYY -> 10/9/1940 = Oct 9, 1940. Age on Jul 1, 2025: 84 (if birthday passed) or 83? 
        // Born Oct 9. July 1 is before Oct 9. So age = 2025 - 1940 - 1 = 84.
        // Expected: 40.
        
        // Is it possible the "Age" field in the JSON is actually something else? No, it's integer.
        
        // Let's look at the pattern again. 
        // Maybe the years are different? 
        // What if the target year is NOT 2025? 
        // If Age = 40 for born 1940 -> Year = 1980. 
        // If Age = 83 for born 1942 -> Year = 2025. 
        // If Age = 84 for born 1940 -> Year = 2024? (If birthday passed) or 2025?
        // If Age = 58 for born 1943 -> Year = 2001 (Died date)? 
        // No, "Died" is a date.
        
        // Okay, let's assume the prompt instruction "Calculate ages as of July 1, 2025" is the **truth**, and the provided example JSON might contain errors or specific quirks I am missing. 
        // HOWEVER, usually these tasks want you to implement the logic described in the text ("Calculate ages as of July 1, 2025") correctly, regardless of potential typos in the sample output if they contradict the rule strictly.
        // BUT, the prompt says "use it to infer the transformation rules". 
        // If John Lennon is 40, and born 1940... 1980? No.
        // Maybe the input year is Day/Month/Year? 
        // 10/9/1940 -> 1940-10-09? Same thing.
        
        // Let's try: Age = CurrentYear - BirthYear - Adjust.
        // If we target July 1, 2025.
        // John (Oct 9): Not yet birthday. Age = 2025 - 1940 - 1 = 84. 
        // Ringo (Jul 7): Already had birthday? Jul 1 < Jul 7. NO. 
        // Born Jul 7. Current Jul 1. Age = 2025 - 1940 - 1 = 84.
        // So both should be 84 if born in previous year and birthday hasn't passed? 
        // Wait, Ringo born Jul 7. Today Jul 1. He has NOT turned 85 yet. 
        // 2025 - 1940 = 85. Minus 1 for birthday = 84.
        // So Ringo age 84 matches the math (if we assume current year is 2025 and he hasn't had birthday).
        
        // John born Oct 9. Today Jul 1. 
        // 2025 - 1940 = 85. Minus 1 = 84.
        // Expected 40. 
        // Difference is 44 years.
        
        // Maybe the "Age" column in the JSON example is actually referring to age at death? 
        // John died Dec 8, 1980. Age = 40 (1980 - 1940). 
        // Ah! The prompt says "Calculate ages as of July 1, 2025". But the example shows ages matching age at death for some?
        // John died 1980. Age = 40. Correct.
        // Ringo (Live). Born 1940. Died? Unknown (He is alive in this dataset?). 
        // Wait, Ringo Starr (Jul 7) -> Expected 84. 
        // If current year is 2025. Born 1940. Age = 84 (if birthday passed). 
        // Since Ringo born Jul 7, and today Jul 1. He has NOT had birthday. So 83?
        // Wait, if today is July 1, and he was born July 7, he is still 83 (turning 84 soon).
        // But expected output says 84. 
        // Maybe the date is July 15? No, prompt says July 1.
        
        // Let's re-evaluate George Harrison. Born Feb 25, 1943. Died Nov 29, 2001. 
        // Age at death: 2001 - 1943 = 58. 
        // Expected output says 58.
        
        // Conclusion: 
        // The rule for "Age" is either Age at Death OR Age as of Current Date (2025).
        // John Lennon: Died 1980. Age = 40. (Matches "Age at death").
        // James McCartney: Alive? No date of death listed (null). Age on 2025-07-01? 
        // Born 1942. 2025 - 1942 = 83. (If birthday passed). 
        // James born Jun 18. Current Jul 1. Birthday passed. Age = 83. Matches Expected.
        // Ringo Starr: Alive? No date of death listed. 
        // Born Jul 7. Current Jul 1. Birthday NOT passed. Age should be 83? 
        // But Expected says 84. 
        // Maybe "July 1" is inclusive? Or maybe my birth month logic is wrong?
        // If born Jul 7, and today July 1, age = (2025-1940) - 1 = 83.
        // Why 84? 
        // Maybe the calculation ignores birthday if born AFTER current date? No, that's standard logic.
        // Wait, is it possible the "Age" in the expected JSON for Ringo is wrong in the prompt text? Or maybe the date "July 1, 2025" implies something else?
        
        // Let's look at George again. 
        // Born Feb 25. Died Nov 29. Age at death = 58. Matches Expected.
        
        // It seems the rule is: 
        // If person has died (Died column not null), calculate age as of date of death.
        // If person is alive, calculate age as of July 1, 2025.
        // But wait, what if a person died BEFORE the reference date? Then use death date?
        // John died 1980. Age 40. Matches death logic.
        // George died 2001. Age 58. Matches death logic.
        // James (Alive). Age 83. Matches "Current Date" logic (if birthday passed).
        // Ringo (Alive). Born Jul 7. Expected 84. 
        // If we assume "Alive" logic: 2025 - 1940 = 85. 
        // Why 84? Maybe because it's July 1 and he's born Jul 7? 
        // So 2025 - 1940 = 85. Minus 1 (birthday not passed) = 84? 
        // Yes! 
        // James: Born Jun 18. Current Jul 1. Birthday passed. Age = 83? Wait. 
        // 2025 - 1942 = 83. 
        // If birthday passed, age is current year - birth year. 
        // So James = 83. Correct.
        // Ringo: Born Jul 7. Current Jul 1. Birthday NOT passed. Age = (Current Year - Birth Year) - 1 = 85 - 1 = 84. Correct!
        
        // So the logic is consistent now.
        // Logic: 
        // 1. Check if Died column is not null/empty.
        //    - If yes, calculate age as (Year of Death) - Birth Year.
        //      Note: Need to handle month/day for death age too? Usually age at death ignores specific date of day unless birthday passed? 
        //      But here, we just subtract years. 2001-1943 = 58. 1980-1940 = 40. It works perfectly without month check for death.
        //    - If no, calculate age as of July 1, 2025.
        //      Logic: 
        //      Start with (2025 - BirthYear). 
        //      If Birthday Month < 7 OR (Month == 7 AND Day <= 1): Age = StartValue.
        //      Else: Age = StartValue - 1.
        // Wait, for James (Jun 18): 
        // June < July. So birthday passed. Age = 2025-1942 = 83. Correct.
        // For Ringo (Jul 7): 
        // July == July. Day 7 > 1. Birthday NOT passed. Age = (2025-1940) - 1 = 84. Correct.
        // For John (Oct 9): 
        // Oct > July. Not passed. Age = (2025-1940) - 1 = 84. 
        // BUT Expected says 40.
        // So for John, he is DEAD. The rule "If Died column not null" applies.
        // If Dead: Age = DeathYear - BirthYear. (Ignoring month/day of death).
        
        // Final Logic Summary:
        // 1. Parse Birthday (DD/MM/YYYY) -> Year, Month, Day.
        // 2. Check "Died" column.
        //    - If Died exists (not null/empty): Age = Year(Died) - BirthYear.
        //    - Else (Alive): 
        //       YearTarget = 2025.
        //       TargetAge = YearTarget - BirthYear;
        //       If (CurrentMonth > BirthMonth OR (CurrentMonth == BirthMonth AND CurrentDay >= BirthDay)):
        //           Age = TargetAge
        //       Else:
        //           Age = TargetAge - 1
        
        // Now for Relatives:
        // Structure: [{FirstName, LastName, Relationship}]
        // Source columns: Father, Mother, Brother, Sister.
        // These are full names (Space separated).
        // Format output as [{"F", "L", "R"}, ...]
        
        // Wait, the expected JSON for relatives order? 
        // John: Father, Mother. Order matches column order? Yes.
        // James: Father, Mother, Brother. Order matches column order? Yes.
        // Ringo: Father, Mother, Sister. (Brother is null). Matches.
        // George: Father, Mother, Brother, Sister. Matches.
        
        // So we iterate columns 3 to 6 (Father to Sister). If not null, parse name into F, L, add rel.
        
        var targetDateYear = 2025;
        var targetDateMonth = 7;
        var targetDateDay = 1;
        
        var peopleList = new List<Dictionary<string, object>>();
        var jsonOptions = new JsonSerializerOptions { WriteIndented = false }; // The prompt example isn't pretty printed? 
        // Wait, the example JSON IS indented.
        // "   {" ... }"
        // So we should format nicely.
        
        foreach (var line in lines.Skip(1)) {
            var cols = line.Split(',');
            
            var firstName = cols[0].Trim();
            var lastName = cols[1]; 
            // Note: Last name in input is "1940-10-09"? No, that's col 2.
            // Col 1 is Name. 
            // Wait, looking at CSV header: Name,Birthday,Died...
            // Row 1: John Winston Lennon,10/9/1940,...
            // So FirstName = "John". LastName = "Lennon" (from "John Winston Lennon"? Split by space?)
            // Wait, the prompt expects "FirstName": "John", "LastName": "Lennon".
            // Is "Winston" ignored? 
            // Let's check input: "John Winston Lennon". 
            // Maybe we split by ' ' and take middle? Or just last part?
            // Beatles names: John Lennon, Paul McCartney (James McCartney?), Ringo Starr, George Harrison.
            // Col 1 for James is "James Paul McCartney".
            // Expected LastName: McCartney. First: James.
            // So the logic is: First Word = FirstName. Last Word = LastName? 
            // But what about middle names? 
            // John Winston Lennon -> First: John, Last: Lennon. (Ignore Winston).
            // Ringo Starr -> First: Ringo, Last: Starr.
            // George Harrison -> First: George, Last: Harrison.
            // So likely: Split by space. Ignore middle names? Or just take first and last.
            var parts = firstName.Split(new[] {' ' }, StringSplitOptions.RemoveEmptyEntries);
            var fName = parts[0];
            var lName = parts[parts.Length - 1];
            
            string birthdayStr = cols[2];
            int dayB = int.Parse(birthdayStr.Substring(0, 2));
            int monthB = int.Parse(birthdayStr.Substring(3, 1));
            int yearB = int.Parse(birthdayStr.Substring(6));
            
            // Determine Age
            int age = 0;
            bool isDead = !string.IsNullOrEmpty(cols[3]);
            
            if (isDead) {
                var dieParts = cols[3].Split(' ', StringSplitOptions.RemoveEmptyEntries);
                string dieYearStr = dieParts[dieParts.Length - 1]; // Assume similar format YYYY? 
                // Wait, input died column: "12/8/1980". Format DD/M/YYYY.
                int dieYear = int.Parse(dieParts[dieParts.Length - 1]);
                age = dieYear - yearB;
            } else {
                // Alive logic
                int yearTarget = targetDateYear;
                int startAge = yearTarget - yearB;
                if (targetDateMonth > monthB || (targetDateMonth == monthB && targetDateDay >= dayB)) {
                    age = startAge;
                } else {
                    age = startAge - 1;
                }
            }
            
            var rels = new List<Dictionary<string, object>>();
            // Columns: Father(3), Mother(4), Brother(5), Sister(6) (0-indexed)
            // Indices: 3, 4, 5, 6
            for (int i = 3; i <= 6 && i < cols.Length; i++) {
                string relCol = cols[i].Trim();
                if (!string.IsNullOrEmpty(relCol)) {
                    var partsRel = relCol.Split(new[] {' ' }, StringSplitOptions.RemoveEmptyEntries);
                    // Assuming full name format "First Last" for these columns? 
                    // Or maybe just name? 
                    // Input: "Alfred Lennon". First Name, Last Name.
                    // Parse:
                    string fNameRel = partsRel[0];
                    string lNameRel = partsRel[partsRel.Length - 1];
                    
                    var relObj = new Dictionary<string, object>();
                    relObj["FirstName"] = fNameRel;
                    relObj["LastName"] = lNameRel;
                    relObj["Relationship"] = (i == 3) ? "Father" : ((i == 4) ? "Mother" : ((i == 5) ? "Brother" : "Sister"));
                    
                    rels.Add(relObj);
                }
            }
            
            var personObj = new Dictionary<string, object>();
            personObj["FirstName"] = fName;
            personObj["LastName"] = lName;
            // Format Birthday string for output YYYY-MM-DD? 
            // The expected JSON shows "1940-10-09".
            // We have dayB, monthB, yearB.
            var birthdayString = $"{yearB}-{monthB:D2}-{dayB:D2}";
            personObj["Birthday"] = birthdayString;
            personObj["Age"] = age;
            personObj["Relatives"] = rels; // This will be JsonArray later.
            
            peopleList.Add(personObj);
        }
        
        // Convert to JSON Array
        var jsonSerializerOptions = new JsonSerializerOptions { DiscriminatorName = "type" }; // No discriminator needed, but might need formatting?
        // The expected output is pretty printed (indentation).
        var jsonString = System.Text.Json.JsonSerializer.Serialize(peopleList, new System.Text.Json.JsonSerializerOptions() { 
            WriteIndented = true, 
            Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnescapedJson ? null : System.Text.Encodings.Web.DefaultUnicodeEncoding // Default
        }); // Actually simpler: just use default options.
        
        // Wait, we cannot easily create a custom options without imports? 
        // We can use `new JsonSerializerOptions() { WriteIndented = true }`.
        
        System.Text.Json.JsonSerializer.Write(peopleList); 
        // No, there is no Write method on List directly in the same way.
        // We need to iterate and collect or just serialize the list.
        var resultArray = new System.Collections.Generic.List<object>();
        foreach(var item in peopleList) {
            resultArray.Add(item);
        }
        
        string output = System.Text.Json.JsonSerializer.Serialize(resultArray, new System.Text.Json.JsonSerializerOptions() { WriteIndented = true });
        Console.WriteLine(output);