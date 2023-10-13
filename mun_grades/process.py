import pathlib
import re
import csv

import ocrmypdf

unprocessed_reports_path = pathlib.Path("./reports")
processed_reports_path = pathlib.Path("./processed_reports")
sidecars_path = pathlib.Path("./sidecars")
csvs_path = pathlib.Path("./csvs")

for directory in [unprocessed_reports_path, processed_reports_path, sidecars_path, csvs_path]:
    directory.mkdir(exist_ok=True)

for path in unprocessed_reports_path.iterdir():
    file_sidecar_path = sidecars_path / path.name.replace(".pdf", ".txt")

    print(f"OCRing {path.name}")

    if file_sidecar_path.exists():
        print("Sidecar already exists, skipping")
        continue

    ocrmypdf.ocr(
        path,
        processed_reports_path / path.name,
        sidecar=file_sidecar_path,
        # 6: "Assume a single uniform block of text."
        # Causes tesseract to not try to detect columns
        tesseract_pagesegmode=6,
    )

# WIP parsing regex. Still has known issues:
# - Doesn't handle courses with no professor listed
# - Doesn't handle courses with letters in course number (such as 499A)
# - Doesn't handle courses with OCR artifacts in the string
#   - This might be a fix on the OCR side, rather than the regex side :person_shrugging:
course_regex = re.compile(r"(\d{6}) ([\S ]+ [\dX]{4})(?:\S*) ([\S ]+) (\S+, [\S ]+) (PASS|FAIL|PWD|A|B|C|D|F)(?: (\d+) (\d+))?")

for sidecar_path in sidecars_path.iterdir():
    print(f"Processing {sidecar_path.name}")

    with open(sidecar_path) as sidecar:
        sidecar_contents = sidecar.read()
    
    with open(csvs_path / sidecar_path.name.replace(".txt", ".csv"), "w") as csv_file:
        csv_writer = csv.writer(csv_file)
        csv_writer.writerow(["Term", "Course Number", "Course Title", "Professor", "Grade", "Grade Count", "Total Enrolled"])

        for match in course_regex.finditer(sidecar_contents):
            csv_writer.writerow(match.groups())
