# Timeouts и cancellation

У server и client разные timeout-границы. Handler передаёт `r.Context()` в долгую работу и прекращает её при disconnect/deadline. Не создавай background context внутри request path: это ломает отмену. Проверяй timeout детерминированно с коротким test deadline.
