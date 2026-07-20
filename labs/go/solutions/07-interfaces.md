# Разбор 07: interfaces и typed nil

**Мисконцепция:** interface равен nil, если лежащий внутри pointer nil.

Interface логически хранит dynamic type и dynamic value. У `var n Notifier = (*emailNotifier)(nil)` type есть, value nil, поэтому `n != nil`. Не нужно прятать это магической проверкой: проектируй nullable dependency явно и не вызывай method на typed nil без допустимого контракта.
