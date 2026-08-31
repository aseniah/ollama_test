import * as fs from 'fs';

    interface Relative {
      FirstName: string;
      LastName: string;
      Relationship: string;
    }

    interface Person {
      FirstName: string;
      LastName: string;
      Birthday: string;
      Age: number;
      Relatives: Relative[];
    }

    function parseDate(dateStr: string): { year: number, month: number, day: number } {
      // Format M/D/YYYY
      const parts = dateStr.split('/');
      return {
        month: parseInt(parts[0], 10),
        day: parseInt(parts[1], 10),
        year: parseInt(parts[2], 10)
      };
    }

    function formatDate(dateStr: string): string {
      // Return YYYY-MM-DD
      const { year, month, day } = parseDate(dateStr);
      return `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`;
    }

    function calculateAge(birthdayStr: string, refDateStr: string): number {
      const birth = parseDate(birthdayStr);
      const ref = parseDate(refDateStr);

      let age = ref.year - birth.year;

      // Check if birthday has passed in the reference date
      if (ref.month < birth.month || (ref.month === birth.month && ref.day < birth.day)) {
        age--;
      }

      return age;
    }

    function main() {
      const csvContent = fs.readFileSync('input/input.csv', 'utf8');
      const lines = csvContent.split('\n').filter(line => line.trim() !== '');

      const headers = lines[0].split(',');
      const rows = lines.slice(1);

      const result: Person[] = [];

      const refDateForAlive = "7/1/2025"; // July 1, 2025

      for (const row of rows) {
        const cols = row.split(',');

        // Map columns to headers
        const data: Record<string, string> = {};
        headers.forEach((h, i) => {
          data[h] = (cols[i] || '').trim();
        });

        // Name parsing
        const fullName = data['Name'];
        const nameParts = fullName.split(' ');
        const firstName = nameParts[0];
        const lastName = nameParts[nameParts.length - 1];

        // Date parsing and age calculation
        const birthdayStr = data['Birthday'];
        const diedStr = data['Died'];

        const refDate = diedStr !== 'null' && diedStr !== '' ? diedStr : refDateForAlive;
        const age = calculateAge(birthdayStr, refDate);

        // Relatives
        const relatives: Relative[] = [];

        const relationshipMap: Record<string, string> = {
          'Father': 'Father',
          'Mother': 'Mother',
          'Brother': 'Brother',
          'Sister': 'Sister'
        };

        for (const relType in relationshipMap) {
          const relName = data[relType];
          if (relName && relName !== 'null') {
            const relParts = relName.split(' ');
            const relFirst = relParts[0];
            const relLast = relParts[relParts.length - 1];

            relatives.push({
              FirstName: relFirst,
              LastName: relLast,
              Relationship: relationshipMap[relType]
            });
          }
        }

        result.push({
          FirstName: firstName,
          LastName: lastName,
          Birthday: formatDate(birthdayStr),
          Age: age,
          Relatives: relatives
        });
      }

      console.log(JSON.stringify(result, null, 2));
    }

    main();