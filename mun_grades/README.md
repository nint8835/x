# mun_grades

Simple script using OCR to process MUN grade reports retrieved via ATIPP, and to turn them into a computer-readable format.

## Usage

1. Install poetry
2. Create & activate a virtualenv
3. `poetry install`
4. Create a `reports` directory. Fill it with PDF copies of ATIPP responses received from MUN.
   - You can get copies of these from Matt Barter's website.
     - [Fall 2022 & Winter 2023](https://mattbarter.ca/2023/08/25/memorial-university-course-grades-for-fall-2022-and-winter-2023/)
     - [Spring 2023](https://mattbarter.ca/2023/10/12/memorial-university-course-grades-for-spring-2023/)
5. `python process.py`
   >
