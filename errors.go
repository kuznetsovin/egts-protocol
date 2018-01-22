package main

//успешно обработано
const EGTS_PC_OK = 0

//в процессе обработки (результат обработки ещё не известен)
const EGTS_PC_IN_PROGRESS = 1

// неподдерживаемый протокол
const EGTS_PC_UNS_PROTOCOL = 128

//ошибка декодирования
const EGTS_PC_DECRYPT_ERROR = 129

//обработка запрещена
const EGTS_PC_PROC_DENIED = 130

//неверный формат заголовка
const EGTS_PC_INC_HEADERFORM = 131

//неверный формат данных
const EGTS_PC_INC_DATAFORM = 132

//неподдерживаемый тип
const EGTS_PC_UNS_TYPE = 133

//неверное количество параметров
const EGTS_PC_NOTEN_PARAMS = 134

//попытка повторной обработки
const EGTS_PC_DBL_PROC = 135

//обработка данных от источника запрещена
const EGTS_PC_PROC_SRC_DENIED = 136

//ошибка контрольной суммы заголовка
const EGTS_PC_HEADERCRC_ERROR = 137

//ошибка контрольной суммы данных
const EGTS_PC_DATACRC_ERROR = 138

//некорректная длина данных
const EGTS_PC_INVDATALEN = 139

//маршрут не найден
const EGTS_PC_ROUTE_NFOUND = 140

//маршрут закрыт
const EGTS_PC_ROUTE_CLOSED = 141

//маршрутизация запрещена
const EGTS_PC_ROUTE_DENIED = 142

//неверный адрес
const EGTS_PC_INVADDR = 143

//превышено количество ретрансляции данных
const EGTS_PC_TTLEXPIRED = 144

//нет подтверждения
const EGTS_PC_NO_ACK = 145

//объект не найден
const EGTS_PC_OBJ_NFOUND = 146

//событие не найдено
const EGTS_PC_EVNT_NFOUND = 147

//сервис не найден
const EGTS_PC_SRVC_NFOUND = 148

//сервис запрещён
const EGTS_PC_SRVC_DENIED = 149

//неизвестный тип сервиса
const EGTS_PC_SRVC_UNKN = 150

//авторизация запрещена
const EGTS_PC_AUTH_DENIED = 151

//объект уже существует
const EGTS_PC_ALREADY_EXISTS = 152

//идентификатор не найден
const EGTS_PC_ID_NFOUND = 153

//неправильная дата и время
const EGTS_PC_INC_DATETIME = 154

//ошибка ввода/вывода
const EGTS_PC_IO_ERROR = 155

//недостаточно ресурсов
const EGTS_PC_NO_RES_AVAIL = 156

//внутренний сбой модуля
const EGTS_PC_MODULE_FAULT = 157

//сбой в работе цепи питания модуля
const EGTS_PC_MODULE_PWR_FLT = 158

//сбой в работе микроконтроллера модуля
const EGTS_PC_MODULE_PROC_FLT = 159

//сбой в работе программы модуля
const EGTS_PC_MODULE_SW_FLT = 160

//сбой в работе внутреннего ПО модуля
const EGTS_PC_MODULE_FW_FLT = 161

//сбой в работе блока ввода/вывода модуля
const EGTS_PC_MODULE_IO_FLT = 162

//сбой в работе внутренней памяти модуля
const EGTS_PC_MODULE_MEM_FLT = 163

//тест не пройден
const EGTS_PC_TEST_FAILED = 164
