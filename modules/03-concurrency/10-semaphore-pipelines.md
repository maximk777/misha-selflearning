# Semaphore и pipeline

Semaphore — buffered channel с capacity равной лимиту: acquire перед работой, release через `defer`. Pipeline: stage читает input и закрывает свой output, когда input исчерпан; fan-out делит работу, fan-in собирает. Каждый stage должен уметь завершиться при cancellation, иначе pipeline течёт.
