# Worker pool

Pool ограничивает число worker и получает jobs из канала. Producer закрывает jobs; owner ждёт `WaitGroup` и закрывает results. При отмене producer и workers прекращают ожидание через context. Размер pool — осознанное ограничение ресурсов, не «побольше для скорости».
