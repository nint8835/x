import pathlib

import ocrmypdf

unprocessed_reports_path = pathlib.Path("./reports")
processed_reports_path = pathlib.Path("./processed_reports")
sidecars_path = pathlib.Path("./sidecars")

for directory in [unprocessed_reports_path, processed_reports_path, sidecars_path]:
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
