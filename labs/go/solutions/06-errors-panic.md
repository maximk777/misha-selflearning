# Разбор 06: errors и panic

**Мисконцепция:** одинакового текста ошибки достаточно для `errors.Is`, или panic подходит для обычного отказа.

`fmt.Errorf("find user %q: %w", id, ErrNotFound)` сохраняет unwrap chain. `%v` только форматирует текст, поэтому `errors.Is` не найдёт sentinel. Ожидаемый not found — `error`; `panic` оставляют для нарушенного internal invariant или boundary с осмысленным `recover`.
