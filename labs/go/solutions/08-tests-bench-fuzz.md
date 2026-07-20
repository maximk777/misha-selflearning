# Разбор 08: tests, benchmark и fuzz

**Мисконцепция:** один passing example доказывает алгоритм, а один benchmark доказывает production performance.

Table cases сохраняют границы, fuzz проверяет инварианту `Add(x, 0) == x` на многих input, а найденный crash/failure становится regression. `b.ReportAllocs` показывает allocations данного benchmark scenario; его сравнивают с baseline на сопоставимой машине и дополняют profile.
