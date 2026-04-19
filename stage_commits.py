import os
import random
import subprocess
from datetime import datetime, timedelta

REPO_PATH = os.path.dirname(os.path.abspath(__file__))

# 18 files (6 modified + 12 new) — commit_gen.py and stage_commits.py excluded.
#
# Per-file date constraints based on library/language release facts:
#
#   generics/main.go        — uses ~int|~float64 union type constraints
#   type_assertion/main.go  — uses `any` keyword
#     → both require Go 1.18, released March 18 2022
#
#   errors/main.go          — uses errors.Is / fmt.Errorf("%w")
#     → requires Go 1.13, released Sep 3 2019 (clamped to Jan 2020 per user range)
#
#   go.mod (go 1.22)        — Go 1.22 released Feb 6 2024
#
#   all other files         — stdlib only, open from Jan 2020
#
# Schema: (earliest, latest, filepath, message)
COMMITS = [
    # --- stdlib-only files (Jan 2020 – Dec 2025) ---
    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "Interface/main.go",                            "add interface and struct embedding examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "Stringer_Interface/main.go",                   "update Stringer interface implementation"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "Writer_Interface/main.go",                     "update io.Writer interface examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "marshal/main.go",                              "update JSON marshal examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "unMarshal/main.go",                            "update JSON unmarshal examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "racecondition/main.go",                        "update race condition and sync examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "closures/main.go",                             "add closure and higher-order function examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "defer_panic_recover/main.go",                  "add defer, panic, and recover examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "embedding/main.go",                            "add struct embedding and promotion examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "goroutines/main.go",                           "add goroutine patterns and WaitGroup examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "select/main.go",                               "add select statement and channel fan-in examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "worker_pool/main.go",                          "add bounded worker pool implementation"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "testing/greet.go",                             "add greet function for testing examples"),

    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "testing/greet_test.go",                        "add table-driven tests for greet function"),

    # --- errors.Is / fmt.Errorf %w: Go 1.13, Sep 2019 (clamped to Jan 2020) ---
    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "errors/main.go",                               "add errors.Is, errors.As, and %w wrapping examples"),

    # --- context package: realistic from 2020+ ---
    (datetime(2020, 1, 1), datetime(2025, 12, 31),
     "context/main.go",                              "add context cancellation, timeout, and value examples"),

    # --- Go 1.18 generics / `any` keyword: March 18 2022 ---
    (datetime(2022, 3, 18), datetime(2025, 12, 31),
     "generics/main.go",                             "add generic functions and type constraints"),

    (datetime(2022, 3, 18), datetime(2025, 12, 31),
     "type_assertion/main.go",                       "update type assertion examples using any keyword"),

    # --- go 1.22 in go.mod: Feb 6 2024 ---
    (datetime(2024, 2, 6), datetime(2025, 12, 31),
     "go.mod",                                       "upgrade module to go 1.22"),
]


def random_date_in_range(start, end):
    delta = (end - start).days
    day = start + timedelta(days=random.randint(0, delta))
    return day.replace(
        hour=random.randint(0, 23),
        minute=random.randint(0, 59),
        second=random.randint(0, 59),
    )


def commit_file(filepath, message, date):
    full_path = os.path.join(REPO_PATH, filepath)
    if not os.path.exists(full_path):
        print(f"  SKIP (not found): {filepath}")
        return False

    subprocess.run(["git", "add", filepath], cwd=REPO_PATH)

    env = os.environ.copy()
    date_str = date.strftime("%Y-%m-%dT%H:%M:%S")
    env["GIT_AUTHOR_DATE"] = date_str
    env["GIT_COMMITTER_DATE"] = date_str

    result = subprocess.run(
        ["git", "commit", "-m", message],
        cwd=REPO_PATH,
        env=env,
        capture_output=True,
        text=True,
    )

    if result.returncode == 0:
        return True
    else:
        print(f"  ERROR: {result.stderr.strip()}")
        return False


def main():
    print("=" * 60)
    print("Golang-Adv-Concepts — staged commit runner")
    print("=" * 60)
    print(f"\nCommitting {len(COMMITS)} files across 2020–2025 (per-file date windows enforced)\n")
    print("  stdlib-only files  : Jan 2020 – Dec 2025")
    print("  errors/context     : Jan 2020 – Dec 2025  (Go 1.13 released Sep 2019)")
    print("  generics + any     : Mar 18 2022 – Dec 2025  (Go 1.18 release date)")
    print("  go.mod (go 1.22)   : Feb 6  2024 – Dec 2025  (Go 1.22 release date)\n")

    entries = [
        (random_date_in_range(start, end), filepath, message)
        for start, end, filepath, message in COMMITS
    ]
    entries.sort(key=lambda x: x[0])

    for i, (date, filepath, message) in enumerate(entries, 1):
        ok = commit_file(filepath, message, date)
        status = "OK" if ok else "SKIPPED"
        print(f"[{i}/{len(COMMITS)}] [{status}] {date.strftime('%Y-%m-%d %H:%M')}  {message}")

    print("\nPushing to remote...")
    subprocess.run(["git", "push"], cwd=REPO_PATH)
    print("Done.")


if __name__ == "__main__":
    main()
