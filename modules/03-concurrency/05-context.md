# context

Context образует дерево: отмена родителя отменяет потомков. Передавай `context.Context` первым параметром, не храни в struct и всегда вызывай `cancel`. В цикле worker проверяй `ctx.Done()`. Context переносит deadline/cancellation, но не заменяет явное владение каналами.
