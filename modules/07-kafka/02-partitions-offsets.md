# Partitions и offsets

Ordering гарантировано только внутри partition. Один key должен стабильно выбирать одну partition; offset — позиция чтения consumer group, не глобальный ID сообщения. Больше partitions повышают параллелизм, но усложняют ordering.
