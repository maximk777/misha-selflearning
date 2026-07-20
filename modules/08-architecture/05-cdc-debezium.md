# CDC и Debezium

CDC читает change log базы и публикует изменения. Debezium снимает polling-нагрузку, но требует schema evolution, snapshots и контроля lag. CDC event отражает изменение данных, не обязательно намерение домена.
