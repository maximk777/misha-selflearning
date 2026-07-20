# HTTP graceful shutdown

На сигнал создай context с deadline и вызови `server.Shutdown(ctx)`: listener закрывается, новые connections не принимаются, активным дают закончить до deadline. Обрабатывай `http.ErrServerClosed` как штатный результат. В тесте вызывай Shutdown напрямую — не посылай настоящий OS signal.
