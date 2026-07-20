# Concurrency

Цель: уметь объяснить и проверить конкурентность, а не «запустить goroutine и надеяться». Порядок: GMP → goroutines/channels/select → context/sync/atomic → диагностика → patterns → shutdown. Для каждого блока: предскажи, запусти, сломай ограниченно по времени, почини и перескажи.
